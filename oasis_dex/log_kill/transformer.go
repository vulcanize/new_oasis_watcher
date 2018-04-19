package log_kill

import (
	"log"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
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
	Repository             Datastore
}

func NewTransformer(db *postgres.DB, blockchain core.Blockchain) shared.Transformer {
	var transformer shared.Transformer
	cnvtr := LogKillConverter{}
	wer := repositories.WatchedEventRepository{DB: db}
	fr := repositories.FilterRepository{DB: db}
	lkr := Repository{db}
	transformer = &Transformer{
		Converter:              cnvtr,
		WatchedEventRepository: wer,
		FilterRepository:       fr,
		Repository:             lkr,
	}
	for _, filter := range constants.LogKillFilters {
		fr.CreateFilter(filter)
	}
	return transformer
}

func (logKillTransformer Transformer) Execute() error {
	for _, filter := range constants.LogKillFilters {
		watchedEvents, err := logKillTransformer.WatchedEventRepository.GetWatchedEvents(filter.Name)
		if err != nil {
			log.Println("Error fetching events for LogKill filter:", err)
			return err
		}
		for _, we := range watchedEvents {
			logKill, err := logKillTransformer.Converter.ToModel(*we)
			if err != nil {
				log.Printf("Error persisting data for LogKill (watchedEvent.LogID %s):\n %s", we.LogID, err)
			}
			logKillTransformer.Repository.Remove(*logKill)
		}
	}
	return nil
}
