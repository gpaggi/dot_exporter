# Dot monitor
A Prometheus exporter to export metrics related to Substrate based Nodes.

## Table Of Contents

* [Requirements](#requirements)
* [How To Build](#how-to-build)


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



Exports metrics about invalid events in Prometheus format, based on existing Flink metrics in Odin's Prometheus.

Expected env vars:
```
LISTENADDR = ":${NOMAD_PORT_http}"
PROMADDR = "http://prometheus.service.ams1.consul:9190/api/v1/query"
# the max_over_time is required to make up for failed scrapes of Flink
QUERY = "sum by (tenant)(max_over_time(flink_odin_invalid_messages[5m]))"
```
For basic auth pass also:
```
USERNAME = "foo"
PASSWORD = "bar"
```

Example of the metrics:
```
odin_invalid_events_total{tenant="comaas-ebay-kleinanzeigen"} 0
odin_invalid_events_total{tenant="comaas-kijiji-ca"} 0
odin_invalid_events_total{tenant="comaas-marktplaats"} 0
odin_invalid_events_total{tenant="comaas-move-ca"} 0
odin_invalid_events_total{tenant="denmark"} 18
odin_invalid_events_total{tenant="ebay-kleinanzeigen"} 0
odin_invalid_events_total{tenant="gumtree-au"} 0
odin_invalid_events_total{tenant="gumtree-uk"} 4
odin_invalid_events_total{tenant="icas-kijiji-ca"} 0
odin_invalid_events_total{tenant="icas-marktplaats"} 86
odin_invalid_events_total{tenant="icas-twh"} 0
odin_invalid_events_total{tenant="kijiji-ca"} 336
odin_invalid_events_total{tenant="marktplaats"} 17
odin_invalid_events_total{tenant="mobile-de"} 338559
odin_invalid_events_total{tenant="move-au"} 0
odin_invalid_events_total{tenant="move-ca"} 74635
odin_invalid_events_total{tenant="total"} 413655
odin_invalid_events_total{tenant="twh"} 0
odin_prometheus_scrape_failures 0
```
In case of failures scraping Prometheus `odin_prometheus_scrape_failures` will be incremented.

It requires the following relabling in the Prometheus configuration:
```
metric_relabel_configs:
  - source_labels: [__name__]
    regex: "flink_taskmanager_job_task_operator_([a-z_]+)_invalid_messages"
    target_label: tenant
    replacement: '${1}'
  - source_labels: [__name__]
    regex: "flink_taskmanager_job_task_operator_.*_invalid_messages"
    target_label: __name__
    replacement: 'flink_odin_invalid_messages'
```
