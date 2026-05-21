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
)

type mockRPCClient struct {
	latestBlock    uint64
	blockNumberErr error
	logs           []types.Log
	filterLogsErr  error
	headers        map[uint64]*types.Header
	headerErr      error
	lastQuery      *ethereum.FilterQuery
}

func (m *mockRPCClient) BlockNumber(ctx context.Context) (uint64, error) {
	if m.blockNumberErr != nil {
		return 0, m.blockNumberErr
	}
	return m.latestBlock, nil
}

func (m *mockRPCClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	m.lastQuery = &q
	if m.filterLogsErr != nil {
		return nil, m.filterLogsErr
	}
	return m.logs, nil
}

func (m *mockRPCClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	if m.headerErr != nil {
		return nil, m.headerErr
	}
	if m.headers == nil {
		return &types.Header{Time: uint64(time.Now().UTC().Unix())}, nil
	}
	h := m.headers[number.Uint64()]
	if h == nil {
		return nil, errors.New("header not found")
	}
	return h, nil
}

func (m *mockRPCClient) Close() {}

type mockRepository struct {
	mu         sync.Mutex
	state      postgres.IndexerState
	getErr     error
	saveCalls  []uint64
	savedNames []string
	flows      []postgres.WithdrawFlow
	saveErr    error
	upsertErr  error
}

func (m *mockRepository) GetIndexerState(ctx context.Context, name string) (postgres.IndexerState, error) {
	if m.getErr != nil {
		return postgres.IndexerState{}, m.getErr
	}
	return m.state, nil
}

func (m *mockRepository) SaveIndexerState(ctx context.Context, name string, lastBlock uint64, processedAt time.Time) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.saveErr != nil {
		return m.saveErr
	}
	m.saveCalls = append(m.saveCalls, lastBlock)
	m.savedNames = append(m.savedNames, name)
	m.state = postgres.IndexerState{Name: name, LastProcessedBlock: lastBlock, LastProcessedTime: processedAt}
	m.getErr = nil
	return nil
}

func (m *mockRepository) UpsertWithdrawFlow(ctx context.Context, flow postgres.WithdrawFlow) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.upsertErr != nil {
		return m.upsertErr
	}
	m.flows = append(m.flows, flow)
	return nil
}

