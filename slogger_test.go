package slogger_test

import (
	"context"
	"sync"
	"testing"

	"github.com/codenaugh/slogger"
	"github.com/stretchr/testify/assert"
)

// run with 'go test -race' to test for a race; would fail without mutexes
func TestRace(t *testing.T) {
	assert.False(t, slogger.IsDebugOutputEnabled())

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.EnableDebugOutput()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.EnableCloudOutput()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Debug("debug message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Debugf("debugf %s", "message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.DebugContext(context.Background(), "debugcontext message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Log(context.Background(), slogger.LevelDebug, "debug log message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Info("info message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Infof("infof %s", "message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.InfoContext(context.Background(), "infocontext message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Log(context.Background(), slogger.LevelInfo, "info log message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Warn("warn message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Warnf("warnf %s", "message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.WarnContext(context.Background(), "warncontext message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Log(context.Background(), slogger.LevelWarn, "warn log message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Error("error message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Errorf("errorf %s", "message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.ErrorContext(context.Background(), "errorcontext message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slogger.Log(context.Background(), slogger.LevelError, "error log message")
	}()

	wg.Wait()

	assert.True(t, slogger.IsDebugOutputEnabled())
}
