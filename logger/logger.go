package logger

import (
	"context"
	"io"
	"time"
	"os"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var (
	AppLog *rotatelogs.RotateLogs
	MiddlewareLog *rotatelogs.RotateLogs
	Logger zerolog.Logger
)

func Initialize() {
	var (
		LogDir = viper.GetString("log.logger_dir")
		LogMaxAge = viper.GetInt("log.log_max_age")
		Debug = viper.GetBool("log.debug")
	)

	LogDir, _ = os.Getwd()
	LogDir = LogDir + "/log"

	if LogMaxAge < 1 {
		LogMaxAge = 15
	}

	AppLog, _ = rotatelogs.New(
		LogDir + "/serv_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(LogDir + "/serv_log"),
		rotatelogs.WithMaxAge(time.Duration(LogMaxAge) * 24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	MiddlewareLog, _ = rotatelogs.New(
		LogDir + "/http_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(LogDir + "/http_log"),
		rotatelogs.WithMaxAge(time.Duration(LogMaxAge) * 24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000000"
	Logger = zerolog.New(AppLog).With().Timestamp().Logger()
}

func Output(w io.Writer) zerolog.Logger {
	return Logger.Output(w)
}

func With() zerolog.Context {
	return Logger.With()
}

func Level(level zerolog.Level) zerolog.Logger {
	return Logger.Level(level)
}

func Sample(s zerolog.Sampler) zerolog.Logger {
	return Logger.Sample(s)
}

func Hook(h zerolog.Hook) zerolog.Logger {
	return Logger.Hook(h)
}

func Debug() *zerolog.Event {
	return Logger.Debug()
}

func Info() *zerolog.Event {
	return Logger.Info()
}

func Warn() *zerolog.Event {
	return Logger.Warn()
}

func Error() *zerolog.Event {
	return Logger.Error()
}

func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

func Panic() *zerolog.Event {
	return Logger.Panic()
}

func WithLevel(level zerolog.Level) *zerolog.Event {
	return Logger.WithLevel(level)
}

func Log() *zerolog.Event {
	return Logger.Log()
}

func Print(v ...interface{}) {
	Logger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
