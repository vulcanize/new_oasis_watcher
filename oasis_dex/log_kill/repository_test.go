package log_kill_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

var _ = Describe("LogKill Repository", func() {
	var db *postgres.DB
	var oasisLogKillRepository log_kill.Datastore

	BeforeEach(func() {
		var err error
		db, err = postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).ToNot(HaveOccurred())
		oasisLogKillRepository = log_kill.Repository{DB: db}
	})

	It("Deletes the offer with the corresponding id", func() {
		lr := repositories.LogRepository{DB: db}
		err := lr.CreateLogs(logs)
		Expect(err).ToNot(HaveOccurred())

		var vulcanizeLogId int64
		err = db.Get(&vulcanizeLogId, `SELECT id FROM public.logs`)
		Expect(err).ToNot(HaveOccurred())

		_, err = db.Exec(`INSERT INTO oasis.offer (id, time, block, tx, vulcanize_log_id) VALUES (1, now(), 1, 1, $1)`, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		var exists bool
		err = db.QueryRow(`SELECT exists (SELECT FROM oasis.offer WHERE id = 1)`).Scan(&exists)
		Expect(err).ToNot(HaveOccurred())
		Expect(exists).To(BeTrue())

		oasisLogKillRepository.Remove(log_kill.LogKillModel{ID: 1})

		err = db.QueryRow(`SELECT exists (SELECT FROM oasis.offer WHERE id = 1)`).Scan(&exists)
		Expect(err).ToNot(HaveOccurred())
		Expect(exists).To(BeFalse())
	})
})
