# Zond Beacon Chain Explorer

The explorer provides a comprehensive and easy to use interface for the Zond beacon chain. It makes it easy to view proposed blocks, follow attestations and monitor your staking activity.

## About

The explorer is built using golang and utilizes a PostgreSQL database for storing and indexing data. 

**Warning:** The explorer is still under heavy active development. More or less everything might change without prior notice and we cannot guarantee any backwards compatibility for now. Once the Zond ecosystem matures we will be able to provide stronger guarantees about the updatability of the explorer.

## Features

- General
  - Open Source (GNU General Public License v3.0)
  - Supports Execution Layer and Consensus Layer
  - Supports multiple networks
  - Written in Golang

- Website
  - Validator Dashboard with status, income, balance, attestations, proposals and charts
  - Overviews about blocks, slots, epochs, transactions, validators, slashings and the mempool
  - Stats and info about Zond clients
  - Charts about various stats

- Tools
  - Income History
  - Profit Calculator
  - Block Visualizer
  - Unit Converter

## Getting started

We currently do not provide any pre-built binaries of the explorer. Docker images are available at https://hub.docker.com/repository/docker/qrledger/zond-beaconchain-explorer.

- Download the latest version of the Qrysm beacon chain client and start it with the `--archive` flag set
- Wait till the client finishes the initial sync
- Setup a PostgreSQL DB and import the `tables.sql` file from the root of this repository
- Install go version 1.13 or higher
- Clone the repository and run `make all` to build the indexer and front-end binaries
- Copy the config-example.yml file and adapt it to your environment
- Start the explorer binary and pass the path to the config file as argument
- To build bootstrap run `npm run --prefix ./bootstrap dist-css` in project folder.

## Developing locally with docker

- Clone the repository
- Run `docker-compose up` to start instances of the following containers `gzond`, `qrysm`, `postgres` and `golang`.
- Open a new terminal in project directory and run `docker run -it --rm --net=host -v $(pwd):/src  postgres psql -f /src/tables.sql -d db -h 0.0.0.0 -U postgres` to create new tables in the database  
- Wait for the client to finish initial sync, you can check this by looking at logs of `qrysm` instance.
- Copy the `config-example.yml` file and adapt it to your environment.\
 For database information check `postgres` section in `docker-compose.yml` file.
- Connect to `golang` instance by running `docker exec -ti golang bash` and run `make all`
- Start the explorer binary and pass the path to the config file as argument 

      ./bin/explorer --config your_config.yml   

## Commercial usage

The explorer uses Highsoft charts which are not free for commercial and governmental use. If you plan to use the explorer for commercial purposes you currently need to purchase an appropriate HighSoft license.
We are planning to switch out the Highsoft chart library with a less restrictive charting library (suggestions are welcome).
