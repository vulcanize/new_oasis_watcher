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

package log_take_test

import (
	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
)

var logMake = core.Log{
	BlockNumber: 5418392,
	Address:     constants.ContractAddress,
	TxHash:      "0xc11979bda618f244c89ad82f4941d1d308bfc264eee2dab623ffeae4b6039d6e",
	Topics: core.Topics{
		constants.LogMakeSignature,
		"0x0000000000000000000000000000000000000000000000000000000000009bef",
		"0x204053929a0ef66ee09fa2295cc078531ee0339fa4d3e02ce9bb3f1a5d0116dd",
		"0x000000000000000000000000168910909606a2fca90d4c28fa39b50407b9c526"},
	Index: 22,
	Data:  "0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a200000000000000000000000000000000000000000000000173ecba01ec780000000000000000000000000000000000000000000000000001158e460913d00000000000000000000000000000000000000000000000000000000000005acd5ada",
}

var logTake1 = core.Log{
	BlockNumber: 5430136,
	Address:     constants.ContractAddress,
	TxHash:      "0xaca917ef9440aaf2d37cd36309872ce8ab6251f56cac62524b6eb63d5c891be8",
	Index:       8,
	Topics: core.Topics{
		constants.LogTakeSignature,
		"0x204053929a0ef66ee09fa2295cc078531ee0339fa4d3e02ce9bb3f1a5d0116dd",
		"0x000000000000000000000000168910909606a2fca90d4c28fa39b50407b9c526",
		"0x0000000000000000000000000016bd4cb70bd98ca07a341da66450b5d22a55aa",
	},
	Data: "0x0000000000000000000000000000000000000000000000000000000000009bef000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a200000000000000000000000000000000000000000000000003ac43f17ce74b6f00000000000000000000000000000000000000000000000002bdb0cb23e5ebe0000000000000000000000000000000000000000000000000000000005acff7c4",
}

var logTake2 = core.Log{
	BlockNumber: 5430139,
	Address:     constants.ContractAddress,
	TxHash:      "0x764bd5e6127d263140b7835920cbc3cb28ca67ce62b73a04aff569e1fa75423c",
	Index:       54,
	Topics: core.Topics{
		constants.LogTakeSignature,
		"0x204053929a0ef66ee09fa2295cc078531ee0339fa4d3e02ce9bb3f1a5d0116dd",
		"0x000000000000000000000000168910909606a2fca90d4c28fa39b50407b9c526",
		"0x0000000000000000000000000016bd4cb70bd98ca07a341da66450b5d22a55aa",
	},
	Data: "0x0000000000000000000000000000000000000000000000000000000000009bef000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a200000000000000000000000000000000000000000000000003ac43f17ce74b6f00000000000000000000000000000000000000000000000002bdb0cb23e5ebe0000000000000000000000000000000000000000000000000000000005acff7dd",
}

var otherLog = core.Log{
	BlockNumber: 0,
	TxHash:      "0xHASH",
	Address:     "0xADDRESS",
	Topics:      core.Topics{},
	Index:       0,
	Data:        "0xDATA",
}

var logs = []core.Log{logMake, logTake1, logTake2, otherLog}

var expectedLogTake1Model = log_take.LogTakeModel{
	ID:        39919,
	Pair:      "0x204053929a0ef66ee09fa2295cc078531ee0339fa4d3e02ce9bb3f1a5d0116dd",
	Guy:       "0x168910909606a2fca90d4c28fa39b50407b9c526",
	Gem:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	Lot:       "197508345201290208",
	Gal:       "0x0016bd4cb70bd98ca07a341da66450b5d22a55aa",
	Pie:       "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2",
	Bid:       "264661182569728879",
	Block:     5430136,
	Tx:        "0xaca917ef9440aaf2d37cd36309872ce8ab6251f56cac62524b6eb63d5c891be8",
	Timestamp: time.Unix(1523578820, 0),
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
	})

	AfterEach(func() {
		_, err := db.Exec(`DELETE FROM oasis.trade`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM log_filters`)
		Expect(err).ToNot(HaveOccurred())
		_, err = db.Exec(`DELETE FROM logs`)
		Expect(err).ToNot(HaveOccurred())
	})

	It("creates oasis.log_take for each LogTake event received", func() {
		blockchain := &fakes.Blockchain{}
		transformer := log_take.NewTransformer(db, blockchain)

		transformer.Execute()

		var count int
		err := db.QueryRow(`SELECT COUNT(*) FROM oasis.log_take`).Scan(&count)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(Equal(2))

		type dbRow struct {
			DBID           uint64 `db:"db_id"`
			VulcanizeLogID int64  `db:"vulcanize_log_id"`
			log_take.LogTakeModel
		}
		var logTake dbRow
		err = db.Get(&logTake, `SELECT * FROM oasis.log_take WHERE block=$1`, logs[1].BlockNumber)

		Expect(err).ToNot(HaveOccurred())
		Expect(logTake.ID).To(Equal(expectedLogTake1Model.ID))
		Expect(logTake.Pair).To(Equal(expectedLogTake1Model.Pair))
		Expect(logTake.Guy).To(Equal(expectedLogTake1Model.Guy))
		Expect(logTake.Gem).To(Equal(expectedLogTake1Model.Gem))
		Expect(logTake.Lot).To(Equal(expectedLogTake1Model.Lot))
		Expect(logTake.Gal).To(Equal(expectedLogTake1Model.Gal))
		Expect(logTake.Pie).To(Equal(expectedLogTake1Model.Pie))
		Expect(logTake.Bid).To(Equal(expectedLogTake1Model.Bid))
		Expect(logTake.Tx).To(Equal(expectedLogTake1Model.Tx))
		Expect(logTake.Block).To(Equal(expectedLogTake1Model.Block))
		Expect(logTake.Timestamp.Equal(expectedLogTake1Model.Timestamp)).To(BeTrue())
	})

})
