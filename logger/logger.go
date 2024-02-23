package logger

import (
	"io"
	"log/slog"
	"os"
	"sync"

	"gocommon/config"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	lo   *slog.Logger
	once sync.Once
)

func Default() *slog.Logger {
	return lo
}

func Init() {
	logconfig := config.GetLog()
	once.Do(func() {
		lo = slog.New(slog.NewTextHandler(
			NewWriter(logconfig.Path), &slog.HandlerOptions{
				AddSource: false,
				Level:     slog.LevelInfo,
			},
		))
	})
}

func NewWriter(path string) io.WriteCloser {
	// return os.Stdout
	if config.Default().GetBool("debug") {
		return os.Stdout
	}
	logconfig := config.GetLog().Roate
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    logconfig.MaxSize, // megabytes
		MaxBackups: logconfig.MaxBackup,
		MaxAge:     logconfig.MaxAge, // dayså
		Compress:   logconfig.Compress,
		LocalTime:  logconfig.Compress,
	}
}
