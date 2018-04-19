package log_make_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/big"
	"strings"

	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_make"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

var _ = Describe("LogMake Converter", func() {

	var event = types.Log{
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

	It("unpacks a log", func() {
		contract := bind.NewBoundContract(common.HexToAddress(constants.ContractAddress), constants.ABI, nil, nil, nil)
		result := &log_make.LogMakeEntity{}

		err := contract.UnpackLog(result, "LogMake", event)
		Expect(err).ToNot(HaveOccurred())

		Expect(common.ToHex(result.Id[:])).To(Equal("0x000000000000000000000000000000000000000000000000000000000000a291"))
		pair := hexutil.Encode(result.Pair[:])
		Expect(pair).To(Equal("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391"))
		Expect(strings.ToLower(result.Maker.Hex())).To(Equal("0x004075e4d4b1ce6c48c81cc940e2bad24b489e64"))
		Expect(strings.ToLower(result.Pay_gem.Hex())).To(Equal("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"))
		Expect(strings.ToLower(result.Buy_gem.Hex())).To(Equal("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"))
		pay_amount := new(big.Int)
		pay_amount.SetString("40000000000000000000", 10)
		Expect(result.Pay_amt).To(Equal(pay_amount))
		buy_amount := new(big.Int)
		buy_amount.SetString("20561563160000000000000", 10)
		Expect(result.Buy_amt).To(Equal(buy_amount))
		Expect(result.Timestamp).To(Equal(uint64(1523632681)))
	})

	It("converts a watched event into a LogMakeEntity", func() {
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

		result, err := log_make.LogMakeConverter{}.ToEntity(watchedEvent)
		Expect(err).NotTo(HaveOccurred())

		Expect(common.ToHex(result.Id[:])).To(Equal("0x000000000000000000000000000000000000000000000000000000000000a291"))
		pair := hexutil.Encode(result.Pair[:])
		Expect(pair).To(Equal("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391"))
		Expect(strings.ToLower(result.Maker.Hex())).To(Equal("0x004075e4d4b1ce6c48c81cc940e2bad24b489e64"))
		Expect(strings.ToLower(result.Pay_gem.Hex())).To(Equal("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"))
		Expect(strings.ToLower(result.Buy_gem.Hex())).To(Equal("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"))
		payAmount := new(big.Int)
		payAmount.SetString("40000000000000000000", 10)
		Expect(result.Pay_amt).To(Equal(payAmount))
		buyAmount := new(big.Int)
		buyAmount.SetString("20561563160000000000000", 10)
		Expect(result.Buy_amt).To(Equal(buyAmount))
		Expect(result.Block).To(Equal(int64(5433832)))
		Expect(result.TransactionHash).To(Equal("0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c"))
		Expect(result.Timestamp).To(Equal(uint64(1523632681)))
	})

	var _ = Describe("ToEntity from entity to model", func() {
		It("converts an LogMakeEntity to an LogMakeModel", func() {
			pa := new(big.Int)
			pa.SetString("40000000000000000000", 10)
			ba := new(big.Int)
			ba.SetString("20561563160000000000000", 10)
			entity := log_make.LogMakeEntity{
				Id:              [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 162, 145},
				Pair:            [32]byte{157, 212, 129, 16, 220, 196, 68, 253, 194, 66, 81, 12, 9, 187, 187, 226, 26, 89, 117, 202, 192, 97, 216, 47, 123, 132, 59, 206, 6, 27, 163, 145},
				Maker:           common.Address{0, 64, 117, 228, 212, 177, 206, 108, 72, 200, 28, 201, 64, 226, 186, 210, 75, 72, 158, 100},
				Pay_gem:         common.Address{192, 42, 170, 57, 178, 35, 254, 141, 10, 14, 92, 79, 39, 234, 217, 8, 60, 117, 108, 194},
				Buy_gem:         common.Address{137, 210, 74, 107, 76, 203, 27, 111, 170, 38, 37, 254, 86, 43, 221, 154, 35, 38, 3, 89},
				Pay_amt:         pa,
				Buy_amt:         ba,
				Block:           5433832,
				Timestamp:       1523632681,
				TransactionHash: "0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c",
			}

			lmc := log_make.LogMakeConverter{}
			model := lmc.ToModel(entity)

			Expect(model.ID).To(Equal(int64(41617)))
			Expect(model.Pair).To(Equal("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391"))
			Expect(model.Guy).To(Equal("0x004075e4d4b1ce6c48c81cc940e2bad24b489e64"))
			Expect(model.Gem).To(Equal("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"))
			Expect(model.Lot).To(Equal(pa.String()))
			Expect(model.Pie).To(Equal("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"))
			Expect(model.Bid).To(Equal(ba.String()))
			Expect(model.Block).To(Equal(int64(5433832)))
			Expect(model.Timestamp.Equal(time.Unix(1523632681, 0))).To(BeTrue())
			Expect(model.Tx).To(Equal("0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c"))
		})

	})
})
