package mojito_extension_tracing

import (
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"log"
)

var (
	serviceName string         = "go-mojito"
	envList     map[string]any = make(map[string]any)
)

type ExporterConfig struct {
	ProviderURL string
}

func Configure(exporterType Exporter, name string, env map[string]any, config ExporterConfig) error {
	serviceName = name
	envList = env

	switch exporterType {
	case JAEGER:
		provider, err := configureJaegerTracer(config.ProviderURL)

		if err != nil {
			return err
		}

		err = ConfigureWithCustomTracerProvider(provider)

		if err != nil {
			return err
		}
	}

	return nil
}

func ConfigureWithCustomTracerProvider(tracer *tracesdk.TracerProvider) error {
	otel.SetTracerProvider(tracer)

	return nil
}

func configureMapToSlice() []attribute.KeyValue {
	values := make([]attribute.KeyValue, 0)

	for key, val := range envList {
		switch val.(type) {
		case string:
			values = append(values, attribute.String(key, val.(string)))
		case bool:
			values = append(values, attribute.Bool(key, val.(bool)))
		case int64:
			values = append(values, attribute.Int64(key, val.(int64)))
		case int:
			values = append(values, attribute.Int(key, val.(int)))
		default:
			log.Println(fmt.Sprintf("key '%s' with datatype '%T' currently not supported", key, val))
		}
	}

	return values
}
