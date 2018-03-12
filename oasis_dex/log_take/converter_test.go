package log_take_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/big"

	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

var _ = Describe("Log Take Converter", func() {

	var event = types.Log{
		Address: common.HexToAddress("0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f"),
		Topics: []common.Hash{common.HexToHash("0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f"),
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

	It("unpacks log", func() {
		parser, err := geth.ParseAbi(`[{"constant":false,"inputs":[],"name":"stop","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"haveToken","type":"address"},{"name":"wantToken","type":"address"},{"name":"haveAmount","type":"uint128"},{"name":"wantAmount","type":"uint128"}],"name":"make","outputs":[{"name":"id","type":"bytes32"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"owner_","type":"address"}],"name":"setOwner","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"last_offer_id","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"cancel","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getOffer","outputs":[{"name":"","type":"uint256"},{"name":"","type":"address"},{"name":"","type":"uint256"},{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"bytes32"},{"name":"maxTakeAmount","type":"uint128"}],"name":"take","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"getTime","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"close_time","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"lifetime","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"stopped","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id_","type":"bytes32"}],"name":"bump","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"authority_","type":"address"}],"name":"setAuthority","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"isActive","outputs":[{"name":"active","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"offers","outputs":[{"name":"sell_how_much","type":"uint256"},{"name":"sell_which_token","type":"address"},{"name":"buy_how_much","type":"uint256"},{"name":"buy_which_token","type":"address"},{"name":"owner","type":"address"},{"name":"active","type":"bool"},{"name":"timestamp","type":"uint64"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"bytes32"}],"name":"kill","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"authority","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"isClosed","outputs":[{"name":"closed","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getOwner","outputs":[{"name":"owner","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"quantity","type":"uint256"}],"name":"buy","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"sell_how_much","type":"uint256"},{"name":"sell_which_token","type":"address"},{"name":"buy_how_much","type":"uint256"},{"name":"buy_which_token","type":"address"}],"name":"offer","outputs":[{"name":"id","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"lifetime_","type":"uint256"}],"payable":false,"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint256"}],"name":"ItemUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"sell_how_much","type":"uint256"},{"indexed":true,"name":"sell_which_token","type":"address"},{"indexed":false,"name":"buy_how_much","type":"uint256"},{"indexed":true,"name":"buy_which_token","type":"address"}],"name":"Trade","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"haveToken","type":"address"},{"indexed":false,"name":"wantToken","type":"address"},{"indexed":false,"name":"haveAmount","type":"uint128"},{"indexed":false,"name":"wantAmount","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogMake","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"haveToken","type":"address"},{"indexed":false,"name":"wantToken","type":"address"},{"indexed":false,"name":"haveAmount","type":"uint128"},{"indexed":false,"name":"wantAmount","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogBump","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"haveToken","type":"address"},{"indexed":false,"name":"wantToken","type":"address"},{"indexed":true,"name":"taker","type":"address"},{"indexed":false,"name":"takeAmount","type":"uint128"},{"indexed":false,"name":"giveAmount","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogTake","type":"event"}, {"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"haveToken","type":"address"},{"indexed":false,"name":"wantToken","type":"address"},{"indexed":false,"name":"haveAmount","type":"uint128"},{"indexed":false,"name":"wantAmount","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogKill","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"authority","type":"address"}],"name":"LogSetAuthority","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"}],"name":"LogSetOwner","type":"event"}]`)
		Expect(err).NotTo(HaveOccurred())
		contract := bind.NewBoundContract(common.StringToAddress("0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f"), parser, nil, nil, nil)
		result := &log_take.LogTakeEntity{}

		contract.UnpackLog(result, "LogTake", event)

		p := hexutil.Encode(result.Pair[:])
		Expect(p).To(Equal("0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733"))
		Expect(result.Maker.Hex()).To(Equal("0x00Ca405026e9018c29c26Cb081DcC9653428bFe9"))
		Expect(result.HaveToken.Hex()).To(Equal("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"))
		Expect(result.WantToken.Hex()).To(Equal("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"))
		Expect(result.Taker.Hex()).To(Equal("0x0092Ad2b9ae189D50F9cd8E7F4c3355C2c93e3fc"))
		ta := new(big.Int)
		ta.SetString("34334082741116751270", 10)
		Expect(result.TakeAmount).To(Equal(ta))
		ga := new(big.Int)
		ga.SetString("27055257200000000002", 10)
		Expect(result.GiveAmount).To(Equal(ga))
		Expect(result.Timestamp).To(Equal(uint64(1499649315)))
	})

	It("converts watched event from address 0x83 into LogTakeEntity struct", func() {
		watchedEvent := core.WatchedEvent{
			LogID:       0,
			Name:        "",
			BlockNumber: 4000870,
			Address:     "0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f",
			TxHash:      "0x98237ddc11a618f5546cd3098e57d9ba159418cb18851fb98130cb3114063807",
			Index:       50,
			Topic0:      "0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f",
			Topic1:      "0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733",
			Topic2:      "0x00000000000000000000000000ca405026e9018c29c26cb081dcc9653428bfe9",
			Topic3:      "0x0000000000000000000000000092ad2b9ae189d50f9cd8e7f4c3355c2c93e3fc",
			Data:        "0x00000000000000000000000000000000000000000000000000000000000009c0000000000000000000000000c66ea802717bfb9833400264dd12c2bceaa34a6d000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7000000000000000000000000000000000000000000000001dc7b2a57e1224da60000000000000000000000000000000000000000000000017777951418d46002000000000000000000000000000000000000000000000000000000005962d523",
		}

		result, err := log_take.LogTakeConverter{}.Convert(watchedEvent)

		Expect(err).NotTo(HaveOccurred())
		p := hexutil.Encode(result.Pair[:])
		Expect(p).To(Equal("0x3daf72416fd88e456b745a9b77264e6765d349bc9eda37a2b9347c7e12902733"))
		Expect(result.Maker.Hex()).To(Equal("0x00Ca405026e9018c29c26Cb081DcC9653428bFe9"))
		Expect(result.HaveToken.Hex()).To(Equal("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"))
		Expect(result.WantToken.Hex()).To(Equal("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"))
		Expect(result.Taker.Hex()).To(Equal("0x0092Ad2b9ae189D50F9cd8E7F4c3355C2c93e3fc"))
		ta := new(big.Int)
		ta.SetString("34334082741116751270", 10)
		Expect(result.TakeAmount).To(Equal(ta))
		ga := new(big.Int)
		ga.SetString("27055257200000000002", 10)
		Expect(result.GiveAmount).To(Equal(ga))
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
			Topic0:      "0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f",
			Topic1:      "0xc51ce3446dae6b90a7707b1c95f15be978a22a6cd998284bdc3a52a1387caed4",
			Topic2:      "0x000000000000000000000000ab8d8b74f202f4cd4a918b65da4bac612e086ee7",
			Topic3:      "0x0000000000000000000000000e4555922c52ffddcfb006d3dbc94b21541f0f15",
			Data:        "0x0000000000000000000000000000000000000000000000000000000000001761000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a700000000000000000000000059adcf176ed2f6788a41b8ea4c4904518e62b6a40000000000000000000000000000000000000000000000026b749b4c3d79ac350000000000000000000000000000000000000000000006edace0a94c056213c1000000000000000000000000000000000000000000000000000000005a36c112",
		}

		result, err := log_take.LogTakeConverter{}.Convert(watchedEvent)

		Expect(err).NotTo(HaveOccurred())
		Expect(common.ToHex(result.Id[:])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000001761"))
		p := hexutil.Encode(result.Pair[:])
		Expect(p).To(Equal("0xc51ce3446dae6b90a7707b1c95f15be978a22a6cd998284bdc3a52a1387caed4"))
		Expect(result.Maker.Hex()).To(Equal("0xab8D8b74F202f4cD4A918B65dA4bAc612e086Ee7"))
		Expect(result.HaveToken.Hex()).To(Equal("0xECF8F87f810EcF450940c9f60066b4a7a501d6A7"))
		Expect(result.WantToken.Hex()).To(Equal("0x59aDCF176ED2f6788A41B8eA4c4904518e62B6A4"))
		Expect(result.Taker.Hex()).To(Equal("0x0E4555922c52FFDdcfb006D3dBc94B21541F0F15"))
		ta := new(big.Int)
		ta.SetString("44636472558527032373", 10)
		Expect(result.TakeAmount).To(Equal(ta))
		ga := new(big.Int)
		ga.SetString("32718534385400314729409", 10)
		Expect(result.GiveAmount).To(Equal(ga))
		Expect(result.Timestamp).To(Equal(uint64(1513537810)))
	})

	It("converts watched event from address 0x14 into LogTakeEntity struct", func() {
		watchedEvent := core.WatchedEvent{
			LogID:       0,
			Name:        "",
			BlockNumber: 5211385,
			Address:     "0x14fbca95be7e99c15cc2996c6c9d841e54b79425",
			TxHash:      "0x73c95b64c079f301d6e915441a1730c5c5a146e0c9a877c6aa431eea3603c4f5",
			Index:       10,
			Topic0:      "0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f",
			Topic1:      "0x203e3b62e033f1548774797f9135c574dddd995f86e1c55fe1bab1610d35094f",
			Topic2:      "0x0000000000000000000000008142e5658a611821b3aa32f177a1d975c60a92f4",
			Topic3:      "0x0000000000000000000000000016bd4cb70bd98ca07a341da66450b5d22a55aa",
			Data:        "0x00000000000000000000000000000000000000000000000000000000000064750000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a2000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000262f9b28cc6e5db00000000000000000000000000000000000000000000000002be9f26eeb18855000000000000000000000000000000000000000000000000000000005a9f95a1",
		}

		result, err := log_take.LogTakeConverter{}.Convert(watchedEvent)

		Expect(err).NotTo(HaveOccurred())
		Expect(common.ToHex(result.Id[:])).To(Equal("0x0000000000000000000000000000000000000000000000000000000000006475"))
		p := hexutil.Encode(result.Pair[:])
		Expect(p).To(Equal("0x203e3b62e033f1548774797f9135c574dddd995f86e1c55fe1bab1610d35094f"))
		Expect(result.Maker.Hex()).To(Equal("0x8142E5658a611821b3Aa32F177a1d975c60A92F4"))
		Expect(result.HaveToken.Hex()).To(Equal("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"))
		Expect(result.WantToken.Hex()).To(Equal("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"))
		Expect(result.Taker.Hex()).To(Equal("0x0016BD4cb70Bd98CA07A341DA66450b5d22A55aa"))
		ta := new(big.Int)
		ta.SetString("171974281054840283", 10)
		Expect(result.TakeAmount).To(Equal(ta))
		ga := new(big.Int)
		ga.SetString("197770423213066325", 10)
		Expect(result.GiveAmount).To(Equal(ga))
		Expect(result.Timestamp).To(Equal(uint64(1520407969)))
	})
})
