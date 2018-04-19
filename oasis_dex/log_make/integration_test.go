package log_make_test

import (
	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_make"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
)

var logMake = core.Log{
	BlockNumber: 5433832,
	Address:     constants.ContractAddress,
	TxHash:      "0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c",
	Index:       31,
	Topics: core.Topics{
		constants.LogMakeSignature,
		"0x000000000000000000000000000000000000000000000000000000000000a291",
		"0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		"0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64",
	},
	Data: "0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a232603590000000000000000000000000000000000000000000000022b1c8c1227a0000000000000000000000000000000000000000000000000045aa502b2307e598000000000000000000000000000000000000000000000000000000000005ad0ca29",
}

var logs = []core.Log{
	logMake,
	{
		BlockNumber: 0,
		TxHash:      "0xHASH",
		Address:     "0xADDRESS",
		Topics:      core.Topics{},
		Index:       0,
		Data:        "0xDATA",
	},
}

var _ = Describe("Integration test with VulcanizeDB", func() {
	var db *postgres.DB

	BeforeEach(func() {
		var err error

		db, err = postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).ToNot(HaveOccurred())
		logRepository := repositories.LogRepository{DB: db}
		err = logRepository.CreateLogs(logs)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		_, err := db.Exec(`DELETE FROM oasis.trade`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM log_filters`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM logs`)
		Expect(err).ToNot(HaveOccurred())
	})

	It("Creates an oasis.offer for each LogTake event received", func() {
		blockchain := &fakes.Blockchain{}
		transformer := log_make.NewTransformer(db, blockchain)

		transformer.Execute()

		var offerCount int
		err := db.QueryRow(`SELECT COUNT(*) FROM oasis.offer where block = $1`, logs[0].BlockNumber).Scan(&offerCount)
		Expect(err).ToNot(HaveOccurred())
		Expect(offerCount).To(Equal(1))

		type dbRow struct {
			ID_            uint64 `db:"db_id"`
			VulcanizeLogID int64  `db:"vulcanize_log_id"`
			log_make.LogMakeModel
		}
		var offer dbRow
		err = db.Get(&offer, `SELECT * from oasis.offer where block = $1`, logs[0].BlockNumber)
		Expect(err).ToNot(HaveOccurred())
		Expect(offer.ID).To(Equal(int64(41617)))
		Expect(offer.Pair).To(Equal("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391"))
		Expect(offer.Guy).To(Equal("0x004075e4d4b1ce6c48c81cc940e2bad24b489e64"))
		Expect(offer.Gem).To(Equal("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"))
		Expect(offer.Pie).To(Equal("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"))
		Expect(offer.Lot).To(Equal("40000000000000000000"))
		Expect(offer.Bid).To(Equal("20561563160000000000000"))
		Expect(offer.Block).To(Equal(int64(5433832)))
		Expect(offer.Timestamp.Equal(time.Unix(1523632681, 0))).To(BeTrue())
		Expect(offer.Tx).To(Equal("0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c"))

	})

})
