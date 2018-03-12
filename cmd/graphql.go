package cmd

import (
	"log"
	"net/http"

	"github.com/8thlight/oasis_watcher/graphql_server"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	// Here lie headaches - graphql-go was recently transferred from neelance to graph-gophers. Attempting to use the
	// neelance package will result in the error "use of internal package not allowed" since the neelance repo now
	// points to internal packages on the new repo
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/spf13/cobra"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

// graphqlCmd represents the graphql command
var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "Run GraphQL to inspect logged transactions",
	Long: `Run GraphQL on to inspect logged transactions. After running the command, you can visit port 9090 and run, for example:

{
	makerHistory(maker: "0x0005ABcBB9533Cf6F9370505ffeF25393E0D2852") {
    transactions {
      id
      maker
      taker
      haveToken
      wantToken
      pair
    }
    total
  }
}

This will return the total number of transactions and transaction details for all logs with maker 0x0005ABcBB9533Cf6F9370505ffeF25393E0D2852`,
	Run: func(cmd *cobra.Command, args []string) {
		schema := parseSchema()
		serve(schema)
	},
}

func init() {
	rootCmd.AddCommand(graphqlCmd)
}

func parseSchema() *graphql.Schema {

	blockchain := geth.NewBlockchain(ipc)
	db, err := postgres.NewDB(databaseConfig, blockchain.Node())
	if err != nil {
		log.Fatal("Can't connect to db")
	}
	oasisLogRepository := log_take.OasisLogRepository{db}
	graphQLRepositories := graphql_server.GraphQLRepositories{
		IOasisLogRepository: oasisLogRepository,
	}
	schema := graphql.MustParseSchema(graphql_server.OasisGraphQLSchema, graphql_server.NewResolver(graphQLRepositories))
	return schema

}

func serve(schema *graphql.Schema) {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":9090", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
