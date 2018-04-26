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

package oasis_dex

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_make"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/vulcanize/vulcanizedb/libraries/shared"
)

func TransformerInitializers() []shared.TransformerInitializer {
	return []shared.TransformerInitializer{
		log_take.NewTransformer,
		log_make.NewTransformer,
		log_kill.NewTransformer,
	}
}
