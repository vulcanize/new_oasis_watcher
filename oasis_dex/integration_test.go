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

package oasis_dex

import (
	"time"

	"database/sql"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/libraries/shared"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
)

var logMake = core.Log{
	BlockNumber: 5497670,
	Address:     constants.ContractAddress,
	TxHash:      "0x21a30773699b014c13e01a7d8fe6253fc31967bcb2901b1cb3bd12961a23ff4b",
	Index:       150,
	Topics: core.Topics{
		constants.LogMakeSignature,
		"0x000000000000000000000000000000000000000000000000000000000000b352",
		"0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		"0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64",
	},
	Data: "0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a232603590000000000000000000000000000000000000000000000022b1c8c1227a0000000000000000000000000000000000000000000000000060cb4106779698f4000000000000000000000000000000000000000000000000000000000005adf32c6",
}

var logTake = core.Log{
	BlockNumber: 5497693,
	Address:     constants.ContractAddress,
	TxHash:      "0xb5328603ef494f0ad3a1728cae0a77fc17dd38e8e0932d89fc62ea18975ec168",
	Index:       52,
	Topics: core.Topics{
		constants.LogTakeSignature,
		"0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		"0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64",
		"0x000000000000000000000000694b206129bea9d2cb734f4fa18278850b228566",
	},
	Data: "0x000000000000000000000000000000000000000000000000000000000000b352000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a23260359000000000000000000000000000000000000000000000000136e4d8e83ef8d9800000000000000000000000000000000000000000000003635c9adc5de9fffdc000000000000000000000000000000000000000000000000000000005adf3417",
}

var logKill = core.Log{
	BlockNumber: 5497708,
	Address:     constants.ContractAddress,
	TxHash:      "0x6220e72531e934fa9d2f871cbded22023532338265e3069c48f6353acfd02e3a",
	Index:       19,
	Topics: core.Topics{
		constants.LogKillSignature,
		"0x000000000000000000000000000000000000000000000000000000000000b352",
		"0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		"0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64",
	},
	Data: "0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a2326035900000000000000000000000000000000000000000000000217ae3e83a3b072680000000000000000000000000000000000000000000005d67e46b9b38aef4024000000000000000000000000000000000000000000000000000000005adf3504",
}

var expectedStateAfterLogMake = TradeStateModel{
	ID:        45906,
	Pair:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
	Guy:       "0x004075e4d4b1ce6c48c81cc940e2bad24b489e64",
	Gem:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	Lot:       "40000000000000000000",
	Gal:       sql.NullString{},
	Pie:       "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359",
	Bid:       "28568534810400000000000",
	Block:     int64(5497670),
	Tx:        "0x21a30773699b014c13e01a7d8fe6253fc31967bcb2901b1cb3bd12961a23ff4b",
	Timestamp: time.Unix(1524576966, 0),
	Deleted:   false,
}

var expectedStateAfterLogTake = TradeStateModel{
	ID:        45906,
	Pair:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
	Guy:       "0x004075e4d4b1ce6c48c81cc940e2bad24b489e64",
	Gem:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	Lot:       "999999999999999999964",
	Gal:       sql.NullString{},
	Pie:       "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359",
	Bid:       "1400141808653012376",
	Block:     int64(5497693),
	Tx:        "0xb5328603ef494f0ad3a1728cae0a77fc17dd38e8e0932d89fc62ea18975ec168",
	Timestamp: time.Unix(1524577303, 0),
	Deleted:   false,
}

