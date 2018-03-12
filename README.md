# Oasis Watcher

## Description
A [VulcanizeDB](https://github.com/vulcanize/VulcanizeDB) handler for watching events related to trading on OasisDex.

## Setup
Setup Postgres and Geth - see [VulcanizeDB README](https://github.com/vulcanize/VulcanizeDB/blob/master/README.md)

Sync the DB from Vulcanize to seed with some log data.

Modify desired config file to point IPC at desired location.

`./oasis_watcher watchLogs --config environments/public.toml`

`./oasis_watcher graphql --config environments/public.toml`

## Notes

There is a possibility for some dependency resolution errors if dependencies are resolving to a vendored file in another project.
A temporary workaround is to temporarily blow away the offending vendored dependency.
In the future, we hope to eliminate the offending vendored dependencies and consistently depend only on go-ethereum.

