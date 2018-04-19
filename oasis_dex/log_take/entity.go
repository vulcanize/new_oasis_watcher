package log_take

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type LogTakeEntity struct {
	Id        [32]byte
	Pair      [32]byte
	Maker     common.Address
	Pay_gem   common.Address //TODO revisit naming when go-eth fixes abi unpacking
	Buy_gem   common.Address //TODO revisit naming when go-eth fixes abi unpacking
	Taker     common.Address
	Take_amt  *big.Int //TODO revisit naming when go-eth fixes abi unpacking
	Give_amt  *big.Int //TODO revisit naming when go-eth fixes abi unpacking
	Block     int64
	Tx        string
	Timestamp uint64
}
