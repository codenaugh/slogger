/*
Package slogger provides both well-formatted human-readable text output and Google Cloud optimized JSON output.

It is a lightweight custom logging package built on top of the Go standard library [slog] package, that
provides several options for severity-based logging.

The default is a text-based console logger that outputs INFO and WARN severity level messages to os.Stdout,
and ERROR level severity messages to os.Stderr. By default, the message is prefixed with the timestamp in the
format of MM-DD-YYYY hh:mm:ss.mmm, then has the originating filename and line number output next to the colorized
severity level, followed by the message.

# Configuration Options

		`EnableCloudOutput`:    Enables Google Cloud optimized JSON output instead of "pretty" message output.
		`EnableDebugOutput`:    Enables DEBUG level output to os.Stdout instead of it being suppressed.
	                         A user may use IsDebugOutputEnabled() to check if debug output is enabled or not.
		`DisableColoredOutput`: The severity levels will no longer be colored in the console output.
		`DisableFileOutput`:    The originating filename and line number will no longer be output in the console output.

# Performance Considerations

The arguments to a log call are always evaluated, even if the log event is discarded. If possible,
defer computation so that it happens only if the value is actually logged. For example, consider the call

	slogger.Debug("starting request", "url", r.URL.String())  // may compute String unnecessarily

The URL.String method will be called even if the logger discards DEBUG level events.

To avoid expensive debug calls using slogger, you may guard the log call with a condition:

	if slogger.IsDebugOutputEnabled() {
	    slog.Debug("starting request", "url", r.URL.String())
	}

You can also use the [LogValuer] interface to avoid unnecessary work in disabled log
calls. Say you need to log some expensive value:

	logger.Debug("frobbing", "value", computeExpensiveValue(arg))

Even if this line is disabled, computeExpensiveValue will be called.
To avoid that, define a type implementing LogValuer:

	type expensive struct { arg int }

	func (e expensive) LogValue() slog.Value {
	    return slog.AnyValue(computeExpensiveValue(e.arg))
	}

Then use a value of that type in log calls:

	logger.Debug("frobbing", "value", expensive{arg})

Now computeExpensiveValue will only be called when the line is enabled.

[slog]: https://pkg.go.dev/log/slog
*/
package slogger
