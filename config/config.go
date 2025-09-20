package config

import (
	"os"

	"github.com/tvoybuket/tblib/tbconfig"
	"github.com/tvoybuket/tblib/tblogger"
)

func MustLoad() *Settings {
	settings := &Settings{}
	if err := tbconfig.LoadConfig(settings); err != nil {
		tblogger.Fatal("failed to load logger settings: %v", err)
	}

	return settings
}

func MustLoadLoggerSettings() *tblogger.Config {
	settings := &LoggerSettings{}
	if err := tbconfig.LoadConfig(settings); err != nil {
		tblogger.Fatal("failed to load logger settings: %v", err)
	}

	level := parseLogLevel(settings.Level)
	format := parseLogFormat(settings.Format)

	cfg := &tblogger.Config{
		Level:       level,
		Format:      format,
		ServiceName: settings.ServiceName,
		Output:      os.Stdout,
	}

	return cfg
}

func parseLogLevel(levelStr string) tblogger.LogLevel {
	switch levelStr {
	case "debug":
		return tblogger.LevelDebug
	case "info":
		return tblogger.LevelInfo
	case "warn":
		return tblogger.LevelWarn
	case "error":
		return tblogger.LevelError
	default:
		return tblogger.LevelInfo
	}
}

func parseLogFormat(formatStr string) tblogger.OutputFormat {
	switch formatStr {
	case "text":
		return tblogger.FormatText
	case "json":
		return tblogger.FormatJSON
	default:
		return tblogger.FormatJSON
	}
}
