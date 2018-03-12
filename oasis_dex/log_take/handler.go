package log_take

import (
	"log"

	"github.com/vulcanize/vulcanizedb/libraries/shared"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/filters"
)

type Handler struct {
	Converter              ILogTakeConverter
	WatchedEventRepository datastore.WatchedEventRepository
	FilterRepository       datastore.FilterRepository
	OasisLogRepository     IOasisLogRepository
}

var logTakeFilters = []filters.LogFilter{
	{
		Name:      "LogTake-0x83",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "0x83ce340889c15a3b4d38cfcd1fc93e5d8497691f",
		Topics:    core.Topics{"0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f"},
	},
	{
		Name:      "LogTake-0x14",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "0x14fbca95be7e99c15cc2996c6c9d841e54b79425",
		Topics:    core.Topics{"0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f"},
	},
	{
		Name:      "LogTake-0x3a",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "0x3aa927a97594c3ab7d7bf0d47c71c3877d1de4a1",
		Topics:    core.Topics{"0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f"},
	},
	// Don't think any of these ever happened before this contract went dormant - should add specific tests
	// to the converter for this example if one pops up
	{
		Name:      "LogTake-0x91",
		FromBlock: 0,
		ToBlock:   -1,
		Address:   "0x91dfe531ff8ba876a505c8f1c98bafede6c7effc",
		Topics:    core.Topics{"0x3383e3357c77fd2e3a4b30deea81179bc70a795d053d14d5b7f2f01d0fd4596f"},
	},
}

func NewLogTakeHandler(db *postgres.DB, blockchain core.Blockchain) shared.Handler {
	var handler shared.Handler
	cnvtr := LogTakeConverter{}
	wer := repositories.WatchedEventRepository{DB: db}
	fr := repositories.FilterRepository{DB: db}
	olr := OasisLogRepository{db}
	handler = &Handler{cnvtr, wer, fr, olr}
	for _, filter := range logTakeFilters {
		fr.CreateFilter(filter)
	}
	return handler
}

func (logTakeHandler *Handler) Execute() error {
	for _, filter := range logTakeFilters {
		watchedEvents, err := logTakeHandler.WatchedEventRepository.GetWatchedEvents(filter.Name)
		if err != nil {
			log.Println("Error fetching events for filter: ", err)
			return err
		}
		for _, watchedEvent := range watchedEvents {
			err = createLogTakeData(watchedEvent, logTakeHandler)
			if err != nil {
				log.Println("Error persisting data for event: ", err)
			}
		}
	}
	return nil
}

func createLogTakeData(watchedEvent *core.WatchedEvent, logTakeHandler *Handler) error {
	logEvent, err := logTakeHandler.Converter.Convert(*watchedEvent)
	if err != nil {
		log.Println("Error converting watched event to log: ", err)
		return err
	}
	err = logTakeHandler.OasisLogRepository.CreateLogTake(*logEvent, watchedEvent.LogID)
	if err != nil {
		log.Println("Error persisting data for event: ", err)
		return err
	}
	return nil
}
