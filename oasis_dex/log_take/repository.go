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
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type Datastore interface {
	Create(logTake LogTakeModel, vulcanizeLogId int64) error
}

type Repository struct {
	*postgres.DB
}

func (repository Repository) Create(logTake LogTakeModel, vulcanizeLogId int64) error {
	_, err := repository.Exec(
		`INSERT INTO oasis.log_take (vulcanize_log_id, id, pair, guy, gem, lot, gal, pie, bid, block, time, tx)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
                ON CONFLICT (vulcanize_log_id) DO NOTHING`,
		vulcanizeLogId, logTake.ID, logTake.Pair, logTake.Guy, logTake.Gem, logTake.Lot, logTake.Gal, logTake.Pie, logTake.Bid, logTake.Block, logTake.Timestamp, logTake.Tx,
	)
	if err != nil {
		return err
	}
	return nil
}
