package slogger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type cloudHandler struct {
	slog.Handler
	stdout slog.Handler
	stderr slog.Handler
}

func newCloudSlogger(opts *slog.HandlerOptions) *sLogger {
	ch := &cloudHandler{
		Handler: slog.NewJSONHandler(os.Stdout, opts),
		stdout:  slog.NewJSONHandler(os.Stdout, opts),
		stderr:  slog.NewJSONHandler(os.Stderr, opts),
	}

	return &sLogger{slog.New(ch)}
}

func (h *cloudHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level >= slog.LevelError {
		return h.stderr.Handle(ctx, r)
	}
	return h.stdout.Handle(ctx, r)
}

// defaultCloudLoggingHandlerOptions returns a set of default options for a
// Google Cloud Logging handler
func defaultCloudLoggingHandlerOptions() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.SourceKey:
				// need to validate this is a slog.Source before modifying
				source, ok := a.Value.Any().(*slog.Source)
				if ok {
					// change key to recognized google cloud logging key
					a.Key = "logging.googleapis.com/sourceLocation"
					// remove the directory from the source's filename
					source.File = filepath.Base(source.File)
				}
			case slog.LevelKey:
				a.Key = "severity"
			case slog.MessageKey:
				a.Key = "message"
			case slog.TimeKey:
				a.Value = slog.StringValue(time.Now().Format(time.RFC3339Nano))
			}
			return a
		},
	}
}
