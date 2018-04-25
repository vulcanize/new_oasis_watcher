package log_kill

import "time"

type LogKillModel struct {
	ID        int64
	Pair      string
	Guy       string
	Gem       string
	Lot       string
	Pie       string
	Bid       string
	Block     int64
	Timestamp time.Time `db:"time"`
	Tx        string
}
