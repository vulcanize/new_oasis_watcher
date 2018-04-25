package log_take_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/helpers"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

var logTakeEntity = log_take.LogTakeEntity{
	Id:        [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 192},
	Pair:      [32]byte{61, 175, 114, 65, 111, 216, 142, 69, 107, 116, 90, 155, 119, 38, 78, 103, 101, 211, 73, 188, 158, 218, 55, 162, 185, 52, 124, 126, 18, 144, 39, 51},
	Maker:     common.HexToAddress("0x00Ca405026e9018c29c26Cb081DcC9653428bFe9"),
	Pay_gem:   common.HexToAddress("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
	Buy_gem:   common.HexToAddress("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"),
	Taker:     common.HexToAddress("0x0092Ad2b9ae189D50F9cd8E7F4c3355C2c93e3fc"),
	Take_amt:  helpers.BigFromString("34334082741116751270"),
	Give_amt:  helpers.BigFromString("27055257200000000002"),
	Block:     4000870,
	Timestamp: uint64(1499649315),
}

var _ = Describe("LogTake Repository", func() {
	var db *postgres.DB
	var logTakeRepository log_take.Repository
	var logRepository repositories.LogRepository
	var err error
	var vulcanizeLogID int64

	BeforeEach(func() {
		db, err = postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).NotTo(HaveOccurred())
		db.Query(`DELETE FROM logs`)
		db.Query(`DELETE FROM log_filters`)
		db.Query(`DELETE FROM oasis.trade`)

		logRepository = repositories.LogRepository{DB: db}
		err = logRepository.CreateLogs([]core.Log{{}})
		Expect(err).ToNot(HaveOccurred())
		err = logRepository.Get(&vulcanizeLogID, `SELECT id FROM logs`)
		Expect(err).ToNot(HaveOccurred())

		logTakeRepository = log_take.Repository{DB: db}

	})

	Describe("Creating a new LogTake record", func() {
		It("inserts a new log take event", func() {
			ltc := log_take.LogTakeConverter{}
			ltm := ltc.ToModel(logTakeEntity)
			err = logTakeRepository.Create(ltm, vulcanizeLogID)
			Expect(err).ToNot(HaveOccurred())
			type DBRow struct {
				DBID           uint64 `db:"db_id"`
				VulcanizeLogID int64  `db:"vulcanize_log_id"`
				log_take.LogTakeModel
			}
			dbResult := DBRow{}
			err = logTakeRepository.DB.QueryRowx(`SELECT * from oasis.log_take`).StructScan(&dbResult)
			Expect(err).ToNot(HaveOccurred())
			Expect(dbResult.ID).To(Equal(ltm.ID))
			Expect(dbResult.Pair).To(Equal(ltm.Pair))
			Expect(dbResult.Guy).To(Equal(ltm.Guy))
			Expect(dbResult.Gem).To(Equal(ltm.Gem))
			Expect(dbResult.Lot).To(Equal(ltm.Lot))
			Expect(dbResult.Gal).To(Equal(ltm.Gal))
			Expect(dbResult.Pie).To(Equal(ltm.Pie))
			Expect(dbResult.Bid).To(Equal(ltm.Bid))
			Expect(dbResult.Block).To(Equal(ltm.Block))
			Expect(dbResult.Tx).To(Equal(ltm.Tx))
			Expect(dbResult.Timestamp.Equal(ltm.Timestamp)).To(BeTrue())
			Expect(dbResult.VulcanizeLogID).To(Equal(vulcanizeLogID))
		})

		It("doesn't duplicate events that have already been seen", func() {
			ltc := log_take.LogTakeConverter{}
			ltm := ltc.ToModel(logTakeEntity)
			err = logTakeRepository.Create(ltm, vulcanizeLogID)
			Expect(err).ToNot(HaveOccurred())
			err = logTakeRepository.Create(ltm, vulcanizeLogID)
			Expect(err).ToNot(HaveOccurred())
			var count int
			err = logTakeRepository.DB.QueryRowx(`SELECT count(*) from oasis.log_take`).Scan(&count)
			Expect(err).ToNot(HaveOccurred())
			Expect(count).To(Equal(1))
		})

		It("removes a LogTake event when corresponding log is removed", func() {
			ltc := log_take.LogTakeConverter{}
			ltm := ltc.ToModel(logTakeEntity)
			err = logTakeRepository.Create(ltm, vulcanizeLogID)
			Expect(err).ToNot(HaveOccurred())

			var exists bool
			err = logTakeRepository.DB.QueryRowx(`SELECT exists (SELECT * FROM oasis.log_take where vulcanize_log_id = $1)`, vulcanizeLogID).Scan(&exists)
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeTrue())

			//log is removed b/c of reorg
			var logCount int
			_, err = logRepository.DB.Exec(`DELETE FROM logs WHERE id = $1`, vulcanizeLogID)
			Expect(err).ToNot(HaveOccurred())
			err = logRepository.Get(&logCount, `SELECT count(*) FROM logs WHERE id = $1`, vulcanizeLogID)
			Expect(err).ToNot(HaveOccurred())
			Expect(logCount).To(BeZero())

			var logTakeCount int
			err = logTakeRepository.DB.QueryRowx(
				`SELECT count(*) FROM oasis.log_take WHERE vulcanize_log_id = $1`, vulcanizeLogID).Scan(&logTakeCount)
			Expect(err).ToNot(HaveOccurred())
			Expect(logTakeCount).To(BeZero())
		})
	})
})
