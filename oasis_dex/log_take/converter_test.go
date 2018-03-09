package log_take_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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
		type LogTake struct {
			Id         [32]byte
			Pair       [32]byte
			Maker      common.Address
			HaveToken  common.Address
			WantToken  common.Address
			Taker      common.Address
			TakeAmount *big.Int
			GiveAmount *big.Int
			Timestamp  uint64
		}
		result := &LogTake{}

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
})
