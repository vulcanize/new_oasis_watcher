package log_make

import (
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type Datastore interface {
	Create(logMakeModel LogMakeModel, vulcanizeLogId int64) error
}

type Repository struct {
	*postgres.DB
}

func (repository Repository) Create(logMakeModel LogMakeModel, vulcanizeLogId int64) error {
	_, err := repository.DB.Exec(
		`INSERT INTO oasis.log_make (id, vulcanize_log_id, pair, gem, lot, pie, bid, guy, block, "time", tx)
               VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			   ON CONFLICT (vulcanize_log_id) DO NOTHING`,
		logMakeModel.ID, vulcanizeLogId, logMakeModel.Pair, logMakeModel.Gem, logMakeModel.Lot, logMakeModel.Pie, logMakeModel.Bid, logMakeModel.Guy, logMakeModel.Block, logMakeModel.Timestamp, logMakeModel.Tx,
	)

	if err != nil {
		return err
	}

	return nil
}
