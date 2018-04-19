package log_kill

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

type Converter interface {
	ToModel(watchedEvent core.WatchedEvent) (*LogKillModel, error)
}

type LogKillConverter struct {
}

func (LogKillConverter) ToModel(watchedEvent core.WatchedEvent) (*LogKillModel, error) {
	id := common.HexToHash(watchedEvent.Topic1).Big().Uint64()
	return &LogKillModel{ID: id}, nil
}
