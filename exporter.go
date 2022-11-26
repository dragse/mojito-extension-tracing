package mojito_extension_tracing

type Exporter string

const (
	JAEGER         Exporter = "jaeger"
	PROMETHEUS     Exporter = "prometheus"
	OTEL_COLLECTOR Exporter = "otel-collector"
	OPENCENSUS     Exporter = "opencensus"
)
