package slogger_test

import (
	"context"
	"errors"
	"flag"

	"github.com/codenaugh/slogger"
)

func expensiveFunc() string {
	// expensive operation
	return "expensive data"
}

func ExampleEnableCloudOutput() {
	var local, debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")
	flag.BoolVar(&local, "local", false, "Running locally; Pretty Text instead of JSON log messages")
	flag.Parse()

	if !local {
		slogger.EnableCloudOutput()
	}
	if debug {
		slogger.EnableDebugOutput()
	}

	slogger.Info("this will output a JSON-formatted log message")
}

func ExampleEnableDebugOutput() {
	var local, debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")
	flag.BoolVar(&local, "local", false, "Running locally; Pretty Text instead of JSON log messages")
	flag.Parse()

	if !local {
		slogger.EnableCloudOutput()
	}
	if debug {
		slogger.EnableDebugOutput()
	}

	slogger.Debug("this will output if debug is enabled, or be suppressed if not")
}

func ExampleIsDebugOutputEnabled() {
	if slogger.IsDebugOutputEnabled() {
		slogger.Debugf("let's output some expensive data: %v", expensiveFunc())
	}
}

func ExampleDisableColoredOutput() {
	slogger.DisableColoredOutput()

	slogger.Info("this will output without any added color")
}

func ExampleDisableFileOutput() {
	slogger.DisableFileOutput()

	slogger.Info("this will output without any file info")

	// Output
	// 11-20-2024 16:12:02.556 INFO this will output without any file info
}

func ExampleDebug() {
	slogger.EnableDebugOutput()

	argsString := "args"

	slogger.Debug("debug message")
	slogger.Debug("debug message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: DEBUG debug message
	// 11-20-2024 16:12:02.556 main.go:17: DEBUG debug message with extra args 10
}

func ExampleInfo() {
	argsString := "args"

	slogger.Info("info message")
	slogger.Info("info message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: INFO info message
	// 11-20-2024 16:12:02.556 main.go:17: INFO message with extra args 10
}

func ExampleWarn() {
	argsString := "args"

	slogger.Warn("warning message")
	slogger.Warn("warning message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: WARN warning message
	// 11-20-2024 16:12:02.556 main.go:17: WARN warning message with extra args 10
}

func ExampleError() {
	argsString := "args"
	err := errors.New("error message")

	slogger.Error(err)
	slogger.Error("error message")
	slogger.Error("error message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: ERROR error message
	// 11-20-2024 16:12:02.556 main.go:17: ERROR error message
	// 11-20-2024 16:12:02.556 main.go:17: ERROR error message with extra args 10
}

func ExampleFatal() {
	err := errors.New("fatal error")

	slogger.Fatal(err)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: ERROR fatal error
	// exit status 1

	slogger.Fatal("fatal message:", err)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: ERROR fatal message: fatal error
	// exit status 1
}

func ExampleDebugf() {
	slogger.EnableDebugOutput()

	argsString := "args"

	slogger.Debugf("debug message")
	slogger.Debugf("debug message %s %s %v: %d", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: DEBUG debug message
	// 11-20-2024 16:12:02.556 main.go:17: DEBUG debug message with extra args: 10
}

func ExampleInfof() {
	argsString := "args"

	slogger.Infof("info message")
	slogger.Infof("info message %s %s %v: %d", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: INFO info message
	// 11-20-2024 16:12:02.556 main.go:17: INFO info message with extra args: 10
}

func ExampleWarnf() {
	argsString := "args"

	slogger.Warnf("warning message")
	slogger.Warnf("warning message %s %s %v: %d", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: WARN warning message
	// 11-20-2024 16:12:02.556 main.go:17: WARN warning message with extra args: 10
}

func ExampleErrorf() {
	argsString := "args"

	slogger.Errorf("error message")
	slogger.Errorf("error message %s %s %v: %d", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: ERROR error message
	// 11-20-2024 16:12:02.556 main.go:17: ERROR error message with extra args: 10
}

func ExampleFatalf() {
	err := errors.New("fatal error")
	slogger.Fatalf("slogger fatal message: %v", err)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: ERROR slogger fatal message: fatal error
	// exit status 1
}

func ExampleDebugContext() {
	slogger.EnableDebugOutput()

	ctx := context.Background()
	argsString := "args"

	slogger.DebugContext(ctx, "debug message")
	slogger.DebugContext(ctx, "debug message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: DEBUG debug message
	// 11-20-2024 16:12:02.556 main.go:17: DEBUG debug message with extra args 10
}

func ExampleInfoContext() {
	ctx := context.Background()
	argsString := "args"

	slogger.InfoContext(ctx, "info message")
	slogger.InfoContext(ctx, "info message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: INFO info message
	// 11-20-2024 16:12:02.556 main.go:17: INFO info message with extra args 10
}

func ExampleWarnContext() {
	ctx := context.Background()
	argsString := "args"

	slogger.WarnContext(ctx, "warning message")
	slogger.WarnContext(ctx, "warning message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: WARN warning message
	// 11-20-2024 16:12:02.556 main.go:17: WARN warning message with extra args 10
}

func ExampleErrorContext() {
	ctx := context.Background()
	argsString := "args"

	slogger.ErrorContext(ctx, "error message")
	slogger.ErrorContext(ctx, "error message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: ERROR error message
	// 11-20-2024 16:12:02.556 main.go:17: ERROR error message with extra args 10
}

func ExampleLog() {
	argsString := "args"

	slogger.Log(context.Background(), slogger.LevelInfo, "slogger log message")
	slogger.Log(context.Background(), slogger.LevelInfo, "slogger log message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: INFO slogger log message
	// 11-20-2024 16:12:02.556 main.go:17: INFO slogger log message with extra args 10
}

func ExampleOutput() {
	argsString := "args"

	slogger.Output(2, slogger.LevelInfo, "slogger output message")
	slogger.Output(2, slogger.LevelInfo, "slogger output message", "with", "extra", argsString, 10)

	// Output
	// 11-20-2024 16:12:02.556 main.go:17: INFO slogger output message
	// 11-20-2024 16:12:02.556 main.go:17: INFO slogger output message with extra args 10
}
