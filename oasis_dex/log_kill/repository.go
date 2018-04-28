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
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type Datastore interface {
	Create(model LogKillModel, vulcanizeLogId int64) error
}

type Repository struct {
	*postgres.DB
}

func (repository Repository) Create(logKillModel LogKillModel, vulcanizeLogId int64) error {
	_, err := repository.DB.Exec(

		`INSERT INTO oasis.kill (id, vulcanize_log_id, pair, gem, lot, pie, bid, guy, block, "time", tx)
               VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
                ON CONFLICT (vulcanize_log_id) DO NOTHING`,
		logKillModel.ID, vulcanizeLogId, logKillModel.Pair, logKillModel.Gem, logKillModel.Lot, logKillModel.Pie, logKillModel.Bid, logKillModel.Guy, logKillModel.Block, logKillModel.Timestamp, logKillModel.Tx)

	if err != nil {
		return err
	}

	return nil
}
