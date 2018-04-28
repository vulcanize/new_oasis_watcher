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

package log_make_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/log_make"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

var _ = Describe("LogMake Repository", func() {
	var db *postgres.DB
	var repository log_make.Repository
	var logRepository repositories.LogRepository
	var vulcanizeLogId int64

	var logMakeEntity = log_make.LogMakeEntity{
		Id:        [32]byte{1, 2, 3, 4},
		Pair:      [32]byte{5, 6, 7, 8},
		Maker:     common.HexToAddress("maker"),
		Pay_gem:   common.HexToAddress("pay_gem"),
		Buy_gem:   common.HexToAddress("buy_gem"),
		Pay_amt:   big.NewInt(123),
		Buy_amt:   big.NewInt(123),
		Block:     int64(12345),
		Timestamp: uint64(1),
	}

	BeforeEach(func() {
		var err error
		db, err = postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).NotTo(HaveOccurred())
		db.Query(`DELETE FROM logs`)
		db.Query(`DELETE FROM log_filters`)
		db.Query(`DELETE FROM oasis.offer`)

		logRepository = repositories.LogRepository{DB: db}
		err = logRepository.CreateLogs([]core.Log{{}})
		Expect(err).ToNot(HaveOccurred())

		err = logRepository.Get(&vulcanizeLogId, `SELECT id FROM logs`)
		Expect(err).ToNot(HaveOccurred())

		repository = log_make.Repository{DB: db}
	})

	AfterEach(func() {
		repository.Exec(`DELETE FROM oasis.offer`)
		repository.Exec(`DELETE FROM oasis.log_make`)
	})

	It("Creates a new LogMake record", func() {
		lmc := log_make.LogMakeConverter{}
		model := lmc.ToModel(logMakeEntity)
		err := repository.Create(model, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		type DBRow struct {
			DBID           uint64 `db:"db_id"`
			VulcanizeLogID int64  `db:"vulcanize_log_id"`
			log_make.LogMakeModel
		}
		dbResult := DBRow{}

		err = repository.DB.QueryRowx(`SELECT * FROM oasis.log_make`).StructScan(&dbResult)
		Expect(err).ToNot(HaveOccurred())

		Expect(dbResult.VulcanizeLogID).To(Equal(vulcanizeLogId))
		Expect(dbResult.Pair).To(Equal(common.ToHex(logMakeEntity.Pair[:])))
		Expect(dbResult.Gem).To(Equal(logMakeEntity.Pay_gem.Hex()))
		Expect(dbResult.Block).To(Equal(logMakeEntity.Block))
		Expect(dbResult.Lot).To(Equal(logMakeEntity.Pay_amt.String()))
		Expect(dbResult.Pie).To(Equal(logMakeEntity.Buy_gem.Hex()))
		Expect(dbResult.Bid).To(Equal(logMakeEntity.Buy_amt.String()))
		Expect(dbResult.Guy).To(Equal(logMakeEntity.Maker.Hex()))
		convertedLogMakeTime := time.Unix(int64(logMakeEntity.Timestamp), 0)
		Expect(dbResult.Timestamp.Equal(convertedLogMakeTime)).To(BeTrue())
		Expect(dbResult.Tx).To(Equal(logMakeEntity.TransactionHash))
	})

	It("does not duplicate LogMakes that have already been seen", func() {
		lmc := log_make.LogMakeConverter{}
		model := lmc.ToModel(logMakeEntity)

		err := repository.Create(model, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		err = repository.Create(model, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())

		var count int
		err = repository.DB.QueryRowx(`SELECT count(*) FROM oasis.log_make`).Scan(&count)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(Equal(1))
	})

	It("removes a LogMake event when corresponding log is removed", func() {
		ltc := log_make.LogMakeConverter{}
		ltm := ltc.ToModel(logMakeEntity)
		err := repository.Create(ltm, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())

		var exists bool
		err = repository.DB.QueryRowx(`SELECT exists (SELECT * FROM oasis.log_make where vulcanize_log_id = $1)`, vulcanizeLogId).Scan(&exists)
		Expect(err).ToNot(HaveOccurred())
		Expect(exists).To(BeTrue())

		//log is removed b/c of reorg
		var logCount int
		_, err = logRepository.DB.Exec(`DELETE FROM logs WHERE id = $1`, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		err = logRepository.Get(&logCount, `SELECT count(*) FROM logs WHERE id = $1`, vulcanizeLogId)
		Expect(err).ToNot(HaveOccurred())
		Expect(logCount).To(BeZero())

		var logTakeCount int
		err = repository.DB.QueryRowx(
			`SELECT count(*) FROM oasis.log_take WHERE vulcanize_log_id = $1`, vulcanizeLogId).Scan(&logTakeCount)
		Expect(err).ToNot(HaveOccurred())
		Expect(logTakeCount).To(BeZero())
	})
})
