package graphql_server_test

import (
	"context"
	"log"

	"encoding/json"

	"math/big"

	"fmt"

	"github.com/8thlight/oasis_watcher/graphql_server"
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
	"github.com/ethereum/go-ethereum/common"
	"github.com/graph-gophers/graphql-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

func formatJSON(data []byte) []byte {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		log.Fatalf("invalid JSON: %s", err)
	}
	formatted, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return formatted
}

var _ = Describe("GraphQL", func() {
	var graphQLRepositories graphql_server.GraphQLRepositories
	var id = [32]byte{1, 1, 1, 1}
	var pair = [32]byte{2, 2, 2, 2}
	var maker = common.StringToAddress("Maker")
	var haveToken = common.StringToAddress("HaveToken")
	var wantToken = common.StringToAddress("WantToken")
	var taker = common.StringToAddress("Taker")
	var takeAmount = big.NewInt(123)
	var giveAmount = big.NewInt(456)
	var timestamp = uint64(789)

	BeforeEach(func() {
		node := core.Node{GenesisBlock: "GENESIS", NetworkID: 1, ID: "x123", ClientName: "geth"}
		db, err := postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, node)
		Expect(err).NotTo(HaveOccurred())
		db.Query(`DELETE FROM oasis.log_takes`)
		graphQLRepositories = graphql_server.GraphQLRepositories{
			log_take.OasisLogRepository{db},
		}
		logTake := log_take.LogTakeEntity{
			Id:         id,
			Pair:       pair,
			Maker:      maker,
			HaveToken:  haveToken,
			WantToken:  wantToken,
			Taker:      taker,
			TakeAmount: takeAmount,
			GiveAmount: giveAmount,
			Timestamp:  timestamp,
		}
		lr := repositories.LogRepository{
			DB: db,
		}
		err = lr.CreateLogs([]core.Log{{}})
		Expect(err).ToNot(HaveOccurred())
		var ethLogID int64
		err = lr.Get(&ethLogID, `Select id from logs`)
		Expect(err).NotTo(HaveOccurred())
		err = graphQLRepositories.CreateLogTake(logTake, ethLogID)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Queries example schema for maker transactions", func() {
		var variables map[string]interface{}
		resolver := graphql_server.NewResolver(graphQLRepositories)
		var schema = graphql.MustParseSchema(graphql_server.OasisGraphQLSchema, resolver)

		response := schema.Exec(context.Background(),
			`{
	                      makerHistory(maker: "0x0000000000000000000000000000004d616B6572") {
	                       total
	                       transactions{
	                           id
	                           pair
	                           maker
	                           haveToken
	                           wantToken
	                           taker
	                           takeAmount
	                           giveAmount
	                           timestamp
	                         }
	                       }
	                   }`,
			"",
			variables)

		expected := fmt.Sprintf(`{
	                  "makerHistory":
	                     {
	                       "total": 1,
	                       "transactions": [
	                           {
								"id": "%v",
								"pair": "%v",
								"maker": "%v",
								"haveToken": "%v",
								"wantToken": "%v",
								"taker": "%v",
								"takeAmount": "%d",
								"giveAmount": "%d",
								"timestamp": %d
	                           }
	                       ]
	                     }
	               }`, common.Bytes2Hex(id[:]), common.Bytes2Hex(pair[:]), maker.Hex(), haveToken.Hex(), wantToken.Hex(), taker.Hex(), takeAmount, giveAmount, timestamp)
		var v interface{}
		if len(response.Errors) != 0 {
			log.Fatal(response.Errors)
		}
		err := json.Unmarshal(response.Data, &v)
		Expect(err).ToNot(HaveOccurred())
		actualJSON := formatJSON(response.Data)
		expectedJSON := formatJSON([]byte(expected))
		Expect(actualJSON).To(Equal(expectedJSON))
	})
})
