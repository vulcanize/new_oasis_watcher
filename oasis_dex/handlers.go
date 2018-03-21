package oasis_dex

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/vulcanize/vulcanizedb/libraries/shared"
)

func TransformerInitializers() []shared.TransformerInitializer {
	return []shared.TransformerInitializer{
		log_take.NewLogTakeTransformer,
	}
}
