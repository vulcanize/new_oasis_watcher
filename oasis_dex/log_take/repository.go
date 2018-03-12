package log_take

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type IOasisLogRepository interface {
	CreateLogTake(logTake LogTakeEntity, ethLogId int64) error
	GetLogTakesByMaker(maker string) ([]LogTakeModel, error)
}

type OasisLogRepository struct {
	*postgres.DB
}

func (repository OasisLogRepository) GetLogTakesByMaker(maker string) ([]LogTakeModel, error) {
	logTakes := []LogTakeModel{}

	err := repository.Select(&logTakes, "SELECT oasis_log_id, pair, maker, have_token, want_token, taker, take_amount, give_amount, timestamp FROM oasis.log_takes WHERE maker = $1", maker)
	if err != nil {
		return logTakes, err
	}

	return logTakes, nil
}

func (repository OasisLogRepository) CreateLogTake(logTake LogTakeEntity, ethLogId int64) error {
	oasisLogId := common.Bytes2Hex(logTake.Id[:])
	pair := common.Bytes2Hex(logTake.Pair[:])
	maker := logTake.Maker.Hex()
	haveToken := logTake.HaveToken.Hex()
	wantToken := logTake.WantToken.Hex()
	taker := logTake.Taker.Hex()
	takeAmount := logTake.TakeAmount.String()
	giveAmount := logTake.GiveAmount.String()

	_, err := repository.Exec(
		`INSERT INTO oasis.log_takes (eth_log_id, oasis_log_id, pair, maker, have_token, want_token, taker, take_amount, give_amount, timestamp)
                SELECT $1, $2, $3, $4, $5, $6, $7, $8::NUMERIC, $9::NUMERIC, $10
                WHERE NOT EXISTS (SELECT eth_log_id FROM oasis.log_takes WHERE eth_log_id = $1)`,
		ethLogId, oasisLogId, pair, maker, haveToken, wantToken, taker, takeAmount, giveAmount, logTake.Timestamp)
	if err != nil {
		return err
	}
	return nil

}
