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

package mocks

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_make"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/filters"
)

type MockWatchedEventsRepository struct {
	watchedEvents []*core.WatchedEvent
	Names         []string
}

func (mwer *MockWatchedEventsRepository) SetWatchedEvents(watchedEvents []*core.WatchedEvent) {
	mwer.watchedEvents = watchedEvents
}

func (mwer *MockWatchedEventsRepository) GetWatchedEvents(name string) ([]*core.WatchedEvent, error) {
	mwer.Names = append(mwer.Names, name)
	result := mwer.watchedEvents
	// clear watched events once returned so same events are returned for every filter while testing
	mwer.watchedEvents = []*core.WatchedEvent{}
	return result, nil
}

type MockLogMakeRepo struct {
	LogMakes        []log_make.LogMakeModel
	VulcanizeLogIDs []int64
}

func (molr *MockLogMakeRepo) Create(offerModel log_make.LogMakeModel, vulcanizeLogId int64) error {
	molr.LogMakes = append(molr.LogMakes, offerModel)
	molr.VulcanizeLogIDs = append(molr.VulcanizeLogIDs, vulcanizeLogId)
	return nil
}

type MockLogKillRepo struct {
	LogKills        []log_kill.LogKillModel
	VulcanizeLogIDs []int64
}

func (molk *MockLogKillRepo) Create(model log_kill.LogKillModel, vulcanizeLogID int64) error {
	molk.LogKills = append(molk.LogKills, model)
	molk.VulcanizeLogIDs = append(molk.VulcanizeLogIDs, vulcanizeLogID)
	return nil
}

type MockLogTakeRepo struct {
	LogTakes        []log_take.LogTakeModel
	VulcanizeLogIDs []int64
}

func (molr *MockLogTakeRepo) Create(logTake log_take.LogTakeModel, vulcanizeLogID int64) error {
	molr.LogTakes = append(molr.LogTakes, logTake)
	molr.VulcanizeLogIDs = append(molr.VulcanizeLogIDs, vulcanizeLogID)
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
