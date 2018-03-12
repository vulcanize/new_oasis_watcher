package log_take

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type LogTakeEntity struct {
	Id         [32]byte
	Pair       [32]byte
	Maker      common.Address
	HaveToken  common.Address
	WantToken  common.Address
	Taker      common.Address
	TakeAmount *big.Int
	GiveAmount *big.Int
	Timestamp  uint64
}
