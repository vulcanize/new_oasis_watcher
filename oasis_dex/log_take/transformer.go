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
