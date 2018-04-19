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
