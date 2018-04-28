// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log_kill_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/helpers"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

var logKillEntity = log_kill.LogKillEntity{
	Id:      [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 175, 33},
	Pair:    [32]byte{157, 212, 129, 16, 220, 196, 68, 253, 194, 66, 81, 12, 9, 187, 187, 226, 26, 89, 117, 202, 192, 97, 216, 47, 123, 132, 59, 206, 6, 27, 163, 145},
	Maker:   common.Address{61, 195, 137, 224, 166, 157, 99, 100, 166, 106, 182, 78, 189, 81, 35, 77, 169, 86, 146, 132},
	Pay_gem: common.Address{192, 42, 170, 57, 178, 35, 254, 141, 10, 14, 92, 79, 39, 234, 217, 8, 60, 117, 108, 194},
	Buy_gem: common.Address{137, 210, 74, 107, 76, 203, 27, 111, 170, 38, 37, 254, 86, 43, 221, 154, 35, 38, 3, 89},
	Pay_amt: helpers.BigFromString("4120000000000000000"),
	Buy_amt: helpers.BigFromString("2702394520000000000000"),
}

var _ = Describe("LogKill Repository", func() {
	var db *postgres.DB
	var repository log_kill.Repository
	var logRepository repositories.LogRepository
	var vulcanizeLogId int64

	BeforeEach(func() {
		var err error
		db, err = postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).NotTo(HaveOccurred())

		logRepository = repositories.LogRepository{DB: db}
		err = logRepository.CreateLogs([]core.Log{{}})
		Expect(err).ToNot(HaveOccurred())

		err = logRepository.Get(&vulcanizeLogId, `SELECT id FROM logs`)
		Expect(err).ToNot(HaveOccurred())

		repository = log_kill.Repository{DB: db}

	})

	AfterEach(func() {
		db.Query(`DELETE FROM logs`)
		db.Query(`DELETE FROM log_filters`)
		db.Query(`DELETE FROM oasis.kill`)

		repository.DB.Exec(`DELETE FROM oasis.kill`)
	})

	It("Creates a new LogKill record", func() {
		converter := log_kill.LogKillConverter{}
		model := converter.ToModel(logKillEntity)
		err := repository.Create(model, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		type DBRow struct {
			DBID           uint64 `db:"db_id"`
			VulcanizeLogID int64  `db:"vulcanize_log_id"`
			log_kill.LogKillModel
		}
		dbResult := DBRow{}

		err = repository.QueryRowx(`SELECT * FROM oasis.kill`).StructScan(&dbResult)
		Expect(err).ToNot(HaveOccurred())

		Expect(dbResult.VulcanizeLogID).To(Equal(vulcanizeLogId))
		Expect(dbResult.Pair).To(Equal(model.Pair))
		Expect(dbResult.Gem).To(Equal(model.Gem))
		Expect(dbResult.Block).To(Equal(model.Block))
		Expect(dbResult.Lot).To(Equal(model.Lot))
		Expect(dbResult.Pie).To(Equal(model.Pie))
		Expect(dbResult.Bid).To(Equal(model.Bid))
		Expect(dbResult.Guy).To(Equal(model.Guy))
		Expect(dbResult.Timestamp.Equal(model.Timestamp)).To(BeTrue())
		Expect(dbResult.Tx).To(Equal(model.Tx))
	})

	It("does not duplicate LogKills that have already been seen", func() {
		converter := log_kill.LogKillConverter{}
		model := converter.ToModel(logKillEntity)

		err := repository.Create(model, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		err = repository.Create(model, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())

		var count int
		err = repository.DB.QueryRowx(`SELECT count(*) FROM oasis.kill`).Scan(&count)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(Equal(1))
	})

	It("Removes an LogKill record when the corresponding log is removed", func() {
		var exists bool

		converter := log_kill.LogKillConverter{}
		model := converter.ToModel(logKillEntity)
		err := repository.Create(model, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())

		err = repository.DB.QueryRow(`SELECT exists (SELECT * FROM oasis.kill WHERE vulcanize_log_id = $1)`, vulcanizeLogId).Scan(&exists)
		Expect(err).ToNot(HaveOccurred())
		Expect(exists).To(BeTrue())

		var logCount int
		_, err = logRepository.DB.Exec(`DELETE FROM logs WHERE id = $1`, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		err = logRepository.Get(&logCount, `SELECT count(*) FROM logs WHERE id = $1`, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		Expect(logCount).To(BeZero())

		var LogKillCount int
		err = repository.DB.QueryRowx(
			`SELECT count(*) FROM oasis.kill WHERE vulcanize_log_id = $1`, vulcanizeLogId).Scan(&LogKillCount)
		Expect(err).ToNot(HaveOccurred())
		Expect(LogKillCount).To(BeZero())
	})
})