func TestRunCycle_FirstRunInitializesCheckpointAndProcessesStartBlockRange(t *testing.T) {
	repo := &mockRepository{getErr: sql.ErrNoRows}
	rpc := &mockRPCClient{latestBlock: startBlock}
	svc := newTestService(t, rpc, repo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processed, err := svc.RunCycle(ctx, cancel)

	if err != nil {
		t.Fatalf("RunCycle returned error: %v", err)
	}
	if processed {
		t.Fatalf("expected processed=false on first initialization run")
	}
	if len(repo.saveCalls) != 2 {
		t.Fatalf("expected two checkpoint saves (init + cycle), got %d", len(repo.saveCalls))
	}
	if repo.saveCalls[0] != startBlock {
		t.Fatalf("expected first checkpoint=%d, got %d", startBlock, repo.saveCalls[0])
	}
	if repo.saveCalls[1] != startBlock {
		t.Fatalf("expected second checkpoint=%d, got %d", startBlock, repo.saveCalls[1])
	}
	if rpc.lastQuery == nil {
		t.Fatalf("expected log query for first-cycle start block")
	}
	if rpc.lastQuery.FromBlock.Uint64() != startBlock || rpc.lastQuery.ToBlock.Uint64() != startBlock {
		t.Fatalf("expected first cycle range %d..%d, got %d..%d", startBlock, startBlock, rpc.lastQuery.FromBlock.Uint64(), rpc.lastQuery.ToBlock.Uint64())
	}
}

func TestRunCycle_ProcessesLogsAndAdvancesCheckpoint(t *testing.T) {
	repo := &mockRepository{state: postgres.IndexerState{LastProcessedBlock: startBlock}}
	rpc := &mockRPCClient{latestBlock: startBlock + 5}

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
	rpc.logs = []types.Log{{
		Address:     common.HexToAddress("0xa484ab92fe32b143aee7019fc1502b1daa522d31"),
		Topics:      []common.Hash{common.HexToHash("0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419"), common.BytesToHash(user.Bytes())},
		Data:        data,
		BlockNumber: logBlock,
		TxHash:      common.HexToHash("0x2222222222222222222222222222222222222222222222222222222222222222"),
		Index:       3,
	}}
	rpc.headers = map[uint64]*types.Header{logBlock: {Time: uint64(1700000100)}}

	svc := newTestService(t, rpc, repo)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processed, err := svc.RunCycle(ctx, cancel)
	if err != nil {
		t.Fatalf("RunCycle returned error: %v", err)
	}
	if !processed {
		t.Fatalf("expected processed=true when logs are present")
	}
	if rpc.lastQuery == nil {
		t.Fatalf("expected FilterLogs to be called")
	}
	if rpc.lastQuery.FromBlock.Uint64() != startBlock {
		t.Fatalf("unexpected from block: %d", rpc.lastQuery.FromBlock.Uint64())
	}
	if len(repo.flows) != 1 {
		t.Fatalf("expected one upserted flow, got %d", len(repo.flows))
	}
	if got := repo.flows[0].SenderAddress; got != user.Hex() {
		t.Fatalf("unexpected sender address: %s", got)
	}
	if got := repo.flows[0].AmountRaw; got != amount.String() {
		t.Fatalf("unexpected amount raw: %s", got)
	}
	if len(repo.saveCalls) == 0 {
		t.Fatalf("expected checkpoint to be saved")
	}
	if got := repo.saveCalls[len(repo.saveCalls)-1]; got != startBlock+5 {
		t.Fatalf("expected checkpoint %d, got %d", startBlock+5, got)
	}
}

func TestRunCycle_ProcessesLogsWithManyLogs(t *testing.T) {
	repo := &mockRepository{getErr: sql.ErrNoRows}
	rpc := &mockRPCClient{latestBlock: startBlock + 1000000}

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

	rpc.logs = []types.Log{{
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
	rpc.headers = map[uint64]*types.Header{
		firstLogBlock: {Time: uint64(1700000100)},
		lastLogBlock:  {Time: uint64(1701000000)},
	}

	svc := newTestService(t, rpc, repo)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processed, err := svc.RunCycle(ctx, cancel)

	if err != nil {
		t.Fatalf("RunCycle returned error: %v", err)
	}
	if !processed {
		t.Fatalf("expected processed=true when logs are present")
	}
	if rpc.lastQuery == nil {
		t.Fatalf("expected FilterLogs to be called")
	}
	if rpc.lastQuery.ToBlock.Uint64() != startBlock+1000000 {
		t.Fatalf("unexpected to block: %d", rpc.lastQuery.ToBlock.Uint64())
	}
	// +2 last block with two mocked logs
	if len(repo.flows) != 500*2+2 {
		t.Fatalf("expected %d upserted flows, got %d", 500*2+2, len(repo.flows))
	}
	if got := repo.flows[0].SenderAddress; got != user.Hex() {
		t.Fatalf("unexpected sender address: %s", got)
	}
	if got := repo.flows[0].AmountRaw; got != amount.String() {
		t.Fatalf("unexpected amount raw: %s", got)
	}
	if len(repo.saveCalls) == 0 {
		t.Fatalf("expected checkpoint to be saved")
	}
	if got := repo.saveCalls[len(repo.saveCalls)-1]; got != startBlock+1000000 {
		t.Fatalf("expected checkpoint %d, got %d", startBlock+1000000, got)
	}
}

func TestRunCycle_ReturnsErrorWhenGetCheckpointFails(t *testing.T) {
	repo := &mockRepository{getErr: errors.New("db down")}
	rpc := &mockRPCClient{latestBlock: startBlock + 10}
	svc := newTestService(t, rpc, repo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processed, err := svc.RunCycle(ctx, cancel)
	if err == nil {
		t.Fatalf("expected error")
	}
	if processed {
		t.Fatalf("expected processed=false on error")
	}
}

func newTestService(t *testing.T, rpc rpcClient, repo repository) *Service {
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
