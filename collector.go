package main

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v3"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	address string
}

type Collector struct {
	activeValidators      *prometheus.Desc
	scrapeFailures        *prometheus.Desc
	cfg                   Config
	scrapeFailuresCount   float64
	activeValidatorsCount int
}

type metric struct {
	validator string
	value     float64
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.activeValidators
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	api, err := gsrpc.NewSubstrateAPI(c.cfg.address)
	if err != nil {
		log.Warn(err)
		c.scrapeFailuresCount++
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Warn(err)
		c.scrapeFailuresCount++
	}

	var validators []types.AccountID
	key, err := types.CreateStorageKey(meta, "Session", "Validators", nil)
	_, err = api.RPC.State.GetStorageLatest(key, &validators)
	if err != nil {
		log.Warn(err)
		c.scrapeFailuresCount++
	}

	c.activeValidatorsCount = len(validators)
	log.Debugf("There are %d validators", c.activeValidatorsCount)

	ch <- prometheus.MustNewConstMetric(
		c.scrapeFailures,
		prometheus.CounterValue,
		c.scrapeFailuresCount,
	)

	ch <- prometheus.MustNewConstMetric(
		c.activeValidators,
		prometheus.CounterValue,
		float64(c.activeValidatorsCount),
	)
}

func NewCollector(cfg Config) prometheus.Collector {
	return &Collector{
		cfg:              cfg,
		activeValidators: prometheus.NewDesc("dot_monitor_active_validators_count", "Count of active validators", nil, nil),
		scrapeFailures:   prometheus.NewDesc("dot_monitor_scrape_failures", "Number of times we failed to scrape", nil, nil),
	}
}
