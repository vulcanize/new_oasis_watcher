package graphql_server

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/log_take"
)

var OasisGraphQLSchema = `
	schema {
		query: Query
	}

	type Query {
		makerHistory(maker: String!): TransactionList
	}

	type TransactionList {
		total: Int!
		transactions: [Transaction]!
	}

	type Transaction {
		id: String!
		pair: String!
		maker: String!
		haveToken: String!
		wantToken: String!
		taker: String!
		takeAmount: String!
		giveAmount: String!
		block: Int!
		timestamp: Int!
	}

`

type GraphQLRepositories struct {
	log_take.IOasisLogRepository
}

type Resolver struct {
	graphQLRepositories GraphQLRepositories
}

func NewResolver(repositories GraphQLRepositories) *Resolver {
	return &Resolver{graphQLRepositories: repositories}
}

func (r *Resolver) MakerHistory(args struct {
	Maker string
}) (*transactionsResolver, error) {
	transactions, err := r.graphQLRepositories.GetLogTakesByMaker(args.Maker)
	if err != nil {
		return &transactionsResolver{}, err
	}
	return &transactionsResolver{transactions: transactions}, err
}

type transactionsResolver struct {
	transactions []log_take.LogTakeModel
}

func (csr transactionsResolver) Transactions() []*transactionResolver {
	return resolveTransactions(csr.transactions)
}

func (csr transactionsResolver) Total() int32 {
	return int32(len(csr.transactions))
}

func resolveTransactions(transactions []log_take.LogTakeModel) []*transactionResolver {
	transactionResolvers := make([]*transactionResolver, 0)
	for _, transaction := range transactions {
		transactionResolvers = append(transactionResolvers, &transactionResolver{transaction})
	}
	return transactionResolvers
}

type transactionResolver struct {
	c log_take.LogTakeModel
}

func (cr transactionResolver) Id() string {
	return cr.c.LogID
}

func (cr transactionResolver) Pair() string {
	return cr.c.Pair
}

func (cr transactionResolver) Maker() string {
	return cr.c.Maker
}

func (cr transactionResolver) HaveToken() string {
	return cr.c.HaveToken
}

func (cr transactionResolver) WantToken() string {
	return cr.c.WantToken
}

func (cr transactionResolver) Taker() string {
	return cr.c.Taker
}

func (cr transactionResolver) TakeAmount() string {
	return cr.c.TakeAmount
}

func (cr transactionResolver) GiveAmount() string {
	return cr.c.GiveAmount
}

func (cr transactionResolver) Block() int32 {
	return int32(cr.c.Block)
}

func (cr transactionResolver) Timestamp() int32 {
	return int32(cr.c.Timestamp)
}
