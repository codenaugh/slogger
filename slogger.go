package slogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type sLogger struct {
	*slog.Logger
}

type config struct {
	// cloudOutput is true when the output is JSON and false when it is text to the console.
	cloudOutput bool
	// debugOutput is true when the level is set to Debug, and false when set to Info, which is the default level.
	debugOutput bool
	// coloredOutput is optional for console sLoggers, but always disabled for cloud sLoggers.
	// The default, true, means that each message is prefixed with a colorized severity level string.
	coloredOutput bool
	// fileOutput is optional for console sLoggers, but always enabled for cloud sLoggers.
	// The default, true, means that each message includes the originating file name and line number.
	fileOutput bool
}

var (
	_mu      sync.RWMutex
	_config  = defaultConfig()
	_slogger = newConsoleSLogger(LevelInfo, true)
)

func defaultConfig() config {
	return config{
		cloudOutput:   false,
		debugOutput:   false,
		coloredOutput: true,
		fileOutput:    true,
	}
}

// reconfigSLogger replaces the global sLogger based on the global config.
// It is called after a change to the global config.
func reconfigSLogger() {
	_mu.Lock()
	defer _mu.Unlock()

	level := LevelInfo

	if _config.debugOutput {
		level = LevelDebug
	}

	if _config.cloudOutput {
		hOpts := defaultCloudLoggingHandlerOptions()
		hOpts.Level = slog.Level(level)
		_slogger = newCloudSlogger(hOpts)
		return
	}

	_slogger = newConsoleSLogger(level, _config.fileOutput)
}

// EnableDebugOutput enables debug output for the sLogger
// without changing other previously configured options.
func EnableDebugOutput() {
	_mu.Lock()
	_config.debugOutput = true
	_mu.Unlock()

	reconfigSLogger()
}

// EnableCloudOutput enables cloud output for the sLogger
// without changing other previously configured options.
func EnableCloudOutput() {
	_mu.Lock()
	_config.cloudOutput = true
	_mu.Unlock()

	reconfigSLogger()
}

// IsDebugOutputEnabled returns true if debug output is enabled for the
// sLogger.
//
// This may be useful for avoiding expensive debug calls (see doc.go)
func IsDebugOutputEnabled() bool {
	_mu.RLock()
	defer _mu.RUnlock()

	return _config.debugOutput
}

// DisableColoredOutput disables the coloring of the Level text.
//
// This is a no-op if using the cloud logging handler.
func DisableColoredOutput() {
	_mu.Lock()
	if _config.cloudOutput {
		_mu.Unlock()
		return
	}
	_config.coloredOutput = false
	_mu.Unlock()

	reconfigSLogger()
}

// DisableFileOutput disables the file name and line number output.
//
// This is a no-op if using the cloud logging handler.
func DisableFileOutput() {
	_mu.Lock()
	if _config.cloudOutput {
		_mu.Unlock()
		return
	}
	_config.fileOutput = false
	_mu.Unlock()

	reconfigSLogger()
}

// Debug logs a message at the debug level, as long as debug mode is enabled,
// with args being of any type and treated as space-separated strings
func Debug(args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	if _config.debugOutput {
		var pcs [1]uintptr
		runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

		var msg string
		for _, arg := range args {
			msg += fmt.Sprintf(" %v", arg)
		}
		msg = strings.TrimSpace(msg)

		r := slog.NewRecord(time.Now(), slog.LevelDebug, msg, pcs[0])
		_ = _slogger.Handler().Handle(context.Background(), r)
	}
}

// Info logs a message at the info level with args being of any type
// and treated as space-separated strings
func Info(args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, pcs[0])
	_ = _slogger.Handler().Handle(context.Background(), r)
}

// Warn logs a message at the warn level with args being of any type
// and treated as space-separated strings
func Warn(args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.LevelWarn, msg, pcs[0])
	_ = _slogger.Handler().Handle(context.Background(), r)
}

// Error logs a message at the error level with args being of any type
// and treated as space-separated strings
func Error(args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0])
	_ = _slogger.Handler().Handle(context.Background(), r)
}

