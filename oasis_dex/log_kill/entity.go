package log_kill

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type LogKillEntity struct {
	Id              [32]byte
	Pair            [32]byte
	Maker           common.Address
	Pay_gem         common.Address
	Buy_gem         common.Address
	Pay_amt         *big.Int
	Buy_amt         *big.Int
	Block           int64
	Timestamp       uint64
	TransactionHash string
}
