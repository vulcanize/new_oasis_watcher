package log_take_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/filters"
)

type MockLogTakeConverter struct {
}

func (MockLogTakeConverter) Convert(watchedEvent core.WatchedEvent) (*log_take.LogTakeEntity, error) {
	return &log_take.LogTakeEntity{}, nil
}

type MockWatchedEventsRepository struct {
	watchedEvents []*core.WatchedEvent
	names         []string
}

func (mwer *MockWatchedEventsRepository) SetWatchedEvents(watchedEvents []*core.WatchedEvent) {
	mwer.watchedEvents = watchedEvents
}

func (mwer *MockWatchedEventsRepository) GetWatchedEvents(name string) ([]*core.WatchedEvent, error) {
	mwer.names = append(mwer.names, name)
	result := mwer.watchedEvents
	// clear watched events once returned so same events are returned for every filter while testing
	mwer.watchedEvents = []*core.WatchedEvent{}
	return result, nil
}

type MockOasisLogRepository struct {
	logTakes  []log_take.LogTakeEntity
	ethLogIDs []int64
}

func (molr *MockOasisLogRepository) GetLogTakesByMaker(maker string) ([]log_take.LogTakeModel, error) {
	panic("implement me")
}

func (molr *MockOasisLogRepository) CreateLogTake(logTake log_take.LogTakeEntity, ethLogId int64) error {
	molr.logTakes = append(molr.logTakes, logTake)
	molr.ethLogIDs = append(molr.ethLogIDs, ethLogId)
	return nil
}

type MockFilterRepository struct {
}

func (MockFilterRepository) CreateFilter(filter filters.LogFilter) error {
	panic("implement me")
}

func (MockFilterRepository) GetFilter(name string) (filters.LogFilter, error) {
	panic("implement me")
}

var logID1 = int64(123)
var logID2 = int64(456)
var fakeWatchedEvents = []*core.WatchedEvent{
	&core.WatchedEvent{
		LogID:       logID1,
		Name:        "LogTake",
		BlockNumber: 5211385,
		Address:     common.StringToAddress("address").String(),
		TxHash:      "",
		Index:       0,
		Topic0:      common.StringToAddress("topic0").Hex(),
		Topic1:      common.StringToAddress("topic1").Hex(),
		Topic2:      common.StringToAddress("topic2").Hex(),
		Topic3:      common.StringToAddress("topic3").Hex(),
		Data:        common.StringToAddress("data").Hex(),
	},
	&core.WatchedEvent{
		LogID:       logID2,
		Name:        "LogTake",
		BlockNumber: 0,
		Address:     common.StringToAddress("address").String(),
		TxHash:      "",
		Index:       0,
		Topic0:      common.StringToAddress("topic0").Hex(),
		Topic1:      common.StringToAddress("topic1").Hex(),
		Topic2:      common.StringToAddress("topic2").Hex(),
		Topic3:      common.StringToAddress("topic3").Hex(),
		Data:        common.StringToAddress("data").Hex(),
	},
}

var _ = Describe("LogTakeEntity Handler", func() {
	It("persists a record for each log take watched event", func() {
		logTakeConverter := MockLogTakeConverter{}
		watchedEventsRepo := MockWatchedEventsRepository{}
		watchedEventsRepo.SetWatchedEvents(fakeWatchedEvents)
		oasisLogRepo := MockOasisLogRepository{}
		filterRepo := MockFilterRepository{}
		handler := log_take.Handler{
			&logTakeConverter,
			&watchedEventsRepo,
			filterRepo,
			&oasisLogRepo,
		}

		handler.Execute()

		Expect(len(oasisLogRepo.logTakes)).To(Equal(2))
		Expect(oasisLogRepo.ethLogIDs).To(ConsistOf(logID1, logID2))
	})
})
