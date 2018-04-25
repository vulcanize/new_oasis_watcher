package log_make_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/helpers"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_make"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

var _ = Describe("LogMake Converter", func() {
	var eventLog = types.Log{
		Address: common.HexToAddress(constants.ContractAddress),
		Topics: []common.Hash{common.HexToHash(constants.LogMakeSignature),
			common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000a291"),
			common.HexToHash("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391"),
			common.HexToHash("0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64")},
		Data:        hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a232603590000000000000000000000000000000000000000000000022b1c8c1227a0000000000000000000000000000000000000000000000000045aa502b2307e598000000000000000000000000000000000000000000000000000000000005ad0ca29"),
		BlockNumber: 5433832,
		TxHash:      common.HexToHash("0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c"),
		TxIndex:     50,
		BlockHash:   common.HexToHash("0x32f8b12023b3a1b4c73f9a46da976931b0355714ada8b8044ebcb2cd295751a9"),
		Index:       50,
		Removed:     false,
	}

	watchedEvent := core.WatchedEvent{
		LogID:       0,
		Name:        "",
		BlockNumber: 5433832,
		Address:     constants.ContractAddress,
		TxHash:      "0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c",
		Index:       50,
		Topic0:      constants.LogMakeSignature,
		Topic1:      "0x000000000000000000000000000000000000000000000000000000000000a291",
		Topic2:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		Topic3:      "0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64",
		Data:        "0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a232603590000000000000000000000000000000000000000000000022b1c8c1227a0000000000000000000000000000000000000000000000000045aa502b2307e598000000000000000000000000000000000000000000000000000000000005ad0ca29",
	}

	var expectedEntity = log_make.LogMakeEntity{
		Id:              [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 162, 145},
		Pair:            [32]byte{157, 212, 129, 16, 220, 196, 68, 253, 194, 66, 81, 12, 9, 187, 187, 226, 26, 89, 117, 202, 192, 97, 216, 47, 123, 132, 59, 206, 6, 27, 163, 145},
		Maker:           common.Address{0, 64, 117, 228, 212, 177, 206, 108, 72, 200, 28, 201, 64, 226, 186, 210, 75, 72, 158, 100},
		Pay_gem:         common.Address{192, 42, 170, 57, 178, 35, 254, 141, 10, 14, 92, 79, 39, 234, 217, 8, 60, 117, 108, 194},
		Buy_gem:         common.Address{137, 210, 74, 107, 76, 203, 27, 111, 170, 38, 37, 254, 86, 43, 221, 154, 35, 38, 3, 89},
		Pay_amt:         helpers.BigFromString("40000000000000000000"),
		Buy_amt:         helpers.BigFromString("20561563160000000000000"),
		Block:           5433832,
		Timestamp:       1523632681,
		TransactionHash: "0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c",
	}

	var expectedModel = log_make.LogMakeModel{
		ID:        41617,
		Pair:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		Guy:       "0x004075e4d4b1ce6c48c81cc940e2bad24b489e64",
		Gem:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		Lot:       "40000000000000000000",
		Pie:       "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359",
		Bid:       "20561563160000000000000",
		Block:     5433832,
		Timestamp: time.Unix(1523632681, 0),
		Tx:        "0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c",
	}

	It("unpacks a log", func() {
		contract := bind.NewBoundContract(common.HexToAddress(constants.ContractAddress), constants.ABI, nil, nil, nil)
		logMakeEntity := &log_make.LogMakeEntity{}

		err := contract.UnpackLog(logMakeEntity, "LogMake", eventLog)
		Expect(err).ToNot(HaveOccurred())

		Expect(logMakeEntity.Id).To(Equal(expectedEntity.Id))
		Expect(logMakeEntity.Pair).To(Equal(expectedEntity.Pair))
		Expect(logMakeEntity.Maker).To(Equal(expectedEntity.Maker))
		Expect(logMakeEntity.Pay_gem).To(Equal(expectedEntity.Pay_gem))
		Expect(logMakeEntity.Buy_gem).To(Equal(expectedEntity.Buy_gem))
		Expect(logMakeEntity.Pay_amt).To(Equal(expectedEntity.Pay_amt))
		Expect(logMakeEntity.Buy_amt).To(Equal(expectedEntity.Buy_amt))
		Expect(logMakeEntity.Timestamp).To(Equal(expectedEntity.Timestamp))
	})

	It("converts a watched event into a LogMakeEntity", func() {
		convertedEvent, err := log_make.LogMakeConverter{}.ToEntity(watchedEvent)
		Expect(err).NotTo(HaveOccurred())

		Expect(convertedEvent.Id).To(Equal(expectedEntity.Id))
		Expect(convertedEvent.Pair).To(Equal(expectedEntity.Pair))
		Expect(convertedEvent.Maker).To(Equal(expectedEntity.Maker))
		Expect(convertedEvent.Pay_gem).To(Equal(expectedEntity.Pay_gem))
		Expect(convertedEvent.Buy_gem).To(Equal(expectedEntity.Buy_gem))
		Expect(convertedEvent.Pay_amt).To(Equal(expectedEntity.Pay_amt))
		Expect(convertedEvent.Buy_amt).To(Equal(expectedEntity.Buy_amt))
		Expect(convertedEvent.Block).To(Equal(expectedEntity.Block))
		Expect(convertedEvent.TransactionHash).To(Equal(expectedEntity.TransactionHash))
		Expect(convertedEvent.Timestamp).To(Equal(expectedEntity.Timestamp))
	})

	var _ = Describe("ToEntity from entity to model", func() {
		It("converts an LogMakeEntity to an LogMakeModel", func() {
			lmc := log_make.LogMakeConverter{}
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
})
