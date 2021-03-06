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

package log_kill

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
	ToEntity(watchedEvent core.WatchedEvent) (*LogKillEntity, error)
	ToModel(entity LogKillEntity) LogKillModel
}

type LogKillConverter struct {
}

func (LogKillConverter) ToEntity(watchedEvent core.WatchedEvent) (*LogKillEntity, error) {
	result := &LogKillEntity{}
	contract := bind.NewBoundContract(common.HexToAddress(constants.ContractAddress), constants.ABI, nil, nil, nil)
	event := helpers.ConvertToLog(watchedEvent)
	err := contract.UnpackLog(result, "LogKill", event)
	if err != nil {
		return result, err
	}
	result.Block = watchedEvent.BlockNumber
	result.TransactionHash = watchedEvent.TxHash

	return result, nil
}

func (LogKillConverter) ToModel(LogKillEntity LogKillEntity) LogKillModel {
	id := common.BytesToHash(LogKillEntity.Id[:]).Big().Int64()
	pair := strings.ToLower(common.ToHex(LogKillEntity.Pair[:]))
	guy := strings.ToLower(LogKillEntity.Maker.Hex())
	gem := strings.ToLower(LogKillEntity.Pay_gem.Hex())
	lot := LogKillEntity.Pay_amt.String()
	pie := strings.ToLower(LogKillEntity.Buy_gem.Hex())
	bid := LogKillEntity.Buy_amt.String()
	block := LogKillEntity.Block
	timestamp := time.Unix(int64(LogKillEntity.Timestamp), 0)
	tx := LogKillEntity.TransactionHash
	return LogKillModel{
		ID:        id,
		Pair:      pair,
		Guy:       guy,
		Gem:       gem,
		Lot:       lot,
		Pie:       pie,
		Bid:       bid,
		Block:     block,
		Timestamp: timestamp,
		Tx:        tx,
	}
}
