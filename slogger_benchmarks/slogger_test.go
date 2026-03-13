package slogger_benchmarks

import (
	"context"
	"errors"
	"testing"

	"github.com/codenaugh/slogger"
)

func BenchmarkConsoleSLogger(b *testing.B) {
	slogger.EnableDebugOutput()
	e := errors.New("error")
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		slogger.Debug("debug test")
		slogger.Info("info", "test")
		slogger.Warn("warn", "test", 3)
		slogger.Errorf("%v", e)
		slogger.Log(ctx, slogger.LevelError, "error", "code", 6)
		slogger.InfoContext(ctx, "info for you", "<3")
	}
}

func BenchmarkConsoleSLoggerWithoutColor(b *testing.B) {
	slogger.EnableDebugOutput()
	slogger.DisableColoredOutput()
	e := errors.New("error")
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		slogger.Debug("debug test")
		slogger.Info("info", "test")
		slogger.Warn("warn", "test", 3)
		slogger.Errorf("%v", e)
		slogger.Log(ctx, slogger.LevelError, "error", "code", 6)
		slogger.InfoContext(ctx, "info for you", "<3")
	}
}

func BenchmarkConsoleSLoggerWithoutFileInfo(b *testing.B) {
	slogger.EnableDebugOutput()
	slogger.DisableFileOutput()
	e := errors.New("error")
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		slogger.Debug("debug test")
		slogger.Info("info", "test")
		slogger.Warn("warn", "test", 3)
		slogger.Errorf("%v", e)
		slogger.Log(ctx, slogger.LevelError, "error", "code", 6)
		slogger.InfoContext(ctx, "info for you", "<3")
	}
}

func BenchmarkConsoleSLoggerWithoutColorAndFileInfo(b *testing.B) {
	slogger.EnableDebugOutput()
	slogger.DisableColoredOutput()
	slogger.DisableFileOutput()
	e := errors.New("error")
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		slogger.Debug("debug test")
		slogger.Info("info", "test")
		slogger.Warn("warn", "test", 3)
		slogger.Errorf("%v", e)
		slogger.Log(ctx, slogger.LevelError, "error", "code", 6)
		slogger.InfoContext(ctx, "info for you", "<3")
	}
}

func BenchmarkCloudSLogger(b *testing.B) {
	slogger.EnableDebugOutput()
	slogger.EnableCloudOutput()
	e := errors.New("error")
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		slogger.Debug("debug test")
		slogger.Info("info", "test")
		slogger.Warn("warn", "test", 3)
		slogger.Errorf("%v", e)
		slogger.Log(ctx, slogger.LevelError, "error", "code", 6)
		slogger.InfoContext(ctx, "info for you", "<3")
	}
}