var _ = Describe("oasis.state represents current state of offer", func() {
	var db *postgres.DB
	var logRepository repositories.LogRepository
	var watcher shared.Watcher

	BeforeEach(func() {
		var err error

		db, err = postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).ToNot(HaveOccurred())
		logRepository = repositories.LogRepository{DB: db}
		blockchain := &fakes.Blockchain{}

		watcher = shared.Watcher{
			DB:         *db,
			Blockchain: blockchain,
		}
		watcher.AddTransformers(TransformerInitializers())
	})

	AfterEach(func() {
		_, err := db.Exec(`DELETE FROM oasis.offer`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM oasis.trade`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM oasis.kill`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM log_filters`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM logs`)
		Expect(err).ToNot(HaveOccurred())
	})

	It("displays oasis.state for each LogMake added", func() {
		err := logRepository.CreateLogs([]core.Log{logMake})
		Expect(err).ToNot(HaveOccurred())

		watcher.Execute()

		var recordCount int
		err = db.QueryRow(`SELECT COUNT(*) FROM oasis.state`).Scan(&recordCount)
		Expect(err).ToNot(HaveOccurred())
		Expect(recordCount).To(Equal(1))

		type dbRow struct {
			ID_            uint64 `db:"db_id"`
			VulcanizeLogID int64  `db:"vulcanize_log_id"`
			TradeStateModel
		}
		var tradeStates dbRow
		err = db.Get(&tradeStates, `SELECT * FROM oasis.state`)
		Expect(err).ToNot(HaveOccurred())
		Expect(tradeStates.ID).To(Equal(expectedStateAfterLogMake.ID))
		Expect(tradeStates.Pair).To(Equal(expectedStateAfterLogMake.Pair))
		Expect(tradeStates.Guy).To(Equal(expectedStateAfterLogMake.Guy))
		Expect(tradeStates.Gem).To(Equal(expectedStateAfterLogMake.Gem))
		Expect(tradeStates.Pie).To(Equal(expectedStateAfterLogMake.Pie))
		Expect(tradeStates.Lot).To(Equal(expectedStateAfterLogMake.Lot))
		Expect(tradeStates.Bid).To(Equal(expectedStateAfterLogMake.Bid))
		Expect(tradeStates.Deleted).To(BeFalse())
		Expect(tradeStates.Block).To(Equal(expectedStateAfterLogMake.Block))
		Expect(tradeStates.Timestamp.Equal(expectedStateAfterLogMake.Timestamp))
		Expect(tradeStates.Tx).To(Equal(expectedStateAfterLogMake.Tx))
	})

	It("updates oasis.state for each LogTake event received", func() {
		err := logRepository.CreateLogs([]core.Log{logMake, logTake})
		Expect(err).ToNot(HaveOccurred())

		watcher.Execute()

		var recordCount int
		err = db.QueryRow(`SELECT COUNT(*) FROM oasis.state`).Scan(&recordCount)
		Expect(err).ToNot(HaveOccurred())
		Expect(recordCount).To(Equal(1))

		type dbRow struct {
			ID_            uint64 `db:"db_id"`
			VulcanizeLogID int64  `db:"vulcanize_log_id"`
			TradeStateModel
		}
		var tradeStates dbRow
		err = db.Get(&tradeStates, `SELECT * from oasis.state`)
		Expect(err).ToNot(HaveOccurred())

		Expect(tradeStates.ID).To(Equal(expectedStateAfterLogTake.ID))
		Expect(tradeStates.Pair).To(Equal(expectedStateAfterLogTake.Pair))
		Expect(tradeStates.Guy).To(Equal(expectedStateAfterLogTake.Guy))
		Expect(tradeStates.Gem).To(Equal(expectedStateAfterLogTake.Gem))
		Expect(tradeStates.Pie).To(Equal(expectedStateAfterLogTake.Pie))
		Expect(tradeStates.Lot).To(Equal(expectedStateAfterLogTake.Lot))
		Expect(tradeStates.Bid).To(Equal(expectedStateAfterLogTake.Bid))
		Expect(tradeStates.Deleted).To(BeFalse())
		Expect(tradeStates.Block).To(Equal(expectedStateAfterLogTake.Block))
		Expect(tradeStates.Timestamp.Equal(expectedStateAfterLogTake.Timestamp)).To(BeTrue())
		Expect(tradeStates.Tx).To(Equal(expectedStateAfterLogTake.Tx))
	})

	It("shows oasis.state as deleted for each LogKill event received", func() {
		err := logRepository.CreateLogs([]core.Log{logMake, logKill})
		Expect(err).ToNot(HaveOccurred())

		watcher.Execute()

		var count int
		err = db.QueryRow(`SELECT COUNT(*) FROM oasis.state`).Scan(&count)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(Equal(1))

		type dbRow struct {
			DBID           uint64 `db:"db_id"`
			VulcanizeLogID int64  `db:"vulcanize_log_id"`
			TradeStateModel
		}
		var tradeState dbRow
		err = db.Get(&tradeState, `SELECT * FROM oasis.state`)
		Expect(err).ToNot(HaveOccurred())
		Expect(tradeState.ID).To(Equal(int64(45906)))
		Expect(tradeState.Deleted).To(BeTrue())
	})

	Describe("oasis.state reflects the correct state when a log event removed due to reorg", func() {
		It("does not include the oasis.state record when the associated LogMake event is removed", func() {
			err := logRepository.CreateLogs([]core.Log{logMake})
			Expect(err).ToNot(HaveOccurred())
			var logId int64
			err = logRepository.DB.Get(&logId, `SELECT id FROM logs`)
			Expect(err).ToNot(HaveOccurred())

			watcher.Execute()

			var count int
			err = db.QueryRow(`SELECT COUNT(*) FROM oasis.state`).Scan(&count)
			Expect(err).ToNot(HaveOccurred())
			Expect(count).To(Equal(1))

			// remove LogMake event to simulate a reorg
			_, err = db.Exec(`DELETE FROM logs WHERE id = $1`, logId)
			Expect(err).ToNot(HaveOccurred())

			err = db.QueryRow(`SELECT COUNT(*) FROM oasis.state`).Scan(&count)
			Expect(err).ToNot(HaveOccurred())
			Expect(count).To(Equal(0))
		})

		It("rolls back marking an oasis.state record as 'deleted' when the associated LogKill event is removed", func() {
			err := logRepository.CreateLogs([]core.Log{logMake, logKill})
			Expect(err).ToNot(HaveOccurred())
			var logId []int64
			err = logRepository.DB.Select(&logId, `SELECT id FROM logs`)
			Expect(err).ToNot(HaveOccurred())

			watcher.Execute()

			var count int
			err = db.QueryRow(`SELECT COUNT(*) FROM oasis.state`).Scan(&count)
			Expect(err).ToNot(HaveOccurred())
			Expect(count).To(Equal(1))

			type dbRow struct {
				DBID           uint64 `db:"db_id"`
				VulcanizeLogID int64  `db:"vulcanize_log_id"`
				TradeStateModel
			}
			var tradeState dbRow
			err = db.Get(&tradeState, `SELECT * FROM oasis.state`)
			Expect(err).ToNot(HaveOccurred())
			Expect(tradeState.ID).To(Equal(expectedStateAfterLogMake.ID))
			Expect(tradeState.Deleted).To(BeTrue())

			// remove LogKill event to simulate reorg
			_, err = db.Exec(`DELETE FROM logs WHERE id = $1`, logId[1])
			Expect(err).ToNot(HaveOccurred())

			err = db.Get(&tradeState, `SELECT * FROM oasis.state`)
			Expect(err).ToNot(HaveOccurred())
			Expect(tradeState.ID).To(Equal(expectedStateAfterLogMake.ID))
			Expect(tradeState.Deleted).To(BeFalse())
		})

		It("rolls back updates to an oasis.state record when the associated LogTake event is removed", func() {
			err := logRepository.CreateLogs([]core.Log{logMake, logTake})
			Expect(err).ToNot(HaveOccurred())
			var logId []int64
			err = logRepository.DB.Select(&logId, `SELECT id FROM logs`)
			Expect(err).ToNot(HaveOccurred())

			watcher.Execute()

			var recordCount int
			err = db.QueryRow(`SELECT COUNT(*) FROM oasis.state`).Scan(&recordCount)
			Expect(err).ToNot(HaveOccurred())
			Expect(recordCount).To(Equal(1))

			type dbRow struct {
				ID_            uint64 `db:"db_id"`
				VulcanizeLogID int64  `db:"vulcanize_log_id"`
				TradeStateModel
			}
			var tradeStates dbRow
			err = db.Get(&tradeStates, `SELECT * from oasis.state`)
			Expect(err).ToNot(HaveOccurred())
			Expect(tradeStates.ID).To(Equal(expectedStateAfterLogTake.ID))
			Expect(tradeStates.Lot).To(Equal(expectedStateAfterLogTake.Lot))
			Expect(tradeStates.Bid).To(Equal(expectedStateAfterLogTake.Bid))
			Expect(tradeStates.Deleted).To(BeFalse())

			// remove LogTake event to simulate reorg
			_, err = db.Exec(`DELETE FROM logs WHERE id = $1`, logId[1])
			Expect(err).ToNot(HaveOccurred())

			err = db.Get(&tradeStates, `SELECT * from oasis.state`)
			Expect(err).ToNot(HaveOccurred())
			Expect(tradeStates.ID).To(Equal(expectedStateAfterLogMake.ID))
			Expect(tradeStates.Lot).To(Equal(expectedStateAfterLogMake.Lot))
			Expect(tradeStates.Bid).To(Equal(expectedStateAfterLogMake.Bid))
			Expect(tradeStates.Deleted).To(BeFalse())
		})
	})

})
