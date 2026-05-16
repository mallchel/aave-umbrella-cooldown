package indexer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

// Config is runtime configuration for umbrella indexing.
type Config struct {
	ChainID            int64
	RPCURL             string
	BatchBlockRange    uint64
	FinalityDepth      uint64
	IndexerStateName   string
	ProxyAddress       common.Address
	ImplementationAddr common.Address
	CooldownTopic0     common.Hash
}

type fileConfig struct {
	ChainID  int64 `json:"chainId"`
	Umbrella struct {
		ProxyAddress          string `json:"proxyAddress"`
		ImplementationAddress string `json:"implementationAddress"`
		Events                struct {
			StakerCooldownUpdated struct {
				Topic0 string `json:"topic0"`
			} `json:"stakerCooldownUpdated"`
		} `json:"events"`
	} `json:"umbrella"`
}

// LoadConfig reads indexer config from json file and environment overrides.
func LoadConfig(configPath string, rpcURL string, batchRange uint64, finalityDepth uint64) (Config, error) {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	var fc fileConfig
	if err := json.Unmarshal(b, &fc); err != nil {
		return Config{}, fmt.Errorf("decode config file: %w", err)
	}

	if !common.IsHexAddress(fc.Umbrella.ProxyAddress) {
		return Config{}, fmt.Errorf("invalid proxy address: %s", fc.Umbrella.ProxyAddress)
	}
	if !common.IsHexAddress(fc.Umbrella.ImplementationAddress) {
		return Config{}, fmt.Errorf("invalid implementation address: %s", fc.Umbrella.ImplementationAddress)
	}
	if !common.IsHexHash(fc.Umbrella.Events.StakerCooldownUpdated.Topic0) {
		return Config{}, fmt.Errorf("invalid cooldown topic0: %s", fc.Umbrella.Events.StakerCooldownUpdated.Topic0)
	}

	if batchRange == 0 {
		batchRange = 2_000
	}

	cfg := Config{
		ChainID:            fc.ChainID,
		RPCURL:             rpcURL,
		BatchBlockRange:    batchRange,
		FinalityDepth:      finalityDepth,
		IndexerStateName:   "umbrella-mainnet-indexer",
		ProxyAddress:       common.HexToAddress(fc.Umbrella.ProxyAddress),
		ImplementationAddr: common.HexToAddress(fc.Umbrella.ImplementationAddress),
		CooldownTopic0:     common.HexToHash(fc.Umbrella.Events.StakerCooldownUpdated.Topic0),
	}

	return cfg, nil
}
