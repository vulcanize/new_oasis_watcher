package log_take

type LogTakeModel struct {
	LogID     string `db:"oasis_log_id"`
	Pair      string
	Maker     string
	HaveToken string `db:"have_token"`
	WantToken string `db:"want_token"`
	Taker     string
	// should maybe convert these to big ints if there's a use case beyond GraphQL
	TakeAmount string `db:"take_amount"`
	GiveAmount string `db:"give_amount"`
	Timestamp  int
}
