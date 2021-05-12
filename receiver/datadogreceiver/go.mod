module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver

go 1.14

require (
	github.com/DataDog/datadog-agent/pkg/trace/exportable v0.0.0-20201016145401-4646cf596b02
	github.com/gorilla/mux v1.8.0
	github.com/tinylib/msgp v1.1.2
	go.opentelemetry.io/collector v0.26.0
)
