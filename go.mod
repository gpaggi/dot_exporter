module github.com/gpaggi/dot_monitor

go 1.15

require (
	github.com/centrifuge/go-substrate-rpc-client/v3 v3.0.0
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
)

replace github.com/centrifuge/go-substrate-rpc-client/v3 => ./dep/github.com/ParthDesai/go-substrate-rpc-client/v3
