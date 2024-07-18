package telemetry

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryService struct {
	tracer trace.Tracer
}

func NewTelemetryService() *TelemetryService {
	return &TelemetryService{
		tracer: otel.Tracer("orderapp-fiber-server"),
	}
}

func (t *TelemetryService) Tracer() trace.Tracer {
	return t.tracer
}

func (t *TelemetryService) InitTracer() *sdktrace.TracerProvider {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("my-service"),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func (t *TelemetryService) OtelFiberMiddleware(c *fiber.Ctx) error {
	ctx := c.UserContext()
	spanCtx, span := t.tracer.Start(ctx, c.Path())
	defer span.End()

	c.SetUserContext(spanCtx)
	return c.Next()
}

// func (t *TelemetryService) TraceMiddleware(handler fiber.Handler) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		ctx := c.UserContext()
// 		_, span := t.tracer.Start(ctx, c.Route().Path)
// 		defer span.End()

// 		c.SetUserContext(ctx)
// 		return handler(c)
// 	}
// }
