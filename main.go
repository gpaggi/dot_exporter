package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	var address, listen string

	ws, ok := os.LookupEnv("WSADDR")
	if !ok {
		address = "wss://kusama-rpc.polkadot.io/"
	} else {
		address = ws
	}

	la, ok := os.LookupEnv("LISTENADDR")
	if !ok {
		listen = "localhost:9090"
	} else {
		listen = la
	}

	lvl := os.Getenv("LOGLEVEL")
	if lvl != "" {
		lvl, err := log.ParseLevel(lvl)
		if err != nil {
			log.Fatal(err)
		}
		log.SetLevel(lvl)
	}

	cfg := Config{
		address: address,
	}

	prometheus.MustRegister(NewCollector(cfg))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(listen, nil)
}
