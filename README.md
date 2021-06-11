# Dot monitor
A Prometheus exporter to export metrics related to Substrate based Nodes.

## Table Of Contents

* [Intro](#intro)
* [Requirements](#requirements)
* [How To Build](#how-to-build)
* [Known Limitations](#known-limitations)
* [Versioning](#Versioning)


## Intro
Metrics are exported in Prometheus format.  
The following env vars can be used for configuration:
```
WSADDR = "wss://kusama-rpc.polkadot.io/"
LISTENADDR = "localhost:9090"
LOGLEVEL = INFO
```
Example of the metrics:
```
# HELP dot_monitor_active_validators_count Count of active validators
# TYPE dot_monitor_active_validators_count counter
dot_monitor_active_validators_count 900
# HELP dot_monitor_scrape_failures Number of times we failed to scrape
# TYPE dot_monitor_scrape_failures counter
dot_monitor_scrape_failures 0
```
In case of failures scraping `dot_monitor_scrape_failures` will be incremented.

## Requirements
* Golang (for local build)
* Docker

## How To Build
Build locally:
```
make build
```

Build with Docker:
```
make build-docker
```

Tag and push to the Docker registry:
```
make distribute
```

## How To Run
```
make build
./dot_exporter
```

## Known Limitations
* Calls to gsrpc should be passed a context with deadline to avoid the whole exporter hanging in case of issues while calling
* In this implementation, and following Prometheus development's guidelines, the metrics collection scheduling is left to Prometheus. In future implementations, and when handling large input list of validators, it is recommended to collect the metrics asynchronously.

## Versioning

For the versions available see git history of the VERSION file