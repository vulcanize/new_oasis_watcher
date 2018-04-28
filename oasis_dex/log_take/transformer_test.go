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

package log_take_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/8thlight/oasis_watcher/oasis_dex/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

type MockLogTakeConverter struct {
	watchedEvents     []*core.WatchedEvent
	convertedIDValues []*uint64
	entitiesToConvert []log_take.LogTakeEntity
	toModelCount      int64
	toEntityCount     int64
}

func (mltc *MockLogTakeConverter) ToModel(entity log_take.LogTakeEntity) log_take.LogTakeModel {
	mltc.entitiesToConvert = append(mltc.entitiesToConvert, entity)
	m := log_take.LogTakeModel{ID: mltc.toModelCount}
	mltc.toModelCount++
	return m
}

func (mltc *MockLogTakeConverter) ToEntity(watchedEvent core.WatchedEvent) (*log_take.LogTakeEntity, error) {
	mltc.watchedEvents = append(mltc.watchedEvents, &watchedEvent)
	e := &log_take.LogTakeEntity{Block: mltc.toEntityCount}
	mltc.toEntityCount++
	return e, nil
}

var logID1 = int64(123)
var logID2 = int64(456)

var fakeWatchedEvents = []*core.WatchedEvent{
	{
		LogID:       logID1,
		Name:        "LogTake",
		BlockNumber: 5430136,
		Address:     constants.ContractAddress,
		TxHash:      "0xaca917ef9440aaf2d37cd36309872ce8ab6251f56cac62524b6eb63d5c891be8",
		Index:       8,
		Topic0:      constants.LogTakeSignature,
		Topic1:      "0x204053929a0ef66ee09fa2295cc078531ee0339fa4d3e02ce9bb3f1a5d0116dd",
		Topic2:      "0x000000000000000000000000168910909606a2fca90d4c28fa39b50407b9c526",
		Topic3:      "0x0000000000000000000000000016bd4cb70bd98ca07a341da66450b5d22a55aa",
		Data:        "0000000000000000000000000000000000000000000000000000000000009bef000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a200000000000000000000000000000000000000000000000003ac43f17ce74b6f00000000000000000000000000000000000000000000000002bdb0cb23e5ebe0000000000000000000000000000000000000000000000000000000005acff7c4",
	},
	{
		LogID:       logID2,
		Name:        "LogTake",
		BlockNumber: 5430139,
		Address:     constants.ContractAddress,
		TxHash:      "0x764bd5e6127d263140b7835920cbc3cb28ca67ce62b73a04aff569e1fa75423c",
		Index:       54,
		Topic0:      constants.LogTakeSignature,
		Topic1:      "0x204053929a0ef66ee09fa2295cc078531ee0339fa4d3e02ce9bb3f1a5d0116dd",
		Topic2:      "0x000000000000000000000000168910909606a2fca90d4c28fa39b50407b9c526",
		Topic3:      "0x0000000000000000000000000016bd4cb70bd98ca07a341da66450b5d22a55aa",
		Data:        "0000000000000000000000000000000000000000000000000000000000009bef000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a200000000000000000000000000000000000000000000000003ac43f17ce74b6f00000000000000000000000000000000000000000000000002bdb0cb23e5ebe0000000000000000000000000000000000000000000000000000000005acff7dd",
	},
}

var _ = Describe("LogTake transformer", func() {
	var mockLogTakeConverter *MockLogTakeConverter
	var watchedEventsRepo *mocks.MockWatchedEventsRepository
	var oasisLogRepo *mocks.MockLogTakeRepo
	var filterRepo mocks.MockFilterRepository
	var logMakeRepo *mocks.MockLogMakeRepo

	BeforeEach(func() {
		mockLogTakeConverter = &MockLogTakeConverter{}
		watchedEventsRepo = &mocks.MockWatchedEventsRepository{}
		watchedEventsRepo.SetWatchedEvents(fakeWatchedEvents)
		oasisLogRepo = &mocks.MockLogTakeRepo{}
		filterRepo = mocks.MockFilterRepository{}
		logMakeRepo = &mocks.MockLogMakeRepo{}

		transformer := log_take.Transformer{
			Converter:              mockLogTakeConverter,
			WatchedEventRepository: watchedEventsRepo,
			FilterRepository:       filterRepo,
			LogMakeRepository:      logMakeRepo,
			Repository:             oasisLogRepo,
		}
		transformer.Execute()

	})

	It("calls the watched events repo with correct filter", func() {
		Expect(len(watchedEventsRepo.Names)).To(Equal(4))
		Expect(watchedEventsRepo.Names).To(ConsistOf("LogTake-0x83", "LogTake-0x14", "LogTake-0x3a", "LogTake-0x91"))
	})

	It("calls the LogTake converter with the watched events", func() {
		Expect(len(mockLogTakeConverter.watchedEvents)).To(Equal(2))
		Expect(mockLogTakeConverter.watchedEvents).To(ConsistOf(fakeWatchedEvents))
	})

	It("converts an LogTake entity to a model", func() {
		Expect(len(mockLogTakeConverter.entitiesToConvert)).To(Equal(2))
		Expect(mockLogTakeConverter.entitiesToConvert[0].Block).To(Equal(int64(0)))
		Expect(mockLogTakeConverter.entitiesToConvert[1].Block).To(Equal(int64(1)))
	})

	It("creates a LogTake record for each event", func() {
		Expect(len(oasisLogRepo.LogTakes)).To(Equal(2))
		Expect(oasisLogRepo.VulcanizeLogIDs).To(ConsistOf(logID1, logID2))
	})

})
