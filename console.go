package slogger

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime"

	"github.com/fatih/color"
)

type consoleHandler struct {
	slog.Handler
	out *log.Logger
	err *log.Logger
}

const maxDepth = 16

// newConsoleSLogger returns a text-based sLogger, that sends well-formatted colorized output
// to os.Stdout or os.Stderr, based on the level of the message.
func newConsoleSLogger(level Level, addFile bool) *sLogger {
	slogOpts := &slog.HandlerOptions{
		AddSource: addFile,
		Level:     slog.Level(level),
	}

	ch := &consoleHandler{
		Handler: slog.NewTextHandler(os.Stderr, slogOpts),
		out:     log.New(os.Stdout, "", 0),
		err:     log.New(os.Stderr, "", 0),
	}

	if addFile {
		ch.out.SetFlags(log.Lshortfile)
		ch.err.SetFlags(log.Lshortfile)
	}

	return &sLogger{slog.New(ch)}
}

func (h *consoleHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String()
	if _config.coloredOutput {
		switch r.Level {
		case slog.LevelDebug:
			level = color.MagentaString(level)
		case slog.LevelInfo:
			level = color.BlueString(level)
		case slog.LevelWarn:
			level = color.YellowString(level)
		case slog.LevelError:
			level = color.RedString(level)
		}
	}

	// determine the call depth for the log message; default to 2
	calldepth := 2
	if _config.fileOutput {
		var pcs [maxDepth]uintptr

		// Capture the current call stack
		n := runtime.Callers(0, pcs[:])
		for depth := 0; depth < n; depth++ {
			if pcs[depth] == r.PC {
				calldepth = depth // setting the matching depth
			}
		}
	}

	// errors go to stderr, everything else goes to stdout
	logger := h.out
	if r.Level >= slog.LevelError {
		logger = h.err
	}
	logger.SetPrefix(fmt.Sprintf("%s ", r.Time.Format("01-02-2006 15:04:05.000")))

	return logger.Output(calldepth, fmt.Sprintf("%s %s", level, r.Message))
}
