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
	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/helpers"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

var expectedModel = log_kill.LogKillModel{
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

var expectedEntity = log_kill.LogKillEntity{
	Id:              [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 175, 33},
	Pair:            [32]byte{157, 212, 129, 16, 220, 196, 68, 253, 194, 66, 81, 12, 9, 187, 187, 226, 26, 89, 117, 202, 192, 97, 216, 47, 123, 132, 59, 206, 6, 27, 163, 145},
	Maker:           common.Address{61, 195, 137, 224, 166, 157, 99, 100, 166, 106, 182, 78, 189, 81, 35, 77, 169, 86, 146, 132},
	Pay_gem:         common.Address{192, 42, 170, 57, 178, 35, 254, 141, 10, 14, 92, 79, 39, 234, 217, 8, 60, 117, 108, 194},
	Buy_gem:         common.Address{137, 210, 74, 107, 76, 203, 27, 111, 170, 38, 37, 254, 86, 43, 221, 154, 35, 38, 3, 89},
	Pay_amt:         helpers.BigFromString("4120000000000000000"),
	Buy_amt:         helpers.BigFromString("2702394520000000000000"),
	Block:           5488076,
	Timestamp:       1524432574,
	TransactionHash: "0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad6546ae",
}

var watchedEvent = core.WatchedEvent{
	LogID:       1,
	Name:        "",
	BlockNumber: 5488076,
	Address:     constants.ContractAddress,
	TxHash:      "0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad6546ae",
	Index:       110,
	Topic0:      constants.LogKillSignature,
	Topic1:      "0x000000000000000000000000000000000000000000000000000000000000af21",
	Topic2:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
	Topic3:      "0x0000000000000000000000003dc389e0a69d6364a66ab64ebd51234da9569284",
	Data:        "0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a23260359000000000000000000000000000000000000000000000000392d2e2bda9c00000000000000000000000000000000000000000000000000927f41fa0a4a418000000000000000000000000000000000000000000000000000000000005adcfebe",
}

var _ = Describe("LogKill Converter", func() {
	It("converts a watched event into a LogKillEntity", func() {

		result, err := log_kill.LogKillConverter{}.ToEntity(watchedEvent)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Id).To(Equal(expectedEntity.Id))
		Expect(result.Pair).To(Equal(expectedEntity.Pair))
		Expect(result.Maker).To(Equal(expectedEntity.Maker))
		Expect(result.Pay_gem).To(Equal(expectedEntity.Pay_gem))
		Expect(result.Buy_gem).To(Equal(expectedEntity.Buy_gem))
		Expect(result.Pay_amt).To(Equal(expectedEntity.Pay_amt))
		Expect(result.Buy_amt).To(Equal(expectedEntity.Buy_amt))
		Expect(result.Block).To(Equal(expectedEntity.Block))
		Expect(result.Timestamp).To(Equal(expectedEntity.Timestamp))
		Expect(result.TransactionHash).To(Equal(expectedEntity.TransactionHash))
	})

	It("converts a LogKillEntity to an LogKillModel", func() {

		lmc := log_kill.LogKillConverter{}
		model := lmc.ToModel(expectedEntity)

		Expect(model.ID).To(Equal(expectedModel.ID))
		Expect(model.Pair).To(Equal(expectedModel.Pair))
		Expect(model.Guy).To(Equal(expectedModel.Guy))
		Expect(model.Gem).To(Equal(expectedModel.Gem))
		Expect(model.Lot).To(Equal(expectedModel.Lot))
		Expect(model.Pie).To(Equal(expectedModel.Pie))
		Expect(model.Bid).To(Equal(expectedModel.Bid))
		Expect(model.Block).To(Equal(expectedModel.Block))
		Expect(model.Timestamp.Equal(expectedModel.Timestamp)).To(BeTrue())
		Expect(model.Tx).To(Equal(expectedModel.Tx))
	})

})
