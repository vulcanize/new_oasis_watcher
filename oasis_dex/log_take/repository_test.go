package log_take_test

import (
	"math/big"

	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

var _ = Describe("Logs Repository", func() {
	var db *postgres.DB
	var oasisLogRepository log_take.OasisLogRepository
	var filterRepository repositories.FilterRepository
	var logRepository repositories.LogRepository

	BeforeEach(func() {
		var err error
		db, err = postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).NotTo(HaveOccurred())
		db.Query(`DELETE FROM maker.cups`)
		db.Query(`DELETE FROM maker.peps`)
		db.Query(`DELETE FROM logs`)
		db.Query(`DELETE FROM log_filters`)
		logRepository = repositories.LogRepository{DB: db}
		filterRepository = repositories.FilterRepository{DB: db}
		oasisLogRepository = log_take.OasisLogRepository{DB: db}
	})

	It("has a test", func() {
		Expect(1).To(Equal(1))
	})

	Describe("Creating a new cups record", func() {
		It("inserts a new cup", func() {
			err := logRepository.CreateLogs([]core.Log{{}})
			Expect(err).ToNot(HaveOccurred())
			var ethLogID int64
			err = logRepository.Get(&ethLogID, `Select id from logs`)
			Expect(err).ToNot(HaveOccurred())

			oasisLogId := [32]byte{1, 2, 3, 4, 5}
			pair := [32]byte{6, 7, 8, 9, 0}
			maker := common.StringToAddress("maker")
			haveToken := common.StringToAddress("have_token")
			wantToken := common.StringToAddress("want_token")
			taker := common.StringToAddress("taker")
			takeAmount := big.NewInt(123)
			giveAmount := big.NewInt(456)
			timestamp := uint64(54321)
			logTake := log_take.LogTakeEntity{
				Id:         oasisLogId,
				Pair:       pair,
				Maker:      maker,
				HaveToken:  haveToken,
				WantToken:  wantToken,
				Taker:      taker,
				TakeAmount: takeAmount,
				GiveAmount: giveAmount,
				Timestamp:  timestamp,
			}

			err = oasisLogRepository.CreateLogTake(logTake, ethLogID)
			Expect(err).ToNot(HaveOccurred())

			var DBethLogID int64
			var DBoasisLogID string
			var DBpair string
			var DBmaker string
			var DBhaveToken string
			var DBwantToken string
			var DBtaker string
			var DBtakeAmount int64
			var DBgiveAmount int64
			var DBtimestamp int64
			err = oasisLogRepository.DB.QueryRowx(
				`SELECT eth_log_id, oasis_log_id, pair, maker, have_token, want_token, taker, take_amount, give_amount, timestamp FROM oasis.log_takes`).
				Scan(&DBethLogID, &DBoasisLogID, &DBpair, &DBmaker, &DBhaveToken, &DBwantToken, &DBtaker, &DBtakeAmount, &DBgiveAmount, &DBtimestamp)
			Expect(err).ToNot(HaveOccurred())
			Expect(DBethLogID).To(Equal(ethLogID))
			Expect(common.Hex2Bytes(DBoasisLogID)).To(Equal(oasisLogId[:]))
			Expect(common.Hex2Bytes(DBpair)).To(Equal(pair[:]))
			Expect(DBmaker).To(Equal(maker.Hex()))
			Expect(DBhaveToken).To(Equal(haveToken.Hex()))
			Expect(DBwantToken).To(Equal(wantToken.Hex()))
			Expect(DBtaker).To(Equal(taker.Hex()))
			Expect(DBtakeAmount).To(Equal(takeAmount.Int64()))
			Expect(DBgiveAmount).To(Equal(giveAmount.Int64()))
			Expect(DBtimestamp).To(Equal(int64(timestamp)))
		})
	})
})
