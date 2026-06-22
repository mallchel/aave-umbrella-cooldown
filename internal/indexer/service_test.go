package indexer

import (
	"context"
	"database/sql"
	"errors"
	"math/big"
	"sync"
	"testing"
	"time"

	"1-task/internal/indexer/bindings"
	"1-task/internal/storage/postgres"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/mock/gomock"
)

type runCycleCaptures struct {
	mu        sync.Mutex
	lastQuery *ethereum.FilterQuery
	saves     []uint64
	flows     []postgres.WithdrawFlow
}

func (c *runCycleCaptures) recordQuery(q ethereum.FilterQuery) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lastQuery = &q
}

func (c *runCycleCaptures) query() *ethereum.FilterQuery {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.lastQuery
}

func (c *runCycleCaptures) recordSave(lastBlock uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.saves = append(c.saves, lastBlock)
}

func (c *runCycleCaptures) saveCalls() []uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]uint64(nil), c.saves...)
}

func (c *runCycleCaptures) recordFlow(flow postgres.WithdrawFlow) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.flows = append(c.flows, flow)
}

func (c *runCycleCaptures) withdrawFlows() []postgres.WithdrawFlow {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]postgres.WithdrawFlow(nil), c.flows...)
}

func TestRunCycle_FirstRunInitializesCheckpointAndProcessesStartBlockRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)
	rpc := NewMockRPCClient(ctrl)
	captures := &runCycleCaptures{}

	repo.EXPECT().GetIndexerState(gomock.Any(), "umbrella-mainnet-indexer").Return(postgres.IndexerState{}, sql.ErrNoRows)
	expectSaveIndexerState(repo, captures).Times(2)
	rpc.EXPECT().BlockNumber(gomock.Any()).Return(startBlock, nil)
	expectFilterLogs(rpc, captures, nil).Times(1)

	svc := newTestService(t, rpc, repo)

	processed, err := svc.RunCycle(context.Background())

	if err != nil {
		t.Fatalf("RunCycle returned error: %v", err)
	}
	if processed {
		t.Fatalf("expected processed=false on first initialization run")
	}
	saveCalls := captures.saveCalls()
	if len(saveCalls) != 2 {
		t.Fatalf("expected two checkpoint saves (init + cycle), got %d", len(saveCalls))
	}
	if saveCalls[0] != startBlock {
		t.Fatalf("expected first checkpoint=%d, got %d", startBlock, saveCalls[0])
	}
	if saveCalls[1] != startBlock {
		t.Fatalf("expected second checkpoint=%d, got %d", startBlock, saveCalls[1])
	}
	lastQuery := captures.query()
	if lastQuery == nil {
		t.Fatalf("expected log query for first-cycle start block")
	}
	if lastQuery.FromBlock.Uint64() != startBlock || lastQuery.ToBlock.Uint64() != startBlock {
		t.Fatalf("expected first cycle range %d..%d, got %d..%d", startBlock, startBlock, lastQuery.FromBlock.Uint64(), lastQuery.ToBlock.Uint64())
	}
}

