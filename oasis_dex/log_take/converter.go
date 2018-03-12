package log_take

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

type ILogTakeConverter interface {
	Convert(watchedEvent core.WatchedEvent) (*LogTakeEntity, error)
}

type LogTakeConverter struct {
}

func (LogTakeConverter) Convert(watchedEvent core.WatchedEvent) (*LogTakeEntity, error) {
	result := &LogTakeEntity{}
	parser, err := geth.ParseAbi(`[{"constant":false,"inputs":[],"name":"stop","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"haveToken","type":"address"},{"name":"wantToken","type":"address"},{"name":"haveAmount","type":"uint128"},{"name":"wantAmount","type":"uint128"}],"name":"make","outputs":[{"name":"id","type":"bytes32"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"owner_","type":"address"}],"name":"setOwner","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"last_offer_id","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"cancel","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getOffer","outputs":[{"name":"","type":"uint256"},{"name":"","type":"address"},{"name":"","type":"uint256"},{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"bytes32"},{"name":"maxTakeAmount","type":"uint128"}],"name":"take","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"getTime","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"close_time","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"lifetime","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"stopped","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id_","type":"bytes32"}],"name":"bump","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"authority_","type":"address"}],"name":"setAuthority","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"isActive","outputs":[{"name":"active","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"offers","outputs":[{"name":"sell_how_much","type":"uint256"},{"name":"sell_which_token","type":"address"},{"name":"buy_how_much","type":"uint256"},{"name":"buy_which_token","type":"address"},{"name":"owner","type":"address"},{"name":"active","type":"bool"},{"name":"timestamp","type":"uint64"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"bytes32"}],"name":"kill","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"authority","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"isClosed","outputs":[{"name":"closed","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getOwner","outputs":[{"name":"owner","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"quantity","type":"uint256"}],"name":"buy","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"sell_how_much","type":"uint256"},{"name":"sell_which_token","type":"address"},{"name":"buy_how_much","type":"uint256"},{"name":"buy_which_token","type":"address"}],"name":"offer","outputs":[{"name":"id","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"lifetime_","type":"uint256"}],"payable":false,"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint256"}],"name":"ItemUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"sell_how_much","type":"uint256"},{"indexed":true,"name":"sell_which_token","type":"address"},{"indexed":false,"name":"buy_how_much","type":"uint256"},{"indexed":true,"name":"buy_which_token","type":"address"}],"name":"Trade","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"haveToken","type":"address"},{"indexed":false,"name":"wantToken","type":"address"},{"indexed":false,"name":"haveAmount","type":"uint128"},{"indexed":false,"name":"wantAmount","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogMake","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"haveToken","type":"address"},{"indexed":false,"name":"wantToken","type":"address"},{"indexed":false,"name":"haveAmount","type":"uint128"},{"indexed":false,"name":"wantAmount","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogBump","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"haveToken","type":"address"},{"indexed":false,"name":"wantToken","type":"address"},{"indexed":true,"name":"taker","type":"address"},{"indexed":false,"name":"takeAmount","type":"uint128"},{"indexed":false,"name":"giveAmount","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogTake","type":"event"}, {"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"haveToken","type":"address"},{"indexed":false,"name":"wantToken","type":"address"},{"indexed":false,"name":"haveAmount","type":"uint128"},{"indexed":false,"name":"wantAmount","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogKill","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"authority","type":"address"}],"name":"LogSetAuthority","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"}],"name":"LogSetOwner","type":"event"}]`)
	if err != nil {
		return result, err
	}
	contract := bind.NewBoundContract(common.StringToAddress("0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f"), parser, nil, nil, nil)
	event := convertToLog(watchedEvent)
	err = contract.UnpackLog(result, "LogTake", event)
	if err != nil {
		return result, err
	}
	return result, nil
}

func convertToLog(watchedEvent core.WatchedEvent) types.Log {
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
