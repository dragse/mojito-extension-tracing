package mojito_extension_tracing

import (
	"context"
	"github.com/go-mojito/mojito"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func Middleware(ctx mojito.Context, next func() error) error {
	contextCtx, span := otel.Tracer(serviceName).Start(context.Background(), "Handle HTTP Request")
	span.SetAttributes(
		attribute.String("host", ctx.Request().GetRequest().Host),
		attribute.String("method", ctx.Request().GetRequest().Method),
		attribute.String("uri", ctx.Request().GetRequest().RequestURI),
		attribute.String("user-agent", ctx.Request().GetRequest().UserAgent()),
	)
	defer span.End()
	_ = ctx.Metadata().Set(METADATA_NAME, contextCtx)

	span.AddEvent("Execute next() Function")
	err := next()
	span.AddEvent("next() Function complete")

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}
