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

package log_take

import (
	"strings"
	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

type Converter interface {
	ToEntity(watchedEvent core.WatchedEvent) (*LogTakeEntity, error)
	ToModel(entity LogTakeEntity) LogTakeModel
}

type LogTakeConverter struct {
}

func (LogTakeConverter) ToEntity(watchedEvent core.WatchedEvent) (*LogTakeEntity, error) {
	result := &LogTakeEntity{}
	contract := bind.NewBoundContract(common.HexToAddress(constants.ContractAddress), constants.ABI, nil, nil, nil)
	event := helpers.ConvertToLog(watchedEvent)
	err := contract.UnpackLog(result, "LogTake", event)
	if err != nil {
		return result, err
	}
	result.Block = watchedEvent.BlockNumber
	result.Tx = watchedEvent.TxHash
	return result, nil
}

func (LogTakeConverter) ToModel(entity LogTakeEntity) LogTakeModel {
	id := common.BytesToHash(entity.Id[:]).Big().Int64()
	pair := common.ToHex(entity.Pair[:])
	guy := strings.ToLower(entity.Maker.String())
	gem := strings.ToLower(entity.Pay_gem.String())
	lot := strings.ToLower(entity.Give_amt.String())
	gal := strings.ToLower(entity.Taker.String())
	pie := strings.ToLower(entity.Buy_gem.String())
	bid := strings.ToLower(entity.Take_amt.String())
	timestamp := time.Unix(int64(entity.Timestamp), 0)
	block := entity.Block
	tx := entity.Tx
	return LogTakeModel{
		ID:        id,
		Pair:      pair,
		Guy:       guy,
		Gem:       gem,
		Pie:       pie,
		Gal:       gal,
		Lot:       lot,
		Bid:       bid,
		Block:     block,
		Tx:        tx,
		Timestamp: timestamp,
	}
}
