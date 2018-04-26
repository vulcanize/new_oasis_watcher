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

package log_make

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type LogMakeEntity struct {
	Id              [32]byte
	Pair            [32]byte
	Maker           common.Address
	Pay_gem         common.Address
	Buy_gem         common.Address
	Pay_amt         *big.Int
	Buy_amt         *big.Int
	Block           int64
	Timestamp       uint64
	TransactionHash string
}
