# Oasis Watcher

## Description
A [VulcanizeDB](https://github.com/vulcanize/VulcanizeDB) transformer for watching events related to trading on OasisDex.

## Dependencies
 - Go 1.9+
 - Postgres 10
 - Ethereum Node
   - [Go Ethereum](https://ethereum.github.io/go-ethereum/downloads/) (1.8+)
   - [Parity 1.8.11+](https://github.com/paritytech/parity/releases)
   - [IPFS](https://github.com/ipfs/go-ipfs#build-from-source)

## Installation
1. Setup Postgres and an Ethereum node - see [VulcanizeDB README](https://github.com/vulcanize/VulcanizeDB/blob/master/README.md).
1. `git clone git@github.com:8thlight/oasis_watcher.git`

  _note: `go get` does not work for this project because need to run the (fixlibcrypto)[https://github.com/8thlight/oasis_watcher/blob/master/Makefile] command along with `go build`._
1. Create the database based on the [VulcanizeDB schema](https://github.com/vulcanize/VulcanizeDB/blob/master/db/schema.sql):
    ```
    make setup NAME=vulcanize_public
    ```
1. Run the migrations to add project specific tables to the database:
    ```
    make migrate NAME=vulcanize_public
    ```
1. Build:
    ```
    make build
    ```

## Configuration
- To use a local Ethereum node, copy `environments/public.toml.example` to
  `environments/public.toml` and update the `ipcPath` to the local node's IPC filepath:
  - when using geth:
    - The IPC file is called `geth.ipc`.
    - The geth IPC file path is printed to the console when you start geth.
    - The default location is:
      - Mac: `$HOME/Library/Ethereum`
      - Linux: `$HOME/.ethereum`

  - when using parity:
    - The IPC file is called `jsonrpc.ipc`.
    - The default location is:
      - Mac: `$HOME/Library/Application\ Support/io.parity.ethereum/`
      - Linux: `$HOME/.local/share/io.parity.ethereum/`

- See `environments/infura.toml` to configure commands to run against infura, if a local node is unavailable.

## Running the watchLogs command
`watchLogs` runs a process that gets log events associated with specified Oasis contracts, and then stores transformed values in a `oasis.log_takes` table, and a `oasis.log_takes_with_status` view in the VulcanizeDB database:

1. Run the [VulcanizeDB `sync` command](https://github.com/vulcanize/VulcanizeDB#start-syncing-with-postgres).
1. `./oasis_watcher watchLogs --config <environment config file>`

## GraphQL
We're using [PostGraphile](https://www.graphile.org/postgraphile/) to create a GraphQL API from the VulcanizeDB postgres schema.
1. Ensure that Node.js v8.6 is installed.
1. Install postgraphile:
    ```
    npm install -g postgraphile
    ```
1. Start the postgraphile server:
    ```
    postgraphile -c "postgresql://<user>@localhost:5432/vulcanize_public" --schema=public,oasis
    ```
    - the `-c "postgresql://user@localhost:5432/vulcanize_public"` flag indicates which postgres connection postgraphile should be looking to, where `<user>` is your local postgres user
    - the `--schema=public,oasis` flag indicates which schema(s) postgraphile should use to generate the GraphQL API
