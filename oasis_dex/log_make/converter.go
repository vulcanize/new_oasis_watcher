package log_make

import (
	"strings"
	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

type Converter interface {
	ToEntity(watchedEvent core.WatchedEvent) (*LogMakeEntity, error)
	ToModel(entity LogMakeEntity) LogMakeModel
}

type LogMakeConverter struct {
}

func (LogMakeConverter) ToEntity(watchedEvent core.WatchedEvent) (*LogMakeEntity, error) {
	result := &LogMakeEntity{}
	contract := bind.NewBoundContract(common.HexToAddress(constants.ContractAddress), constants.ABI, nil, nil, nil)
	event := helpers.ConvertToLog(watchedEvent)
	err := contract.UnpackLog(result, "LogMake", event)
	if err != nil {
		return result, err
	}
	result.Block = watchedEvent.BlockNumber
	result.TransactionHash = watchedEvent.TxHash

	return result, nil
}

func (LogMakeConverter) ToModel(logMakeEntity LogMakeEntity) LogMakeModel {
	id := common.BytesToHash(logMakeEntity.Id[:]).Big().Int64()
	pair := strings.ToLower(common.ToHex(logMakeEntity.Pair[:]))
	guy := strings.ToLower(logMakeEntity.Maker.Hex())
	gem := strings.ToLower(logMakeEntity.Pay_gem.Hex())
	lot := logMakeEntity.Pay_amt.String()
	pie := strings.ToLower(logMakeEntity.Buy_gem.Hex())
	bid := logMakeEntity.Buy_amt.String()
	block := logMakeEntity.Block
	timestamp := time.Unix(int64(logMakeEntity.Timestamp), 0)
	tx := logMakeEntity.TransactionHash
	return LogMakeModel{
		ID:        id,
		Pair:      pair,
		Guy:       guy,
		Gem:       gem,
		Lot:       lot,
		Pie:       pie,
		Bid:       bid,
		Block:     block,
		Timestamp: timestamp,
		Tx:        tx,
	}
}
