package log_take_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/big"

	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

var _ = Describe("Log Take Converter", func() {

	blocknumber := int64(4000870)
	var event = types.Log{
		Address: common.HexToAddress("0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f"),
		Topics: []common.Hash{common.HexToHash(constants.LogTakeSignature),
			common.HexToHash("0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733"),
			common.HexToHash("0x00000000000000000000000000ca405026e9018c29c26cb081dcc9653428bfe9"),
			common.HexToHash("0x0000000000000000000000000092ad2b9ae189d50f9cd8e7f4c3355c2c93e3fc")},
		Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000009c0000000000000000000000000c66ea802717bfb9833400264dd12c2bceaa34a6d000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7000000000000000000000000000000000000000000000001dc7b2a57e1224da60000000000000000000000000000000000000000000000017777951418d46002000000000000000000000000000000000000000000000000000000005962d523"),
		BlockNumber: 4000870,
		TxHash:      common.HexToHash("0x98237ddc11a618f5546cd3098e57d9ba159418cb18851fb98130cb3114063807"),
		TxIndex:     50,
		BlockHash:   common.HexToHash("0xca9d5f4507b83030e2cf07534429ab5640b11c18fb38edf22674d9fcdb692cf2"),
		Index:       50,
		Removed:     false,
	}

	It("unpacks a log", func() {
		contract := bind.NewBoundContract(common.HexToAddress("0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f"), constants.ABI, nil, nil, nil)
		result := &log_take.LogTakeEntity{}

		err := contract.UnpackLog(result, "LogTake", event)
		Expect(err).ToNot(HaveOccurred())

		p := hexutil.Encode(result.Pair[:])
		Expect(p).To(Equal("0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733"))
		Expect(result.Maker.Hex()).To(Equal("0x00Ca405026e9018c29c26Cb081DcC9653428bFe9"))
		Expect(result.Pay_gem.Hex()).To(Equal("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"))
		Expect(result.Buy_gem.Hex()).To(Equal("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"))
		Expect(result.Taker.Hex()).To(Equal("0x0092Ad2b9ae189D50F9cd8E7F4c3355C2c93e3fc"))
		ta := new(big.Int)
		ta.SetString("34334082741116751270", 10)
		Expect(result.Take_amt).To(Equal(ta))
		ga := new(big.Int)
		ga.SetString("27055257200000000002", 10)
		Expect(result.Give_amt).To(Equal(ga))
		Expect(result.Timestamp).To(Equal(uint64(1499649315)))
	})

	It("converts watched event from address 0x83 into LogTakeEntity struct", func() {
		watchedEvent := core.WatchedEvent{
			LogID:       0,
			Name:        "",
			BlockNumber: blocknumber,
			Address:     "0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f",
			TxHash:      "0x98237ddc11a618f5546cd3098e57d9ba159418cb18851fb98130cb3114063807",
			Index:       50,
			Topic0:      constants.LogTakeSignature,
			Topic1:      "0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733",
			Topic2:      "0x00000000000000000000000000ca405026e9018c29c26cb081dcc9653428bfe9",
			Topic3:      "0x0000000000000000000000000092ad2b9ae189d50f9cd8e7f4c3355c2c93e3fc",
			Data:        "0x00000000000000000000000000000000000000000000000000000000000009c0000000000000000000000000c66ea802717bfb9833400264dd12c2bceaa34a6d000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7000000000000000000000000000000000000000000000001dc7b2a57e1224da60000000000000000000000000000000000000000000000017777951418d46002000000000000000000000000000000000000000000000000000000005962d523",
		}

		result, err := log_take.LogTakeConverter{}.ToEntity(watchedEvent)
		Expect(err).NotTo(HaveOccurred())

		pair := hexutil.Encode(result.Pair[:])
		Expect(pair).To(Equal("0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733"))
		Expect(result.Maker.Hex()).To(Equal("0x00Ca405026e9018c29c26Cb081DcC9653428bFe9"))
		Expect(result.Pay_gem.Hex()).To(Equal("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"))
		Expect(result.Buy_gem.Hex()).To(Equal("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"))
		Expect(result.Taker.Hex()).To(Equal("0x0092Ad2b9ae189D50F9cd8E7F4c3355C2c93e3fc"))
		takeAmt := new(big.Int)
		takeAmt.SetString("34334082741116751270", 10)
		Expect(result.Take_amt).To(Equal(takeAmt))
		giveAmt := new(big.Int)
		giveAmt.SetString("27055257200000000002", 10)
		Expect(result.Give_amt).To(Equal(giveAmt))
		Expect(result.Block).To(Equal(blocknumber))
		Expect(result.Tx).To(Equal(watchedEvent.TxHash))
		Expect(result.Timestamp).To(Equal(uint64(1499649315)))
	})

	It("converts watched event from address 0x3A into LogTakeEntity struct", func() {
		watchedEvent := core.WatchedEvent{
			LogID:       0,
			Name:        "",
			BlockNumber: 4750060,
			Address:     "0x3aa927a97594c3ab7d7bf0d47c71c3877d1de4a1",
			TxHash:      "0x5a89f89609794bc59838ac53b319ef19df34bb2060eefa759e135c5af63ba132",
			Index:       105,
			Topic0:      constants.LogTakeSignature,
			Topic1:      "0xc51ce3446dae6b90a7707b1c95f15be978a22a6cd998284bdc3a52a1387caed4",
			Topic2:      "0x000000000000000000000000ab8d8b74f202f4cd4a918b65da4bac612e086ee7",
			Topic3:      "0x0000000000000000000000000e4555922c52ffddcfb006d3dbc94b21541f0f15",
			Data:        "0x0000000000000000000000000000000000000000000000000000000000001761000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a700000000000000000000000059adcf176ed2f6788a41b8ea4c4904518e62b6a40000000000000000000000000000000000000000000000026b749b4c3d79ac350000000000000000000000000000000000000000000006edace0a94c056213c1000000000000000000000000000000000000000000000000000000005a36c112",
		}

		result, err := log_take.LogTakeConverter{}.ToEntity(watchedEvent)
		Expect(err).NotTo(HaveOccurred())

		Expect(common.ToHex(result.Id[:])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000001761"))
		p := hexutil.Encode(result.Pair[:])
		Expect(p).To(Equal("0xc51ce3446dae6b90a7707b1c95f15be978a22a6cd998284bdc3a52a1387caed4"))
		Expect(result.Maker.Hex()).To(Equal("0xab8D8b74F202f4cD4A918B65dA4bAc612e086Ee7"))
		Expect(result.Pay_gem.Hex()).To(Equal("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"))
		Expect(result.Buy_gem.Hex()).To(Equal("0x59aDCF176ED2f6788A41B8eA4c4904518e62B6A4"))
		Expect(result.Taker.Hex()).To(Equal("0x0E4555922c52FFDdcfb006D3dBc94B21541F0F15"))
		ta := new(big.Int)
		ta.SetString("44636472558527032373", 10)
		Expect(result.Take_amt).To(Equal(ta))
		ga := new(big.Int)
		ga.SetString("32718534385400314729409", 10)
		Expect(result.Give_amt).To(Equal(ga))
		Expect(result.Block).To(Equal(int64(4750060)))
		Expect(result.Tx).To(Equal(watchedEvent.TxHash))
		Expect(result.Timestamp).To(Equal(uint64(1513537810)))
	})

	It("converts watched event from address 0x14 into LogTakeEntity struct", func() {
		watchedEvent := core.WatchedEvent{
			LogID:       0,
			Name:        "",
			BlockNumber: 5211385,
			Address:     constants.ContractAddress,
			TxHash:      "0x73c95b64c079f301d6e915441a1730c5c5a146e0c9a877c6aa431eea3603c4f5",
			Index:       10,
			Topic0:      constants.LogTakeSignature,
			Topic1:      "0x203e3b62e033f1548774797f9135c574dddd995f86e1c55fe1bab1610d35094f",
			Topic2:      "0x0000000000000000000000008142e5658a611821b3aa32f177a1d975c60a92f4",
			Topic3:      "0x0000000000000000000000000016bd4cb70bd98ca07a341da66450b5d22a55aa",
			Data:        "0x00000000000000000000000000000000000000000000000000000000000064750000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a2000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000262f9b28cc6e5db00000000000000000000000000000000000000000000000002be9f26eeb18855000000000000000000000000000000000000000000000000000000005a9f95a1",
		}

		result, err := log_take.LogTakeConverter{}.ToEntity(watchedEvent)

		Expect(err).NotTo(HaveOccurred())
		Expect(common.ToHex(result.Id[:])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000006475"))
		p := hexutil.Encode(result.Pair[:])
		Expect(p).To(Equal("0x203e3b62e033f1548774797f9135c574dddd995f86e1c55fe1bab1610d35094f"))
		Expect(result.Maker.Hex()).To(Equal("0x8142E5658a611821b3Aa32F177a1d975c60A92F4"))
		Expect(result.Pay_gem.Hex()).To(Equal("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"))
		Expect(result.Buy_gem.Hex()).To(Equal("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"))
		Expect(result.Taker.Hex()).To(Equal("0x0016BD4cb70Bd98CA07A341DA66450b5d22A55aa"))
		ta := new(big.Int)
		ta.SetString("171974281054840283", 10)
		Expect(result.Take_amt).To(Equal(ta))
		ga := new(big.Int)
		ga.SetString("197770423213066325", 10)
		Expect(result.Give_amt).To(Equal(ga))
		Expect(result.Block).To(Equal(int64(5211385)))
		Expect(result.Tx).To(Equal(watchedEvent.TxHash))
		Expect(result.Timestamp).To(Equal(uint64(1520407969)))
	})

	It("converts a LogTakeEntity to a LogTakeModel", func() {
		ga := new(big.Int)
		ga.SetString("27055257200000000002", 10)
		ta := new(big.Int)
		ta.SetString("34334082741116751270", 10)
		lt := log_take.LogTakeEntity{
			Id:        [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 192},
			Pair:      [32]byte{61, 175, 114, 65, 111, 216, 142, 69, 107, 116, 90, 155, 119, 38, 78, 103, 101, 211, 73, 188, 158, 218, 55, 162, 185, 52, 124, 126, 18, 144, 39, 51},
			Maker:     common.HexToAddress("0x00Ca405026e9018c29c26Cb081DcC9653428bFe9"),
			Pay_gem:   common.HexToAddress("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
			Buy_gem:   common.HexToAddress("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"),
			Taker:     common.HexToAddress("0x0092Ad2b9ae189D50F9cd8E7F4c3355C2c93e3fc"),
			Take_amt:  ta,
			Give_amt:  ga,
			Block:     4000870,
			Timestamp: uint64(1499649315),
		}

		ltc := log_take.LogTakeConverter{}
		ltm := ltc.ToModel(lt)

		expectedPair := "0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733"
		expectedID := int64(2496)
		Expect(ltm.ID).To(Equal(expectedID))
		Expect(ltm.Pair).To(Equal(expectedPair))
		Expect(ltm.Guy).To(Equal("0x00ca405026e9018c29c26cb081dcc9653428bfe9"))
		Expect(ltm.Lot).To(Equal(ga.String()))
		Expect(ltm.Gem).To(Equal("0xc66ea802717bfb9833400264dd12c2bceaa34a6d"))
		Expect(ltm.Pie).To(Equal("0xecf8f87f810ecf450940c9f60066b4a7a501d6a7"))
		Expect(ltm.Bid).To(Equal(ta.String()))
		Expect(ltm.Block).To(Equal(lt.Block))
		Expect(ltm.Timestamp).To(Equal(time.Unix(int64(lt.Timestamp), 0)))
	})
})
