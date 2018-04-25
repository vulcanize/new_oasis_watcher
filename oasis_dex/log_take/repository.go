package log_take

import (
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type Datastore interface {
	Create(logTake LogTakeModel, vulcanizeLogId int64) error
}

type Repository struct {
	*postgres.DB
}

func (repository Repository) Create(logTake LogTakeModel, vulcanizeLogId int64) error {
	_, err := repository.Exec(
		`INSERT INTO oasis.log_take (vulcanize_log_id, id, pair, guy, gem, lot, gal, pie, bid, block, time, tx)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
                ON CONFLICT (vulcanize_log_id) DO NOTHING`,
		vulcanizeLogId, logTake.ID, logTake.Pair, logTake.Guy, logTake.Gem, logTake.Lot, logTake.Gal, logTake.Pie, logTake.Bid, logTake.Block, logTake.Timestamp, logTake.Tx,
	)
	if err != nil {
		return err
	}
	return nil
}
