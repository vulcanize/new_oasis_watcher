#! /bin/bash
set -e
echo "Setting up Vulcanize"
trap 'kill $(jobs -pr)' SIGINT SIGTERM EXIT
make setup NAME=vulcanize_public
make migrate NAME=vulcanize_public
make build

vulcanizedb sync --config environments/public.toml --starting-block-number 5410136 &
echo "Running Watcher"
./oasis_watcher watchLogs --config environments/public.toml
