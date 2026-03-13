# slogger

slogger provides both well-formatted human-readable text output and Google Cloud optimized JSON output.

It is a lightweight custom logging package built on top of the Go standard library [`slog`](https://pkg.go.dev/log/slog) package, that
provides several options for severity-based logging.

If you're curious about the performance, check the benchmarks located in each benchmark directory.

## Features

#### Clean Console Output
The default is well-formatted text-based INFO level and higher output sent to stdout and stderr that is easy for humans to read and parse through.

![image](https://github.com/user-attachments/assets/21a97c92-21e6-4470-88f3-42410dbb9271)


#### Clean Google Cloud Logging Output
With `EnableCloudOutput()`, logs are output in JSON format with some fields re-keyed to ensure proper recognition and formatting for [Google Cloud Structured Logging](https://cloud.google.com/logging/docs/structured-logging). This prevents bloated and duplicated info in log entries.

#### Debug Output
With `EnableDebugOutput()`, debug-level messages are output, to console or Google Cloud, instead of being suppressed by default.

#### Colored Severity Level Output
In console mode, the severity level of each log message is color-coded for easy recognition, and so things like WARN and ERROR messages stand out. This can be toggled off with `DisableColoredOutput()`

#### Easy Troubleshooting
In console mode, each message with have the originating file and line number, making it easy to track down where the log message was generated. This can be toggled off with `DisableFileOutput()` if a user doesn't need it or prefers a cleaner look.

## Install

```bash
go get github.com/codenaugh/slogger
```
