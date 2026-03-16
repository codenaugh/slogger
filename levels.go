package slogger

import (
	"fmt"
	"log/slog"
)

type Level slog.Level

const (
	LevelDebug = Level(slog.LevelDebug)
	LevelInfo  = Level(slog.LevelInfo)
	LevelWarn  = Level(slog.LevelWarn)
	LevelError = Level(slog.LevelError)
)

// String returns a name for the level.
// If the level has a name, then that name
// in uppercase is returned.
// If the level is between named values, then
// an integer is appended to the uppercased name.
// Examples:
//
//	LevelWarn.String() => "WARN"
//	(LevelInfo+2).String() => "INFO+2"
func (l Level) String() string {
	str := func(base string, val Level) string {
		if val == 0 {
			return base
		}
		return fmt.Sprintf("%s%+d", base, val)
	}

	switch {
	case l < LevelInfo:
		return str("DEBUG", l-LevelDebug)
	case l < LevelWarn:
		return str("INFO", l-LevelInfo)
	case l < LevelError:
		return str("WARN", l-LevelWarn)
	default:
		return str("ERROR", l-LevelError)
	}
}

// MarshalText implements [encoding.TextMarshaler].
func (l Level) MarshalText() ([]byte, error) {
	return slog.Level(l).MarshalText()
}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (l *Level) UnmarshalText(data []byte) error {
	var sl slog.Level
	if err := sl.UnmarshalText(data); err != nil {
		return err
	}
	*l = Level(sl)
	return nil
}
