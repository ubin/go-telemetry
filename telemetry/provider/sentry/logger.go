package sentry

import (
	"log/slog"

	"github.com/getsentry/sentry-go"
)

// sends slog messages to sentry as events
func CaptureLogMessage(r slog.Record) {

	// Send logs as messages to Sentry
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetExtra("level", r.Level.String())
		scope.SetExtra("message", r.Message)
		r.Attrs(func(a slog.Attr) bool {
			scope.SetExtra(a.Key, a.Value.Any())
			return true
		})
		sentry.CaptureMessage(r.Message)
	})

}