func TestRunCycle_ProcessesLogsAndAdvancesCheckpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)
	rpc := NewMockRPCClient(ctrl)
	captures := &runCycleCaptures{}

	parsedABI, err := bindings.UmbrellaStakeTokenMetaData.GetAbi()
	if err != nil {
		t.Fatalf("parse abi: %v", err)
	}
	nonIndexed := parsedABI.Events["StakerCooldownUpdated"].Inputs.NonIndexed()
	amount := big.NewInt(12345)
	end := big.NewInt(1700000000)
	window := big.NewInt(86400)
	data, err := nonIndexed.Pack(amount, end, window)
	if err != nil {
		t.Fatalf("pack event data: %v", err)
	}

	user := common.HexToAddress("0x1111111111111111111111111111111111111111")
	logBlock := startBlock + 1
	logs := []types.Log{{
		Address:     common.HexToAddress("0xa484ab92fe32b143aee7019fc1502b1daa522d31"),
		Topics:      []common.Hash{common.HexToHash("0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419"), common.BytesToHash(user.Bytes())},
		Data:        data,
		BlockNumber: logBlock,
		TxHash:      common.HexToHash("0x2222222222222222222222222222222222222222222222222222222222222222"),
		Index:       3,
	}}

	repo.EXPECT().GetIndexerState(gomock.Any(), "umbrella-mainnet-indexer").Return(postgres.IndexerState{LastProcessedBlock: startBlock}, nil)
	expectSaveIndexerState(repo, captures).Times(1)
	expectUpsertWithdrawFlow(repo, captures).Times(1)
	rpc.EXPECT().BlockNumber(gomock.Any()).Return(startBlock+5, nil)
	expectFilterLogs(rpc, captures, logs).Times(1)
	expectHeaderByNumber(rpc, map[uint64]*types.Header{logBlock: {Time: uint64(1700000100)}}).Times(1)

	svc := newTestService(t, rpc, repo)
	processed, err := svc.RunCycle(context.Background())
	if err != nil {
		t.Fatalf("RunCycle returned error: %v", err)
	}
	if !processed {
		t.Fatalf("expected processed=true when logs are present")
	}
	lastQuery := captures.query()
	if lastQuery == nil {
		t.Fatalf("expected FilterLogs to be called")
	}
	if lastQuery.FromBlock.Uint64() != startBlock {
		t.Fatalf("unexpected from block: %d", lastQuery.FromBlock.Uint64())
	}
	flows := captures.withdrawFlows()
	if len(flows) != 1 {
		t.Fatalf("expected one upserted flow, got %d", len(flows))
	}
	if got := flows[0].SenderAddress; got != user.Hex() {
		t.Fatalf("unexpected sender address: %s", got)
	}
	if got := flows[0].AmountRaw; got != amount.String() {
		t.Fatalf("unexpected amount raw: %s", got)
	}
	saveCalls := captures.saveCalls()
	if len(saveCalls) == 0 {
		t.Fatalf("expected checkpoint to be saved")
	}
	if got := saveCalls[len(saveCalls)-1]; got != startBlock+5 {
		t.Fatalf("expected checkpoint %d, got %d", startBlock+5, got)
	}
}

func TestRunCycle_ProcessesLogsWithManyLogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)
	rpc := NewMockRPCClient(ctrl)
	captures := &runCycleCaptures{}

	parsedABI, err := bindings.UmbrellaStakeTokenMetaData.GetAbi()
	if err != nil {
		t.Fatalf("parse abi: %v", err)
	}
	nonIndexed := parsedABI.Events["StakerCooldownUpdated"].Inputs.NonIndexed()
	amount := big.NewInt(12345)
	end := big.NewInt(1700000000)
	window := big.NewInt(86400)
	data, err := nonIndexed.Pack(amount, end, window)
	if err != nil {
		t.Fatalf("pack event data: %v", err)
	}

	user := common.HexToAddress("0x1111111111111111111111111111111111111111")
	firstLogBlock := startBlock + 1
	lastLogBlock := startBlock + 1000000
	batchRange := uint64(1000000)
	batches := batchRange/(5*2000) + 1 // 101 loop iterations in RunCycle
	filterCalls := (batches-1)*5 + 1   // 5 calls per full batch, 1 on the final partial batch
	flowUpserts := filterCalls * 2     // two mocked logs returned per FilterLogs call
	saveCallsExpected := batches + 1   // +1 initialization save on sql.ErrNoRows

	logs := []types.Log{{
		Address:     common.HexToAddress("0xa484ab92fe32b143aee7019fc1502b1daa522d31"),
		Topics:      []common.Hash{common.HexToHash("0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419"), common.BytesToHash(user.Bytes())},
		Data:        data,
		BlockNumber: firstLogBlock,
		TxHash:      common.HexToHash("0x2222222222222222222222222222222222222222222222222222222222222222"),
		Index:       3,
	}, {
		Address:     common.HexToAddress("0xa484ab92fe32b143aee7019fc1502b1daa522d31"),
		Topics:      []common.Hash{common.HexToHash("0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419"), common.BytesToHash(user.Bytes())},
		Data:        data,
		BlockNumber: lastLogBlock,
		TxHash:      common.HexToHash("0x3333333333333333333333333333333333333333333333333333333333333333"),
		Index:       7,
	}}
	headers := map[uint64]*types.Header{
		firstLogBlock: {Time: uint64(1700000100)},
		lastLogBlock:  {Time: uint64(1701000000)},
	}

	repo.EXPECT().GetIndexerState(gomock.Any(), "umbrella-mainnet-indexer").Return(postgres.IndexerState{}, sql.ErrNoRows)
	expectSaveIndexerState(repo, captures).Times(int(saveCallsExpected))
	expectUpsertWithdrawFlow(repo, captures).Times(int(flowUpserts))
	rpc.EXPECT().BlockNumber(gomock.Any()).Return(startBlock+1000000, nil)
	expectFilterLogs(rpc, captures, logs).Times(int(filterCalls))
	expectHeaderByNumber(rpc, headers).Times(int(flowUpserts))

	svc := newTestService(t, rpc, repo)
	processed, err := svc.RunCycle(context.Background())

	if err != nil {
		t.Fatalf("RunCycle returned error: %v", err)
	}
	if !processed {
		t.Fatalf("expected processed=true when logs are present")
	}
	lastQuery := captures.query()
	if lastQuery == nil {
		t.Fatalf("expected FilterLogs to be called")
	}
	if lastQuery.ToBlock.Uint64() != startBlock+1000000 {
		t.Fatalf("unexpected to block: %d", lastQuery.ToBlock.Uint64())
	}
	flows := captures.withdrawFlows()
	// +2 last block with two mocked logs
	if len(flows) != 500*2+2 {
		t.Fatalf("expected %d upserted flows, got %d", 500*2+2, len(flows))
	}
	if got := flows[0].SenderAddress; got != user.Hex() {
		t.Fatalf("unexpected sender address: %s", got)
	}
	if got := flows[0].AmountRaw; got != amount.String() {
		t.Fatalf("unexpected amount raw: %s", got)
	}
	saveCalls := captures.saveCalls()
	if len(saveCalls) != int(saveCallsExpected) {
		t.Fatalf("expected %d checkpoint saves, got %d", saveCallsExpected, len(saveCalls))
	}
	if got := saveCalls[len(saveCalls)-1]; got != startBlock+1000000 {
		t.Fatalf("expected checkpoint %d, got %d", startBlock+1000000, got)
	}
}

