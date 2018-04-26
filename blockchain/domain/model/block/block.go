package block

import (
	"encoding/json"
	"github.com/it-chain/yggdrasill/block"
	"time"
	"github.com/it-chain/it-chain-Engine/blockchain/domain/service"
)

func CreateGenesisBlock() (*block.DefaultBlock, error) {
	byteValue, err := service.ConfigFromJson("GenesisBlockConfig.json")
	if err != nil {
		return nil, err
	}

	var GenesisBlock *block.DefaultBlock
	json.Unmarshal(byteValue, &GenesisBlock)
	GenesisBlock.Header.TimeStamp = time.Now()
	return GenesisBlock, nil
}