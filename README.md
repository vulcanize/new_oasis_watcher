# Oasis Watcher

## Description
A [VulcanizeDB](https://github.com/vulcanize/VulcanizeDB) transformer for watching events related to trading on OasisDex.

## Setup
```
make setup NAME=vulcanize_public
make migrate NAME=vulcanize_public
make build
make fixlibcrypto
```

Sync the DB from Vulcanize to seed with some log data.

Modify desired config file to point IPC at desired location.

`./oasis_watcher watchLogs --config environments/public.toml`

`./oasis_watcher graphql --config environments/public.toml`
