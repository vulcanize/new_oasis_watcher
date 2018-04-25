package oasis_dex

import (
	"database/sql"
	"time"
)

type TradeStateModel struct {
	ID        int64          `db:"id"` //id
	Pair      string         //pair
	Guy       string         //maker
	Gem       string         //pay_gem
	Lot       string         //give_amt
	Gal       sql.NullString //taker
	Pie       string         //buy_gem
	Bid       string         //take_amt
	Block     int64
	Tx        string
	Timestamp time.Time `db:"time"`
	Deleted   bool
}
