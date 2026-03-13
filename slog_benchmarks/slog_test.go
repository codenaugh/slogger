package slog_benchmarks_test

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

func BenchmarkSLog(b *testing.B) {
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		slog.Debug("debug test")
		slog.Info("info", "test", "value")
		slog.Warn("warn", "test", 3)
		slog.Error("error")
		slog.Log(ctx, slog.LevelError, "error", "code", 6)
		slog.InfoContext(ctx, "info from slog", "<", 3)
	}
}

func BenchmarkTextSLog(b *testing.B) {
	ctx := context.Background()
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	for i := 0; i < b.N; i++ {
		l.Debug("debug test")
		l.Info("info", "test", "value")
		l.Warn("warn", "test", 3)
		l.Error("error")
		l.Log(ctx, slog.LevelError, "error", "code", 6)
		l.InfoContext(ctx, "info from slog", "<", 3)
	}
}

func BenchmarkTextSLog2(b *testing.B) {
	ctx := context.Background()
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(l)

	for i := 0; i < b.N; i++ {
		slog.Debug("debug test")
		slog.Info("info", "test", "value")
		slog.Warn("warn", "test", 3)
		slog.Error("error")
		slog.Log(ctx, slog.LevelError, "error", "code", 6)
		slog.InfoContext(ctx, "info from slog", "<", 3)
	}
}

func BenchmarkJSONSLog(b *testing.B) {
	ctx := context.Background()
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	for i := 0; i < b.N; i++ {
		l.Debug("debug test")
		l.Info("info", "test", "value")
		l.Warn("warn", "test", 3)
		l.Error("error")
		l.Log(ctx, slog.LevelError, "error", "code", 6)
		l.InfoContext(ctx, "info from slog", "<", 3)
	}
}

func BenchmarkJSONSLog2(b *testing.B) {
	ctx := context.Background()
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(l)

	for i := 0; i < b.N; i++ {
		slog.Debug("debug test")
		slog.Info("info", "test", "value")
		slog.Warn("warn", "test", 3)
		slog.Error("error")
		slog.Log(ctx, slog.LevelError, "error", "code", 6)
		slog.InfoContext(ctx, "info from slog", "<", 3)
	}
}
