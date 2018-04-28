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
	"log"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_make"
	"github.com/vulcanize/vulcanizedb/libraries/shared"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

type Transformer struct {
	Converter              Converter
	WatchedEventRepository datastore.WatchedEventRepository
	FilterRepository       datastore.FilterRepository
	LogMakeRepository      log_make.Datastore
	Repository             Datastore
}

func NewTransformer(db *postgres.DB, blockchain core.Blockchain) shared.Transformer {
	var transformer shared.Transformer
	cnvtr := LogTakeConverter{}
	wer := repositories.WatchedEventRepository{DB: db}
	fr := repositories.FilterRepository{DB: db}
	repo := Repository{db}
	lmr := log_make.Repository{DB: db}
	transformer = &Transformer{
		Converter:              cnvtr,
		WatchedEventRepository: wer,
		FilterRepository:       fr,
		LogMakeRepository:      lmr,
		Repository:             repo,
	}
	for _, filter := range constants.LogTakeFilters {
		fr.CreateFilter(filter)
	}
	return transformer
}

func (logTakeTransformer *Transformer) Execute() error {
	for _, filter := range constants.LogTakeFilters {
		watchedEvents, err := logTakeTransformer.WatchedEventRepository.GetWatchedEvents(filter.Name)
		if err != nil {
			log.Println("Error fetching events for LogTake filter: ", err)
			return err
		}
		for _, watchedEvent := range watchedEvents {
			err = createLogTakeData(watchedEvent, logTakeTransformer)
			if err != nil {
				log.Printf("Error persisting data for LogTake (watchedEvent.LogID %s):\n %s", watchedEvent.LogID, err)
			}
		}
	}
	return nil
}

func createLogTakeData(watchedEvent *core.WatchedEvent, logTakeTransformer *Transformer) error {
	logEvent, err := logTakeTransformer.Converter.ToEntity(*watchedEvent)
	if err != nil {
		log.Println("Error converting LogTake watched event to log: ", err)
		return err
	}
	lm := logTakeTransformer.Converter.ToModel(*logEvent)
	err = logTakeTransformer.Repository.Create(lm, watchedEvent.LogID)
	if err != nil {
		log.Println("Error persisting data for LogTake event: ", err)
		return err
	}

	return nil
}
