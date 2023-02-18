package bootstrap

import (
	"errors"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/zipkin"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// NewJaegerExporter 创建一个jaeger导出器
func NewJaegerExporter(endpoint string) (traceSdk.SpanExporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
}

// NewZipkinExporter 创建一个zipkin导出器
func NewZipkinExporter(endpoint string) (traceSdk.SpanExporter, error) {
	return zipkin.New(endpoint)
}

// NewTracerExporter 创建一个导出器，支持：jaeger和zipkin
func NewTracerExporter(exporterName, endpoint string) (traceSdk.SpanExporter, error) {
	switch exporterName {
	case "jaeger":
		return NewJaegerExporter(endpoint)
	case "zipkin":
		return NewZipkinExporter(endpoint)
	default:
		return nil, errors.New("exporter type not support")
	}
}

// NewTracerProvider 创建一个链路追踪器
func NewTracerProvider(exporterName, endpoint, devEnv string, serviceInfo *ServiceInfo) error {
	opts := []traceSdk.TracerProviderOption{
		traceSdk.WithSampler(traceSdk.ParentBased(traceSdk.TraceIDRatioBased(1.0))),
		traceSdk.WithResource(resource.NewSchemaless(
			semConv.ServiceNameKey.String(serviceInfo.Name),
			semConv.ServiceVersionKey.String(serviceInfo.Version),
			semConv.ServiceInstanceIDKey.String(serviceInfo.Id),
			attribute.String("env", devEnv),
		)),
	}

	if len(endpoint) > 0 {
		exp, err := NewTracerExporter(exporterName, endpoint)
		if err != nil {
			panic(err)
		}

		opts = append(opts, traceSdk.WithBatcher(exp))
	}

	tp := traceSdk.NewTracerProvider(opts...)
	if tp == nil {
		return errors.New("create tracer provider failed")
	}

	otel.SetTracerProvider(tp)

	return nil
}
