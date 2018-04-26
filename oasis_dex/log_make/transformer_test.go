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

package log_make_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_make"
	"github.com/8thlight/oasis_watcher/oasis_dex/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

type MockLogMakeConverter struct {
	watchedEvents     []*core.WatchedEvent
	entitiesToConvert []log_make.LogMakeEntity
	block             int64
}

func (mlmc *MockLogMakeConverter) ToModel(entity log_make.LogMakeEntity) log_make.LogMakeModel {
	mlmc.entitiesToConvert = append(mlmc.entitiesToConvert, entity)
	return log_make.LogMakeModel{}
}

func (mlmc *MockLogMakeConverter) ToEntity(watchedEvent core.WatchedEvent) (*log_make.LogMakeEntity, error) {
	mlmc.watchedEvents = append(mlmc.watchedEvents, &watchedEvent)
	e := &log_make.LogMakeEntity{Block: mlmc.block}
	mlmc.block++
	return e, nil
}

var logID1 int64 = 100
var logID2 int64 = 101

var fakeWatchedEvents = []*core.WatchedEvent{
	{
		LogID:       100,
		Name:        "LogMake",
		BlockNumber: 5433832,
		Address:     constants.ContractAddress,
		TxHash:      "0x5f2a91616d1ca67d0761360d33a5b1cf9d46612d165442d9c170307a1ab2e60c",
		Index:       0,
		Topic0:      constants.LogMakeSignature,
		Topic1:      "0x000000000000000000000000000000000000000000000000000000000000a291",
		Topic2:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		Topic3:      "0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64",
		Data:        "000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a232603590000000000000000000000000000000000000000000000022b1c8c1227a0000000000000000000000000000000000000000000000000045aa502b2307e598000000000000000000000000000000000000000000000000000000000005ad0ca29",
	},
	{
		LogID:       101,
		Name:        "LogMake",
		BlockNumber: 5433810,
		Address:     constants.ContractAddress,
		TxHash:      "0x4844a000ffab6e3532f84dcfb6bba2c915f2865c26f1d73e97981e083ba28218",
		Index:       0,
		Topic0:      constants.LogMakeSignature,
		Topic1:      "0x000000000000000000000000000000000000000000000000000000000000a28e",
		Topic2:      "0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391",
		Topic3:      "0x000000000000000000000000004075e4d4b1ce6c48c81cc940e2bad24b489e64",
		Data:        "000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a232603590000000000000000000000000000000000000000000000022b1c8c1227a000000000000000000000000000000000000000000000000004637e539c63c5682000000000000000000000000000000000000000000000000000000000005ad0c881",
	},
}

var _ = Describe("LogMake transformer", func() {
	var logMakeConverter MockLogMakeConverter
	var watchedEventsRepo mocks.MockWatchedEventsRepository
	var repository mocks.MockLogMakeRepo
	var filterRepo mocks.MockFilterRepository
	var transformer log_make.Transformer

	BeforeEach(func() {
		logMakeConverter = MockLogMakeConverter{}
		watchedEventsRepo.SetWatchedEvents(fakeWatchedEvents)
		repository = mocks.MockLogMakeRepo{}

		transformer = log_make.Transformer{
			Converter:              &logMakeConverter,
			WatchedEventRepository: &watchedEventsRepo,
			FilterRepository:       &filterRepo,
			Repository:             &repository,
		}

		transformer.Execute()
	})

	It("calls the watched events repo with the correct filter", func() {
		Expect(len(watchedEventsRepo.Names)).To(Equal(1))
		Expect(watchedEventsRepo.Names[0]).To(Equal("LogMake"))
	})

	It("calls the LogMake converter with the watched event", func() {
		Expect(len(logMakeConverter.watchedEvents)).To(Equal(2))
		Expect(logMakeConverter.watchedEvents).To(ConsistOf(fakeWatchedEvents))
	})

	It("calls ToModel with LogMake entity", func() {
		Expect(len(logMakeConverter.entitiesToConvert)).To(Equal(2))
		Expect(logMakeConverter.entitiesToConvert[0].Block).To(Equal(int64(0)))
		Expect(logMakeConverter.entitiesToConvert[1].Block).To(Equal(int64(1)))
	})

	It("persists a record for each LogMake event", func() {
		Expect(len(repository.LogMakes)).To(Equal(2))
		Expect(repository.VulcanizeLogIDs).To(ConsistOf(logID1, logID2))
	})
})
