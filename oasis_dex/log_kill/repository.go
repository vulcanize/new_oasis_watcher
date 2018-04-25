package log_kill

import (
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type Datastore interface {
	Create(model LogKillModel, vulcanizeLogId int64) error
}

type Repository struct {
	*postgres.DB
}

func (repository Repository) Create(logKillModel LogKillModel, vulcanizeLogId int64) error {
	_, err := repository.DB.Exec(

		`INSERT INTO oasis.kill (id, vulcanize_log_id, pair, gem, lot, pie, bid, guy, block, "time", tx)
               VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
                ON CONFLICT (vulcanize_log_id) DO NOTHING`,
		logKillModel.ID, vulcanizeLogId, logKillModel.Pair, logKillModel.Gem, logKillModel.Lot, logKillModel.Pie, logKillModel.Bid, logKillModel.Guy, logKillModel.Block, logKillModel.Timestamp, logKillModel.Tx)

	if err != nil {
		return err
	}

	return nil
}
