package indexer

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"sync"
	"time"

	"1-task/internal/indexer/bindings"
	"1-task/internal/storage/postgres"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//go:generate go run go.uber.org/mock/mockgen@v0.6.0 -source=$GOFILE -destination=mocks_test.go -package=$GOPACKAGE

type cooldownEvent struct {
	Amount        *big.Int
	EndOfCooldown *big.Int
	UnstakeWindow *big.Int
}

type withdrawEvent struct {
	Assets *big.Int
	Shares *big.Int
}

type RPCClient interface {
	BlockNumber(ctx context.Context) (uint64, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	Close()
}

type Repository interface {
	GetIndexerState(ctx context.Context, name string) (postgres.IndexerState, error)
	SaveIndexerState(ctx context.Context, name string, lastBlock uint64, processedAt time.Time) error
	UpsertWithdrawFlow(ctx context.Context, flow postgres.WithdrawFlow) error
}

// Service indexes umbrella events and reconciles queued rows.
type Service struct {
	cfg           Config
	client        RPCClient
	repo          Repository
	contractABI   abi.ABI
	assetDecimals uint8
}

// Close closes underlying RPC connections.
func (s *Service) Close() {
	if s == nil || s.client == nil {
		return
	}
	s.client.Close()
}

// NewService creates indexer service.
func NewService(ctx context.Context, cfg Config, repo *postgres.Repository) (*Service, error) {
	client, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("dial rpc: %w", err)
	}

	parsedABI, err := bindings.UmbrellaStakeTokenMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("parse stake token abi: %w", err)
	}

	assetDecimals, err := readAssetDecimals(ctx, client, cfg.ProxyAddress)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("read asset metadata: %w", err)
	}

	return &Service{
		cfg:           cfg,
		client:        client,
		repo:          repo,
		contractABI:   *parsedABI,
		assetDecimals: assetDecimals,
	}, nil
}

func exponentialBackoff[T any](getter func() (T, error), currentBlock, toBlock uint64, ctx context.Context) (T, error) {
	var data T
	var err error
	var zero T

	for attempts := range 20 {
		data, err = getter()

		if err == nil {
			break
		}

		if ctx.Err() != nil {
			return zero, ctx.Err()
		}

		duration := 300 * time.Millisecond * time.Duration(attempts)
		select {
		case <-ctx.Done():
			return zero, ctx.Err()
		case <-time.After(duration):
			fmt.Println("Using backoff for", currentBlock, toBlock, duration)
		}
	}

	return data, err
}

// https://etherscan.io/tx/0xfa2d65feee27b96d70d5d2808c3b89f9d2b7240ed4ff37f679b88d55dba5c658
var startBlock = uint64(16832680)

// RunCycle processes one block range and returns true when something was indexed.
func (s *Service) RunCycle(ctx context.Context) (bool, error) {
	checkpoint, err := s.repo.GetIndexerState(ctx, s.cfg.IndexerStateName)

	if err == sql.ErrNoRows {
		// Initialize checkpoint to configured start block on first run.
		if err := s.repo.SaveIndexerState(ctx, s.cfg.IndexerStateName, startBlock, time.Now().UTC()); err != nil {
			return false, fmt.Errorf("initialize indexer checkpoint: %w", err)
		}
		checkpoint = postgres.IndexerState{LastProcessedBlock: startBlock}
	} else if err != nil {
		return false, fmt.Errorf("load indexer checkpoint: %w", err)
	}

	latest, err := s.client.BlockNumber(ctx)
	fmt.Println("RunCycle", checkpoint, latest)
	if err != nil {
		return false, fmt.Errorf("read latest block: %w", err)
	}

	safeHead := latest
	if s.cfg.FinalityDepth > 0 {
		safeHead = max(0, latest-s.cfg.FinalityDepth)
	}

	fromBlock := min(checkpoint.LastProcessedBlock, safeHead)

	var wg sync.WaitGroup
	// Concurrency throttle
	batchSize := uint64(5)
	hasProcessedLogs := false

	for batchStart := fromBlock; batchStart <= safeHead; batchStart += batchSize * s.cfg.BatchBlockRange {
		// Thread-safe structure to track successfully processed block numbers
		var mu sync.Mutex
		var highestBlockInBatch uint64 = 0

		var toBlock uint64

		// Launch exactly up to 5 goroutines for this batch
		for i := range batchSize {
			if toBlock == safeHead {
				continue
			}
			currentBlock := batchStart + i*s.cfg.BatchBlockRange
			toBlock = min(currentBlock+s.cfg.BatchBlockRange-1, safeHead)

			wg.Add(1)

			go func(toBlock uint64) {
				defer wg.Done()

				fmt.Printf("Fetching batch: blocks %d to %d\n", currentBlock, toBlock)

				getLogs := func() ([]types.Log, error) {
					logs, err := s.client.FilterLogs(ctx, ethereum.FilterQuery{
						FromBlock: new(big.Int).SetUint64(currentBlock),
						ToBlock:   new(big.Int).SetUint64(toBlock),
						Addresses: []common.Address{s.cfg.ProxyAddress},
						Topics: [][]common.Hash{{
							s.cfg.CooldownTopic0,
							s.cfg.WithdrawTopic0,
						}},
					})

					if err != nil {
						fmt.Printf("filter logs: %v", err)
					}

					return logs, err
				}

				logs, err := exponentialBackoff(getLogs, currentBlock, toBlock, ctx)
				if err != nil {
					panic(fmt.Sprintf("Failed to fetch logs for blocks %d to %d after retries: %v\n", currentBlock, toBlock, err))
				}

				mu.Lock()
				if len(logs) > 0 {
					hasProcessedLogs = true
				}
				mu.Unlock()

				blockTimeCache := make(map[uint64]time.Time)
				for _, lg := range logs {
					fmt.Println("logs for block number:", lg.BlockNumber)

					if len(lg.Topics) == 0 {
						continue
					}

					bt, ok := blockTimeCache[lg.BlockNumber]
					if !ok {
						getHeader := func() (*types.Header, error) {
							return s.client.HeaderByNumber(ctx, new(big.Int).SetUint64(lg.BlockNumber))
						}

						h, err := exponentialBackoff(getHeader, lg.BlockNumber, lg.BlockNumber, ctx)

						if err != nil {
							panic(fmt.Sprintf("load block header %d: %v", lg.BlockNumber, err))
						}
						bt = time.Unix(int64(h.Time), 0).UTC()
						blockTimeCache[lg.BlockNumber] = bt
					}

					switch lg.Topics[0] {
					case s.cfg.CooldownTopic0:
						if err := s.handleCooldownLog(ctx, lg, bt); err != nil {
							panic(fmt.Sprintf("handle cooldown log: %v", err))
						}
					case s.cfg.WithdrawTopic0:
						if err := s.handleWithdrawLog(ctx, lg, bt); err != nil {
							panic(fmt.Sprintf("handle withdraw log: %v", err))
						}
					}
				}

				mu.Lock()
				if toBlock > highestBlockInBatch {
					highestBlockInBatch = toBlock
				}
				mu.Unlock()
			}(toBlock)
		}

		wg.Wait()

		fmt.Println("Save highestBlockInBatch:", highestBlockInBatch)
		if err := ctx.Err(); err != nil {
			return hasProcessedLogs, fmt.Errorf("context canceled before saving indexer checkpoint: %w", err)
		}
		if err := s.repo.SaveIndexerState(ctx, s.cfg.IndexerStateName, highestBlockInBatch, time.Now().UTC()); err != nil {
			return false, fmt.Errorf("save indexer checkpoint: %w", err)
		}
	}

	wg.Wait()

	return hasProcessedLogs, nil
}