func TestRunCycle_ReturnsErrorWhenGetCheckpointFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)
	rpc := NewMockRPCClient(ctrl)

	repo.EXPECT().GetIndexerState(gomock.Any(), "umbrella-mainnet-indexer").Return(postgres.IndexerState{}, errors.New("db down"))

	svc := newTestService(t, rpc, repo)

	processed, err := svc.RunCycle(context.Background())
	if err == nil {
		t.Fatalf("expected error")
	}
	if processed {
		t.Fatalf("expected processed=false on error")
	}
}

func expectFilterLogs(rpc *MockRPCClient, captures *runCycleCaptures, logs []types.Log) *gomock.Call {
	return rpc.EXPECT().FilterLogs(gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
			captures.recordQuery(q)
			return logs, nil
		},
	)
}

func expectHeaderByNumber(rpc *MockRPCClient, headers map[uint64]*types.Header) *gomock.Call {
	return rpc.EXPECT().HeaderByNumber(gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, number *big.Int) (*types.Header, error) {
			if headers == nil {
				return &types.Header{Time: uint64(time.Now().UTC().Unix())}, nil
			}

			header := headers[number.Uint64()]
			if header == nil {
				return nil, errors.New("header not found")
			}
			return header, nil
		},
	)
}

func expectSaveIndexerState(repo *MockRepository, captures *runCycleCaptures) *gomock.Call {
	return repo.EXPECT().SaveIndexerState(gomock.Any(), "umbrella-mainnet-indexer", gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, _ string, lastBlock uint64, _ time.Time) error {
			captures.recordSave(lastBlock)
			return nil
		},
	)
}

func expectUpsertWithdrawFlow(repo *MockRepository, captures *runCycleCaptures) *gomock.Call {
	return repo.EXPECT().UpsertWithdrawFlow(gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, flow postgres.WithdrawFlow) error {
			captures.recordFlow(flow)
			return nil
		},
	)
}

func newTestService(t *testing.T, rpc RPCClient, repo Repository) *Service {
	t.Helper()
	parsedABI, err := bindings.UmbrellaStakeTokenMetaData.GetAbi()
	if err != nil {
		t.Fatalf("parse abi: %v", err)
	}

	return &Service{
		cfg: Config{
			ChainID:          1,
			BatchBlockRange:  2000,
			IndexerStateName: "umbrella-mainnet-indexer",
			ProxyAddress:     common.HexToAddress("0xa484ab92fe32b143aee7019fc1502b1daa522d31"),
			CooldownTopic0:   common.HexToHash("0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419"),
			WithdrawTopic0:   common.HexToHash("0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db"),
		},
		client:      rpc,
		repo:        repo,
		contractABI: *parsedABI,
	}
}
