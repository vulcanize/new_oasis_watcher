package cmd

import (
	"log"

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
	blockchain := geth.NewBlockchain(ipc)
	db, err := postgres.NewDB(databaseConfig, blockchain.Node())
	if err != nil {
		log.Fatal("Failed to initialize DB")
	}
	watcher := shared.Watcher{
		DB:         *db,
		Blockchain: blockchain,
	}
	watcher.AddHandlers(oasis_dex.HandlerInitializers())
	watcher.Execute()
}
