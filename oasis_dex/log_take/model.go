package log_take

import (
	"time"
)

type LogTakeModel struct {
	ID        int64  `db:"id"` //id
	Pair      string //pair
	Guy       string //maker
	Gem       string //pay_gem
	Lot       string //give_amt
	Gal       string //taker
	Pie       string //buy_gem
	Bid       string //take_amt
	Block     int64
	Tx        string
	Timestamp time.Time `db:"time"`
}
