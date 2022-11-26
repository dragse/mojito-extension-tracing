package mojito_extension_tracing

import (
	"context"
	"github.com/go-mojito/mojito"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log"
)

func StartTracing(mojitoCtx mojito.Context, name string) trace.Span {
	var tracingCtx context.Context
	var ok bool
	tracingCtxInterface := mojitoCtx.Metadata().GetOrDefault(METADATA_NAME, context.Background())

	tracingCtx, ok = tracingCtxInterface.(context.Context)

	if !ok {
		// problem while converting context. use a new Context now
		log.Println("TEST!!!")
		tracingCtx = context.Background()
	}

	ctx, span := otel.Tracer(serviceName).Start(tracingCtx, name)
	_ = mojitoCtx.Metadata().Set(METADATA_NAME, ctx)

	return span
}
