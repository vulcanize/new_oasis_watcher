package log_kill_test

import (
	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
)

var idOne = "0x000000000000000000000000000000000000000000000000000000000000af21"
var logKill = core.Log{
	BlockNumber: 5488076,
	Address:     constants.ContractAddress,
	TxHash:      "0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad6546ae",
	Index:       0,
	Topics: [4]string{
		constants.LogKillSignature,
		idOne,
		"0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		"0x0000000000000000000000003dc389e0a69d6364a66ab64ebd51234da9569284",
	},
	Data: "0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a23260359000000000000000000000000000000000000000000000000392d2e2bda9c00000000000000000000000000000000000000000000000000927f41fa0a4a418000000000000000000000000000000000000000000000000000000000005adcfebe",
}

var expectedLogKill = log_kill.LogKillModel{
	ID:        44833,
	Pair:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
	Guy:       "0x3dc389e0a69d6364a66ab64ebd51234da9569284",
	Gem:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	Lot:       "4120000000000000000",
	Pie:       "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359",
	Bid:       "2702394520000000000000",
	Block:     5488076,
	Timestamp: time.Unix(1524432574, 0),
	Tx:        "0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad6546ae",
}

//converted logID to assert against
var logs = []core.Log{
	logKill,
	{
		BlockNumber: 0,
		TxHash:      "",
		Address:     "",
		Topics:      core.Topics{},
		Index:       0,
		Data:        "",
	},
}

var _ = Describe("Integration test with vulcanizedb", func() {
	var db *postgres.DB

	BeforeEach(func() {
		var err error
		db, err = postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).NotTo(HaveOccurred())
		lr := repositories.LogRepository{DB: db}
		err = lr.CreateLogs(logs)
		Expect(err).ToNot(HaveOccurred())

		var vulcanizeLogIds []int64
		err = db.Select(&vulcanizeLogIds, `SELECT id FROM public.logs`)
		Expect(err).ToNot(HaveOccurred())

	})

	AfterEach(func() {
		_, err := db.Exec(`DELETE FROM oasis.kill`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM log_filters`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM logs`)
		Expect(err).ToNot(HaveOccurred())
	})

	It("creates oasis.kill for each LogKill event received", func() {
		blockchain := &fakes.Blockchain{}
		transformer := log_kill.NewTransformer(db, blockchain)

		transformer.Execute()

		var count int
		err := db.QueryRow(`SELECT COUNT(*) FROM oasis.kill`).Scan(&count)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(Equal(1))

		type dbRow struct {
			DBID           uint64 `db:"db_id"`
			VulcanizeLogID int64  `db:"vulcanize_log_id"`
			log_kill.LogKillModel
		}
		var logKill dbRow
		err = db.Get(&logKill, `SELECT * FROM oasis.kill WHERE block=$1`, logs[0].BlockNumber)
		Expect(err).ToNot(HaveOccurred())
		Expect(logKill.ID).To(Equal(expectedLogKill.ID))
		Expect(logKill.Pair).To(Equal(expectedLogKill.Pair))
		Expect(logKill.Guy).To(Equal(expectedLogKill.Guy))
		Expect(logKill.Gem).To(Equal(expectedLogKill.Gem))
		Expect(logKill.Lot).To(Equal(expectedLogKill.Lot))
		Expect(logKill.Pie).To(Equal(expectedLogKill.Pie))
		Expect(logKill.Bid).To(Equal(expectedLogKill.Bid))
		Expect(logKill.Block).To(Equal(expectedLogKill.Block))
		Expect(logKill.Tx).To(Equal(expectedLogKill.Tx))
		Expect(logKill.Timestamp.Equal(expectedLogKill.Timestamp)).To(BeTrue())
	})

})