// Fatal logs a message at the error level with args being of any type
// and treated as space-separated strings, then exits
func Fatal(args ...any) {
	_mu.Lock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0])
	_ = _slogger.Handler().Handle(context.Background(), r)

	_mu.Unlock()

	os.Exit(1)
}

// Debugf logs a message at the debug level using the given format string and args,
// if debug mode is enabled.
func Debugf(format string, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	if _config.debugOutput {
		var pcs [1]uintptr
		runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

		r := slog.NewRecord(time.Now(), slog.LevelDebug, fmt.Sprintf(format, args...), pcs[0])
		_ = _slogger.Handler().Handle(context.Background(), r)
	}
}

// Infof logs a message at the info level using the given format string and args.
func Infof(format string, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(format, args...), pcs[0])
	_ = _slogger.Handler().Handle(context.Background(), r)
}

// Warnf logs a message at the warn level using the given format string and args.
func Warnf(format string, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	r := slog.NewRecord(time.Now(), slog.LevelWarn, fmt.Sprintf(format, args...), pcs[0])
	_ = _slogger.Handler().Handle(context.Background(), r)
}

// Errorf logs a message at the error level using the given format string and args.
func Errorf(format string, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(format, args...), pcs[0])
	_ = _slogger.Handler().Handle(context.Background(), r)
}

// Fatalf logs a message at the error level using the given format string and args, then exits.
func Fatalf(format string, args ...any) {
	_mu.Lock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(format, args...), pcs[0])
	_ = _slogger.Handler().Handle(context.Background(), r)

	_mu.Unlock()

	os.Exit(1)
}

// DebugContext logs a message at the debug level, if enabled, with args being of any type
func DebugContext(ctx context.Context, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	if _config.debugOutput {
		var pcs [1]uintptr
		runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

		var msg string
		for _, arg := range args {
			msg += fmt.Sprintf(" %v", arg)
		}
		msg = strings.TrimSpace(msg)

		r := slog.NewRecord(time.Now(), slog.LevelDebug, msg, pcs[0])
		_ = _slogger.Handler().Handle(ctx, r)
	}
}

// InfoContext logs a message at the info level with args being of any type
func InfoContext(ctx context.Context, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, pcs[0])
	_ = _slogger.Handler().Handle(ctx, r)
}

// WarnContext logs a message at the warn level with args being of any type
func WarnContext(ctx context.Context, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.LevelWarn, msg, pcs[0])
	_ = _slogger.Handler().Handle(ctx, r)
}

// ErrorContext logs a message at the error level with args being of any type
func ErrorContext(ctx context.Context, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0])
	_ = _slogger.Handler().Handle(ctx, r)
}

// Log outputs a message at the given level, if enabled, with args being of any type
func Log(ctx context.Context, level Level, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	if !_slogger.Enabled(ctx, slog.Level(level)) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [runtime.Callers, this function]

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.Level(level), msg, pcs[0])
	_ = _slogger.Handler().Handle(ctx, r)
}

// Output logs a message for the given call frame at the given level, if enabled, with args being of any type.
//
// This allows for logging from a specific call frame, rather than the default of the calling function,
// like if you want the log to originate from the parent function of the calling function.
func Output(calldepth int, level Level, args ...any) {
	_mu.Lock()
	defer _mu.Unlock()

	ctx := context.Background()

	if !_slogger.Enabled(ctx, slog.Level(level)) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(calldepth+1, pcs[:]) // skip +1 frames

	var msg string
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	msg = strings.TrimSpace(msg)

	r := slog.NewRecord(time.Now(), slog.Level(level), msg, pcs[0])
	_ = _slogger.Handler().Handle(ctx, r)
}

// MapToGCPSeverity maps slog levels to GCP severity levels
func MapToGCPSeverity(level slog.Level) string {
	switch {
	case level >= slog.LevelError:
		return "ERROR"
	case level >= slog.LevelWarn:
		return "WARNING"
	case level >= slog.LevelInfo:
		return "INFO"
	default:
		return "DEBUG"
	}
}
