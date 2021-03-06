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

package constants

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/filters"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

var ABI abi.ABI
var err error

func init() {
	ABI, err = geth.ParseAbi(`[{"constant":true,"inputs":[],"name":"matchingEnabled","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"sell_gem","type":"address"},{"name":"buy_gem","type":"address"}],"name":"getBestOffer","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"pay_gem","type":"address"},{"name":"pay_amt","type":"uint256"},{"name":"buy_gem","type":"address"},{"name":"min_fill_amount","type":"uint256"}],"name":"sellAllAmount","outputs":[{"name":"fill_amt","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"stop","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"pay_gem","type":"address"},{"name":"buy_gem","type":"address"},{"name":"pay_amt","type":"uint128"},{"name":"buy_amt","type":"uint128"}],"name":"make","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"owner_","type":"address"}],"name":"setOwner","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"buy_gem","type":"address"},{"name":"pay_gem","type":"address"},{"name":"pay_amt","type":"uint256"}],"name":"getBuyAmount","outputs":[{"name":"fill_amt","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"baseToken","type":"address"},{"name":"quoteToken","type":"address"}],"name":"addTokenPairWhitelist","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"baseToken","type":"address"},{"name":"quoteToken","type":"address"}],"name":"remTokenPairWhitelist","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"pay_amt","type":"uint256"},{"name":"pay_gem","type":"address"},{"name":"buy_amt","type":"uint256"},{"name":"buy_gem","type":"address"},{"name":"pos","type":"uint256"}],"name":"offer","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"pos","type":"uint256"}],"name":"insert","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"last_offer_id","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"matchingEnabled_","type":"bool"}],"name":"setMatchingEnabled","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"cancel","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getOffer","outputs":[{"name":"","type":"uint256"},{"name":"","type":"address"},{"name":"","type":"uint256"},{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"del_rank","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"bytes32"},{"name":"maxTakeAmount","type":"uint128"}],"name":"take","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"pay_gem","type":"address"}],"name":"getMinSell","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getTime","outputs":[{"name":"","type":"uint64"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getNextUnsortedOffer","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"close_time","outputs":[{"name":"","type":"uint64"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"_span","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"_best","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"stopped","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"id_","type":"bytes32"}],"name":"bump","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"authority_","type":"address"}],"name":"setAuthority","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"sell_gem","type":"address"},{"name":"buy_gem","type":"address"}],"name":"getOfferCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"buy_gem","type":"address"},{"name":"buy_amt","type":"uint256"},{"name":"pay_gem","type":"address"},{"name":"max_fill_amount","type":"uint256"}],"name":"buyAllAmount","outputs":[{"name":"fill_amt","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"isActive","outputs":[{"name":"active","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"offers","outputs":[{"name":"pay_amt","type":"uint256"},{"name":"pay_gem","type":"address"},{"name":"buy_amt","type":"uint256"},{"name":"buy_gem","type":"address"},{"name":"owner","type":"address"},{"name":"timestamp","type":"uint64"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getFirstUnsortedOffer","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"baseToken","type":"address"},{"name":"quoteToken","type":"address"}],"name":"isTokenPairWhitelisted","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getBetterOffer","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"_dust","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getWorseOffer","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"_menu","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"_near","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"bytes32"}],"name":"kill","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"pay_gem","type":"address"},{"name":"dust","type":"uint256"}],"name":"setMinSell","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"authority","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"isClosed","outputs":[{"name":"closed","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"_rank","outputs":[{"name":"next","type":"uint256"},{"name":"prev","type":"uint256"},{"name":"delb","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getOwner","outputs":[{"name":"owner","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"isOfferSorted","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"buyEnabled_","type":"bool"}],"name":"setBuyEnabled","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"amount","type":"uint256"}],"name":"buy","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"pay_amt","type":"uint256"},{"name":"pay_gem","type":"address"},{"name":"buy_amt","type":"uint256"},{"name":"buy_gem","type":"address"},{"name":"pos","type":"uint256"},{"name":"rounding","type":"bool"}],"name":"offer","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"pay_amt","type":"uint256"},{"name":"pay_gem","type":"address"},{"name":"buy_amt","type":"uint256"},{"name":"buy_gem","type":"address"}],"name":"offer","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"buyEnabled","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"pay_gem","type":"address"},{"name":"buy_gem","type":"address"},{"name":"buy_amt","type":"uint256"}],"name":"getPayAmount","outputs":[{"name":"fill_amt","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"close_time","type":"uint64"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":true,"inputs":[{"indexed":true,"name":"sig","type":"bytes4"},{"indexed":true,"name":"guy","type":"address"},{"indexed":true,"name":"foo","type":"bytes32"},{"indexed":true,"name":"bar","type":"bytes32"},{"indexed":false,"name":"wad","type":"uint256"},{"indexed":false,"name":"fax","type":"bytes"}],"name":"LogNote","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint256"}],"name":"LogItemUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"pay_amt","type":"uint256"},{"indexed":true,"name":"pay_gem","type":"address"},{"indexed":false,"name":"buy_amt","type":"uint256"},{"indexed":true,"name":"buy_gem","type":"address"}],"name":"LogTrade","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"pay_gem","type":"address"},{"indexed":false,"name":"buy_gem","type":"address"},{"indexed":false,"name":"pay_amt","type":"uint128"},{"indexed":false,"name":"buy_amt","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogMake","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"pay_gem","type":"address"},{"indexed":false,"name":"buy_gem","type":"address"},{"indexed":false,"name":"pay_amt","type":"uint128"},{"indexed":false,"name":"buy_amt","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogBump","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"pay_gem","type":"address"},{"indexed":false,"name":"buy_gem","type":"address"},{"indexed":true,"name":"taker","type":"address"},{"indexed":false,"name":"take_amt","type":"uint128"},{"indexed":false,"name":"give_amt","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogTake","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"id","type":"bytes32"},{"indexed":true,"name":"pair","type":"bytes32"},{"indexed":true,"name":"maker","type":"address"},{"indexed":false,"name":"pay_gem","type":"address"},{"indexed":false,"name":"buy_gem","type":"address"},{"indexed":false,"name":"pay_amt","type":"uint128"},{"indexed":false,"name":"buy_amt","type":"uint128"},{"indexed":false,"name":"timestamp","type":"uint64"}],"name":"LogKill","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"authority","type":"address"}],"name":"LogSetAuthority","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"}],"name":"LogSetOwner","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"isEnabled","type":"bool"}],"name":"LogBuyEnabled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"pay_gem","type":"address"},{"indexed":false,"name":"min_amount","type":"uint256"}],"name":"LogMinSell","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"isEnabled","type":"bool"}],"name":"LogMatchingEnabled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint256"}],"name":"LogUnsortedOffer","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint256"}],"name":"LogSortedOffer","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"baseToken","type":"address"},{"indexed":false,"name":"quoteToken","type":"address"}],"name":"LogAddTokenPairWhitelist","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"baseToken","type":"address"},{"indexed":false,"name":"quoteToken","type":"address"}],"name":"LogRemTokenPairWhitelist","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"keeper","type":"address"},{"indexed":false,"name":"id","type":"uint256"}],"name":"LogInsert","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"keeper","type":"address"},{"indexed":false,"name":"id","type":"uint256"}],"name":"LogDelete","type":"event"}]`)
	if err != nil {
		log.Fatal("constants: unable to parse abi")
	}
}

