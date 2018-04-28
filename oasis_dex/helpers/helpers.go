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

package helpers

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

func ConvertToLog(watchedEvent core.WatchedEvent) types.Log {
	return types.Log{
		Address:     common.HexToAddress(watchedEvent.Address),
		Topics:      createTopics(watchedEvent.Topic0, watchedEvent.Topic1, watchedEvent.Topic2, watchedEvent.Topic3),
		Data:        hexutil.MustDecode(watchedEvent.Data),
		BlockNumber: uint64(watchedEvent.BlockNumber),
		TxHash:      common.HexToHash(watchedEvent.TxHash),
		TxIndex:     0,
		BlockHash:   common.HexToHash("0x0"),
		Index:       uint(watchedEvent.Index),
		Removed:     false,
	}
}

func createTopics(topic0 string, topic1 string, topic2 string, topic3 string) []common.Hash {
	return []common.Hash{common.HexToHash(topic0), common.HexToHash(topic1), common.HexToHash(topic2), common.HexToHash(topic3)}
}

func BigFromString(n string) *big.Int {
	b := new(big.Int)
	b.SetString(n, 10)
	return b
}
