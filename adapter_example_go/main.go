package main

import (
	"log"
	"os"
)

func main() {
	// Adaptee: stdlib logger with default settings
	std := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	adaptee := NewStdLoggerAdaptee(std)

	// Adapter to the application's expected Logger interface
	logger := NewStdLoggerAdapter(adaptee)

	// Application code uses Logger interface only
	logger.Info("starting server", map[string]any{"addr": ":8080", "mode": "dev"})
	logger.Debug("cache primed", map[string]any{"entries": 128})
	logger.Warn("rate limit approaching", map[string]any{"remaining": 3})
	logger.Error("database connection failed", map[string]any{"retry_in": "2s", "attempt": 1})
}