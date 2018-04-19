package log_kill_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	"github.com/8thlight/oasis_watcher/oasis_dex/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

type MockLogKillConverter struct {
	watchedEvents []*core.WatchedEvent
}

func (mlkc *MockLogKillConverter) ToModel(watchedEvent core.WatchedEvent) (*log_kill.LogKillModel, error) {
	mlkc.watchedEvents = append(mlkc.watchedEvents, &watchedEvent)
	return &log_kill.LogKillModel{}, nil
}

var blockID1 = int64(5428074)
var blockID2 = int64(5428405)

var fakeWatchedEvents = []*core.WatchedEvent{
	{
		LogID:       113,
		Name:        "LogKill",
		BlockNumber: blockID1,
		Address:     constants.ContractAddress,
		TxHash:      "0x769de518d62d3ec4c4c5b50c51ca8248f27f4f5f833f349fc150adc4b2548cfd",
		Index:       0,
		Topic0:      constants.LogKillSignature,
		Topic1:      "0x0000000000000000000000000000000000000000000000000000000000009eda",
		Topic2:      "0x204053929a0ef66ee09fa2295cc078531ee0339fa4d3e02ce9bb3f1a5d0116dd",
		Topic3:      "0x0000000000000000000000009f87bda86354ba26d0e9250d006876d8b5216622",
		Data:        "000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a200000000000000000000000000000000000000000000000003782dace9d90000000000000000000000000000000000000000000000000000028d1286abf261e2000000000000000000000000000000000000000000000000000000005acf7f72",
	},
	{
		LogID:       100,
		Name:        "LogKill",
		BlockNumber: blockID2,
		Address:     constants.ContractAddress,
		TxHash:      "0xd1b94c5745add6e7aedcb9504458ab75860ff8a9acd19e030273d4d23585e2c7",
		Index:       0,
		Topic0:      constants.LogKillSignature,
		Topic1:      "0x0000000000000000000000000000000000000000000000000000000000009f06",
		Topic2:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		Topic3:      "0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64",
		Data:        "000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a232603590000000000000000000000000000000000000000000000022b1c8c1227a000000000000000000000000000000000000000000000000003ed8ff043a5f56f4000000000000000000000000000000000000000000000000000000000005acf91ed",
	},
}

var _ = Describe("LogKill transformer", func() {
	var logKillConverter MockLogKillConverter
	var watchedEventsRepo mocks.MockWatchedEventsRepository
	var logKillRepo mocks.MockLogKillRepo
	var filterRepo mocks.MockFilterRepository

	BeforeEach(func() {
		logKillConverter = MockLogKillConverter{}
		watchedEventsRepo = mocks.MockWatchedEventsRepository{}
		watchedEventsRepo.SetWatchedEvents(fakeWatchedEvents)
		logKillRepo = mocks.MockLogKillRepo{}
		filterRepo = mocks.MockFilterRepository{}
	})

	It("Calls the watched events repo with correct filter", func() {
		transformer := log_kill.Transformer{
			Converter:              &logKillConverter,
			WatchedEventRepository: &watchedEventsRepo,
			FilterRepository:       filterRepo,
			Repository:             &logKillRepo,
		}
		transformer.Execute()
		Expect(len(watchedEventsRepo.Names)).To(Equal(1))
		Expect(watchedEventsRepo.Names).To(ConsistOf("LogKill"))
	})

	It("Calls the LogKill converter with the watched event", func() {
		transformer := log_kill.Transformer{
			Converter:              &logKillConverter,
			WatchedEventRepository: &watchedEventsRepo,
			FilterRepository:       filterRepo,
			Repository:             &logKillRepo,
		}
		transformer.Execute()
		Expect(len(logKillConverter.watchedEvents)).To(Equal(2))
		Expect(logKillConverter.watchedEvents).To(ConsistOf(fakeWatchedEvents))
	})

})
