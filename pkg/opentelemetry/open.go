package telemetry

import (
	"context"
	"fmt"

	conf "github.com/kowiste/boilerplate/src/config"
	"github.com/kowiste/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func New() (tp *trace.TracerProvider, err error) {
	ctx := context.Background()
	cnf, err := config.Get[conf.BoilerConfig]()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}
	// Configure the OTLP exporter with the provided endpoint
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint(cnf.LogAddress), otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	tp = trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("GinApp"),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp, nil
}