func (s *Service) handleCooldownLog(ctx context.Context, lg types.Log, blockTime time.Time) error {
	if len(lg.Topics) < 2 {
		return fmt.Errorf("cooldown log missing topics")
	}

	var ev cooldownEvent
	if err := s.contractABI.UnpackIntoInterface(&ev, "StakerCooldownUpdated", lg.Data); err != nil {
		return fmt.Errorf("decode cooldown event: %w", err)
	}

	user := common.HexToAddress(lg.Topics[1].Hex())

	flow := postgres.WithdrawFlow{
		ChainID:          s.cfg.ChainID,
		TxHash:           lg.TxHash.Hex(),
		LogIndex:         int32(lg.Index),
		BlockNumber:      int64(lg.BlockNumber),
		BlockTime:        blockTime,
		SenderAddress:    user.Hex(),
		EventType:        postgres.FlowEventTypeRequest,
		AmountRaw:        ev.Amount.String(),
		AmountNormalized: normalizeAmount(ev.Amount, s.assetDecimals),
	}

	if err := s.repo.UpsertWithdrawFlow(ctx, flow); err != nil {
		return fmt.Errorf("upsert cooldown flow: %w", err)
	}

	return nil
}

func (s *Service) handleWithdrawLog(ctx context.Context, lg types.Log, blockTime time.Time) error {
	if len(lg.Topics) < 4 {
		return fmt.Errorf("withdraw log missing topics")
	}

	var ev withdrawEvent
	if err := s.contractABI.UnpackIntoInterface(&ev, "Withdraw", lg.Data); err != nil {
		return fmt.Errorf("decode withdraw event: %w", err)
	}

	sender := common.HexToAddress(lg.Topics[1].Hex())

	flow := postgres.WithdrawFlow{
		ChainID:          s.cfg.ChainID,
		TxHash:           lg.TxHash.Hex(),
		LogIndex:         int32(lg.Index),
		BlockNumber:      int64(lg.BlockNumber),
		BlockTime:        blockTime,
		SenderAddress:    sender.Hex(),
		EventType:        postgres.FlowEventTypeWithdraw,
		AmountRaw:        ev.Assets.String(),
		AmountNormalized: normalizeAmount(ev.Assets, s.assetDecimals),
	}

	if err := s.repo.UpsertWithdrawFlow(ctx, flow); err != nil {
		return fmt.Errorf("upsert withdraw flow: %w", err)
	}

	return nil
}

func readAssetDecimals(ctx context.Context, client *ethclient.Client, proxy common.Address) (uint8, error) {
	proxyCaller, err := bindings.NewUmbrellaStakeTokenCaller(proxy, client)
	if err != nil {
		return 0, fmt.Errorf("create proxy caller: %w", err)
	}

	callOpts := &bind.CallOpts{Context: ctx}

	assetAddr, err := proxyCaller.Asset(callOpts)
	if err != nil {
		return 0, fmt.Errorf("call asset(): %w", err)
	}

	assetCaller, err := bindings.NewUmbrellaStakeTokenCaller(assetAddr, client)
	if err != nil {
		return 0, fmt.Errorf("create asset caller: %w", err)
	}

	decimals, err := assetCaller.Decimals(callOpts)
	if err != nil {
		return 0, fmt.Errorf("call decimals(): %w", err)
	}

	return decimals, nil
}

func normalizeAmount(amount *big.Int, decimals uint8) string {
	if amount == nil {
		return "0"
	}
	if decimals == 0 {
		return amount.String()
	}

	denom := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	rat := new(big.Rat).SetInt(amount)
	rat.Quo(rat, new(big.Rat).SetInt(denom))

	precision := min(int(decimals), 18)

	return rat.FloatString(precision)
}
