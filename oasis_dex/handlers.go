package oasis_dex

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/vulcanize/vulcanizedb/libraries/shared"
)

func HandlerInitializers() []shared.HandlerInitializer {
	return []shared.HandlerInitializer{
		log_take.NewLogTakeHandler,
	}
}
