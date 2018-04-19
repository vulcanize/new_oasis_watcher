package helpers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

func ConvertToLog(watchedEvent core.WatchedEvent) types.Log {
	return types.Log{
		Address:     common.HexToAddress(watchedEvent.Address),
		Topics:      createTopics(watchedEvent.Topic0, watchedEvent.Topic1, watchedEvent.Topic2, watchedEvent.Topic3),
		Data:        hexutil.MustDecode(watchedEvent.Data),
		BlockNumber: uint64(watchedEvent.BlockNumber),
		TxHash:      common.HexToHash(watchedEvent.TxHash),
		TxIndex:     0,
		BlockHash:   common.HexToHash("0x0"),
		Index:       uint(watchedEvent.Index),
		Removed:     false,
	}
}

func createTopics(topic0 string, topic1 string, topic2 string, topic3 string) []common.Hash {
	return []common.Hash{common.HexToHash(topic0), common.HexToHash(topic1), common.HexToHash(topic2), common.HexToHash(topic3)}
}
