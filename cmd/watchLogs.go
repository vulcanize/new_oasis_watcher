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

package cmd

import (
	"log"

	"time"

	"github.com/8thlight/oasis_watcher/oasis_dex"
	"github.com/spf13/cobra"
	"github.com/vulcanize/vulcanizedb/libraries/shared"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

// watchLogsCmd represents the watchLogs command
var watchLogsCmd = &cobra.Command{
	Use:   "watchLogs",
	Short: "Identify all transaction logs and persist data about them.",
	Long: `Creates a filter to watch events where LogTake is invoked in Oasis contracts,
then looks up corresponding logs and persists transaction data to the DB.`,
	Run: func(cmd *cobra.Command, args []string) {
		watchLogs()
	},
}

func init() {
	rootCmd.AddCommand(watchLogsCmd)
}

func watchLogs() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	blockchain := geth.NewBlockchain(ipc)
	db, err := postgres.NewDB(databaseConfig, blockchain.Node())
	if err != nil {
		log.Fatal("Failed to initialize DB")
	}
	watcher := shared.Watcher{
		DB:         *db,
		Blockchain: blockchain,
	}
	watcher.AddTransformers(oasis_dex.TransformerInitializers())
	for range ticker.C {
		watcher.Execute()
	}
}
