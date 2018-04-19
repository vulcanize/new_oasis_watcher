package log_kill_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
)

var idOne = "0x0000000000000000000000000000000000000000000000000000000000009eda"
var logKill = core.Log{
	BlockNumber: 5428074,
	Address:     constants.ContractAddress,
	TxHash:      "0x769de518d62d3ec4c4c5b50c51ca8248f27f4f5f833f349fc150adc4b2548cfd",
	Index:       0,
	Topics: [4]string{
		constants.LogKillSignature,
		idOne,
		"0x204053929a0ef66ee09fa2295cc078531ee0339fa4d3e02ce9bb3f1a5d0116dd",
		"0x0000000000000000000000009f87bda86354ba26d0e9250d006876d8b5216622",
	},
	Data: "000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a200000000000000000000000000000000000000000000000003782dace9d90000000000000000000000000000000000000000000000000000028d1286abf261e2000000000000000000000000000000000000000000000000000000005acf7f72",
}

//converted logID to assert against
var otherLogId int64 = 40000
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

		//LogKill requires some offers be present
		_, err = db.Exec(`INSERT INTO oasis.offer (id, time, vulcanize_log_id, block, tx) VALUES 
			($1, now(), $2, 1, 1),
			($3, now(), $4, 1, 1)`,
			common.HexToHash(idOne).Big().Uint64(), vulcanizeLogIds[0],
			otherLogId, vulcanizeLogIds[1],
		)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		_, err := db.Exec(`DELETE FROM oasis.offer`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM log_filters`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM logs`)
		Expect(err).ToNot(HaveOccurred())
	})

	It("Deletes oasis.offer with corresponding id for each LogKill event received", func() {
		blockchain := &fakes.Blockchain{}
		transformer := log_kill.NewTransformer(db, blockchain)

		transformer.Execute()

		var id int64
		err := db.QueryRow(`SELECT id FROM oasis.offer`).Scan(&id)

		Expect(err).ToNot(HaveOccurred())
		Expect(id).To(Equal(int64(otherLogId)))
	})

})