var ContractAddress = "0x14fbca95be7e99c15cc2996c6c9d841e54b79425"
var LogTakeSignature = "0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f"
var LogMakeSignature = "0x773ff502687307abfa024ac9f62f9752a0d210dac2ffd9a29e38e12e2ea82c82"
var LogKillSignature = "0x9577941d28fff863bfbee4694a6a4a56fb09e169619189d2eaa750b5b4819995"

var LogTakeFilters = []filters.LogFilter{
	{
		Name:      "LogTake-0x83",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f",
		Topics:    core.Topics{LogTakeSignature},
	},
	{
		Name:      "LogTake-0x14",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   ContractAddress,
		Topics:    core.Topics{LogTakeSignature},
	},
	{
		Name:      "LogTake-0x3a",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "0x3aa927a97594c3ab7d7bf0d47c71c3877d1de4a1",
		Topics:    core.Topics{LogTakeSignature},
	},
	// Don't think any of these ever happened before this contract went dormant - should add specific tests
	// to the converter for this example if one pops up
	{
		Name:      "LogTake-0x91",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "0x91dfe531ff8ba876a505c8f1c98bafede6c7effc",
		Topics:    core.Topics{LogTakeSignature},
	},
}

var LogMakeFilters = []filters.LogFilter{
	{
		Name:      "LogMake",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   ContractAddress,
		Topics:    core.Topics{LogMakeSignature},
	},
}

var LogKillFilters = []filters.LogFilter{
	{
		Name:      "LogKill",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   ContractAddress,
		Topics:    core.Topics{LogKillSignature},
	},
}
