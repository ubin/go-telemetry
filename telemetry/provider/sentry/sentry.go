package sentry

import (
	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"github.com/ubin/go-telemetry/telemetry/config"
	"go.opentelemetry.io/otel/sdk/trace"
)

// InitSentry initializes Sentry for error monitoring
func Init(cfg config.Config) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.GetCollectorEndpoint(),
		EnableTracing:    cfg.IsInsecure(),
		Environment:      cfg.GetEnvironment(),
		TracesSampleRate: 1.0, // Adjust for performance (percentage of transactions to capture)
		Debug:            cfg.IsDebugMode(),
	})
	if err != nil {
		return err
	}
	return nil
}

func TracerProvider() *trace.TracerProvider {
	tracerProvider := trace.NewTracerProvider(
		trace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)
	return tracerProvider
}
