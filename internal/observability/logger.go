package observability

import (
	"context"
	"fmt"
	"log/slog"
)

type CoreLogger struct {
	*slog.Logger
}

const LevelFatal = slog.Level(12)

// CaptureFatal logs a fatal error and sends it to Sentry.
func (cl *CoreLogger) CaptureFatal(err error, args ...any) {
	cl.Logger.Log(context.Background(), LevelFatal, err.Error(), args...)

	// if cl.captureException != nil {
	// 	cl.captureException(err, cl.tagsWithArgs(args...))
	// }
}

// CaptureFatalAndPanic logs a fatal error, sends it to Sentry and panics.
func (cl *CoreLogger) CaptureFatalAndPanic(err error, args ...any) {
	cl.CaptureFatal(err, args...)
	if err != nil {
		panic(err)
	}
}

// CaptureWarn logs a warning and sends it to Sentry.
func (cl *CoreLogger) CaptureWarn(msg string, args ...any) {
	cl.Logger.Warn(msg, args...)

	// if cl.captureMessage != nil {
	// 	cl.captureMessage(msg, cl.tagsWithArgs(args...))
	// }
}

// CaptureInfo logs an info message and sends it to Sentry.
func (cl *CoreLogger) CaptureInfo(msg string, args ...any) {
	cl.Logger.Info(msg, args...)

	// if cl.captureMessage != nil {
	// 	cl.captureMessage(msg, cl.tagsWithArgs(args...))
	// }
}

// Reraise reports panics to Sentry.
func (cl *CoreLogger) Reraise(args ...any) {
	// if err := recover(); err != nil {
	// 	cl.reraise(err, cl.tagsWithArgs(args...))
	// }
}

// GetTags returns the tags associated with the logger.
// func (cl *CoreLogger) GetTags() Tags {
// 	return cl.globalTags
// }

// GetLogger returns the underlying slog.Logger.
func (cl *CoreLogger) GetLogger() *slog.Logger {
	return cl.Logger
}

// GetCaptureException returns the function used to capture exceptions.
// func (cl *CoreLogger) GetCaptureException() func(err error, tags map[string]string) {
// 	return cl.captureException
// }

// GetCaptureMessage returns the function used to capture messages.
// func (cl *CoreLogger) GetCaptureMessage() func(msg string, tags map[string]string) {
// 	return cl.captureMessage
// }

type Printer struct{}

func (p *Printer) Write(text string) {
	fmt.Print(text)
}
