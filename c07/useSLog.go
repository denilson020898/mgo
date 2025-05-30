package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	slog.Error("This an ERROR message")
	slog.Debug("This an DEBUG message")
	slog.Info("This an INFO message")
	slog.Warn("This an WARNING message")

	logLevel := &slog.LevelVar{}
	fmt.Println("Log level:", logLevel)
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)
	logLevel.Set(slog.LevelDebug)
	logger.Debug("This is a DEBUG message")

	logJSON := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logJSON.Error("ERROR message in JSON")
}
