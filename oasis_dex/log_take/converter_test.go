package log_take_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/helpers"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

var _ = Describe("Log Take Converter", func() {

	expectedModel := log_take.LogTakeModel{
		ID:        2496,
		Pair:      "0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733",
		Guy:       "0x00ca405026e9018c29c26cb081dcc9653428bfe9",
		Gem:       "0xc66ea802717bfb9833400264dd12c2bceaa34a6d",
		Lot:       "27055257200000000002",
		Gal:       "0x0092ad2b9ae189d50f9cd8e7f4c3355c2c93e3fc",
		Pie:       "0xecf8f87f810ecf450940c9f60066b4a7a501d6a7",
		Bid:       "34334082741116751270",
		Block:     4000870,
		Tx:        "0x73c95b64c079f301d6e915441a1730c5c5a146e0c9a877c6aa431eea3603c4f5",
		Timestamp: time.Unix(1499649315, 0),
	}

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

		expectedEntity := log_take.LogTakeEntity{
			Id:        [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 192},
			Pair:      [32]byte{61, 175, 114, 65, 111, 216, 142, 69, 107, 116, 90, 155, 119, 38, 78, 103, 101, 211, 73, 188, 158, 218, 55, 162, 185, 52, 124, 126, 18, 144, 39, 51},
			Maker:     common.HexToAddress("0x00Ca405026e9018c29c26Cb081DcC9653428bFe9"),
			Pay_gem:   common.HexToAddress("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
			Buy_gem:   common.HexToAddress("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"),
			Taker:     common.HexToAddress("0x0092Ad2b9ae189D50F9cd8E7F4c3355C2c93e3fc"),
			Take_amt:  helpers.BigFromString("34334082741116751270"),
			Give_amt:  helpers.BigFromString("27055257200000000002"),
			Timestamp: uint64(1499649315),
		}

		Expect(result.Pair).To(Equal(expectedEntity.Pair))
		Expect(result.Maker).To(Equal(expectedEntity.Maker))
		Expect(result.Pay_gem).To(Equal(expectedEntity.Pay_gem))
		Expect(result.Buy_gem).To(Equal(expectedEntity.Buy_gem))
		Expect(result.Taker).To(Equal(expectedEntity.Taker))
		Expect(result.Take_amt).To(Equal(expectedEntity.Take_amt))
		Expect(result.Give_amt).To(Equal(expectedEntity.Give_amt))
		Expect(result.Block).To(Equal(expectedEntity.Block))
		Expect(result.Timestamp).To(Equal(expectedEntity.Timestamp))
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

		expectedEntity := log_take.LogTakeEntity{
			Id:        [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 192},
			Pair:      [32]byte{61, 175, 114, 65, 111, 216, 142, 69, 107, 116, 90, 155, 119, 38, 78, 103, 101, 211, 73, 188, 158, 218, 55, 162, 185, 52, 124, 126, 18, 144, 39, 51},
			Maker:     common.Address{0, 202, 64, 80, 38, 233, 1, 140, 41, 194, 108, 176, 129, 220, 201, 101, 52, 40, 191, 233},
			Pay_gem:   common.Address{198, 110, 168, 2, 113, 123, 251, 152, 51, 64, 2, 100, 221, 18, 194, 188, 234, 163, 74, 109},
			Buy_gem:   common.Address{236, 248, 248, 127, 129, 14, 207, 69, 9, 64, 201, 246, 0, 102, 180, 167, 165, 1, 214, 167},
			Taker:     common.Address{0, 146, 173, 43, 154, 225, 137, 213, 15, 156, 216, 231, 244, 195, 53, 92, 44, 147, 227, 252},
			Take_amt:  helpers.BigFromString("34334082741116751270 "),
			Give_amt:  helpers.BigFromString("27055257200000000002"),
			Block:     4000870,
			Tx:        "0x98237ddc11a618f5546cd3098e57d9ba159418cb18851fb98130cb3114063807",
			Timestamp: 1499649315,
		}

		Expect(result.Pair).To(Equal(expectedEntity.Pair))
		Expect(result.Maker).To(Equal(expectedEntity.Maker))
		Expect(result.Pay_gem).To(Equal(expectedEntity.Pay_gem))
		Expect(result.Buy_gem).To(Equal(expectedEntity.Buy_gem))
		Expect(result.Taker).To(Equal(expectedEntity.Taker))
		Expect(result.Take_amt).To(Equal(expectedEntity.Take_amt))
		Expect(result.Give_amt).To(Equal(expectedEntity.Give_amt))
		Expect(result.Block).To(Equal(expectedEntity.Block))
		Expect(result.Tx).To(Equal(watchedEvent.TxHash))
		Expect(result.Timestamp).To(Equal(expectedEntity.Timestamp))
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

		expectedEntity := log_take.LogTakeEntity{
			Id:        [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 23, 97},
			Pair:      [32]byte{197, 28, 227, 68, 109, 174, 107, 144, 167, 112, 123, 28, 149, 241, 91, 233, 120, 162, 42, 108, 217, 152, 40, 75, 220, 58, 82, 161, 56, 124, 174, 212},
			Maker:     common.Address{171, 141, 139, 116, 242, 2, 244, 205, 74, 145, 139, 101, 218, 75, 172, 97, 46, 8, 110, 231},
			Pay_gem:   common.Address{236, 248, 248, 127, 129, 14, 207, 69, 9, 64, 201, 246, 0, 102, 180, 167, 165, 1, 214, 167},
			Buy_gem:   common.Address{89, 173, 207, 23, 110, 210, 246, 120, 138, 65, 184, 234, 76, 73, 4, 81, 142, 98, 182, 164},
			Taker:     common.Address{14, 69, 85, 146, 44, 82, 255, 221, 207, 176, 6, 211, 219, 201, 75, 33, 84, 31, 15, 21},
			Take_amt:  helpers.BigFromString("44636472558527032373"),
			Give_amt:  helpers.BigFromString("32718534385400314729409"),
			Block:     4750060,
			Tx:        "0x5a89f89609794bc59838ac53b319ef19df34bb2060eefa759e135c5af63ba132",
			Timestamp: 1513537810,
		}

		Expect(result.Pair).To(Equal(expectedEntity.Pair))
		Expect(result.Maker).To(Equal(expectedEntity.Maker))
		Expect(result.Pay_gem).To(Equal(expectedEntity.Pay_gem))
		Expect(result.Buy_gem).To(Equal(expectedEntity.Buy_gem))
		Expect(result.Taker).To(Equal(expectedEntity.Taker))
		Expect(result.Take_amt).To(Equal(expectedEntity.Take_amt))
		Expect(result.Give_amt).To(Equal(expectedEntity.Give_amt))
		Expect(result.Block).To(Equal(expectedEntity.Block))
		Expect(result.Tx).To(Equal(watchedEvent.TxHash))
		Expect(result.Timestamp).To(Equal(expectedEntity.Timestamp))
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
		expectedEntity := log_take.LogTakeEntity{
			Id:        [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 117},
			Pair:      [32]byte{32, 62, 59, 98, 224, 51, 241, 84, 135, 116, 121, 127, 145, 53, 197, 116, 221, 221, 153, 95, 134, 225, 197, 95, 225, 186, 177, 97, 13, 53, 9, 79},
			Maker:     common.Address{129, 66, 229, 101, 138, 97, 24, 33, 179, 170, 50, 241, 119, 161, 217, 117, 198, 10, 146, 244},
			Pay_gem:   common.Address{159, 143, 114, 170, 147, 4, 200, 181, 147, 213, 85, 241, 46, 246, 88, 156, 195, 165, 121, 162},
			Buy_gem:   common.Address{192, 42, 170, 57, 178, 35, 254, 141, 10, 14, 92, 79, 39, 234, 217, 8, 60, 117, 108, 194},
			Taker:     common.Address{0, 22, 189, 76, 183, 11, 217, 140, 160, 122, 52, 29, 166, 100, 80, 181, 210, 42, 85, 170},
			Take_amt:  helpers.BigFromString("171974281054840283"),
			Give_amt:  helpers.BigFromString("197770423213066325"),
			Block:     5211385,
			Tx:        "0x73c95b64c079f301d6e915441a1730c5c5a146e0c9a877c6aa431eea3603c4f5",
			Timestamp: 1520407969,
		}

		Expect(result.Pair).To(Equal(expectedEntity.Pair))
		Expect(result.Maker).To(Equal(expectedEntity.Maker))
		Expect(result.Pay_gem).To(Equal(expectedEntity.Pay_gem))
		Expect(result.Buy_gem).To(Equal(expectedEntity.Buy_gem))
		Expect(result.Taker).To(Equal(expectedEntity.Taker))
		Expect(result.Take_amt).To(Equal(expectedEntity.Take_amt))
		Expect(result.Give_amt).To(Equal(expectedEntity.Give_amt))
		Expect(result.Block).To(Equal(expectedEntity.Block))
		Expect(result.Tx).To(Equal(watchedEvent.TxHash))
		Expect(result.Timestamp).To(Equal(expectedEntity.Timestamp))
	})

	It("converts a LogTakeEntity to a LogTakeModel", func() {
		entity := log_take.LogTakeEntity{
			Id:        [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 192},
			Pair:      [32]byte{61, 175, 114, 65, 111, 216, 142, 69, 107, 116, 90, 155, 119, 38, 78, 103, 101, 211, 73, 188, 158, 218, 55, 162, 185, 52, 124, 126, 18, 144, 39, 51},
			Maker:     common.HexToAddress("0x00Ca405026e9018c29c26Cb081DcC9653428bFe9"),
			Pay_gem:   common.HexToAddress("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
			Buy_gem:   common.HexToAddress("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"),
			Taker:     common.HexToAddress("0x0092Ad2b9ae189D50F9cd8E7F4c3355C2c93e3fc"),
			Take_amt:  helpers.BigFromString("34334082741116751270"),
			Give_amt:  helpers.BigFromString("27055257200000000002"),
			Block:     4000870,
			Timestamp: uint64(1499649315),
			Tx:        "0x73c95b64c079f301d6e915441a1730c5c5a146e0c9a877c6aa431eea3603c4f5",
		}

		ltc := log_take.LogTakeConverter{}
		model := ltc.ToModel(entity)

		Expect(model.ID).To(Equal(expectedModel.ID))
		Expect(model.Pair).To(Equal(expectedModel.Pair))
		Expect(model.Guy).To(Equal(expectedModel.Guy))
		Expect(model.Lot).To(Equal(expectedModel.Lot))
		Expect(model.Gal).To(Equal(expectedModel.Gal))
		Expect(model.Gem).To(Equal(expectedModel.Gem))
		Expect(model.Pie).To(Equal(expectedModel.Pie))
		Expect(model.Bid).To(Equal(expectedModel.Bid))
		Expect(model.Block).To(Equal(expectedModel.Block))
		Expect(model.Tx).To(Equal(expectedModel.Tx))
		Expect(model.Timestamp).To(Equal(expectedModel.Timestamp))
	})
})
