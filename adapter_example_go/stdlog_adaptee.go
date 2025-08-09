package main

import (
	"fmt"
	"log"
)

// StdLoggerAdaptee represents the concrete type we cannot change (std log.Logger)
type StdLoggerAdaptee struct {
	logger *log.Logger
}

func NewStdLoggerAdaptee(l *log.Logger) *StdLoggerAdaptee {
	return &StdLoggerAdaptee{logger: l}
}

// PrintLine is the incompatible API we need to adapt to our Logger interface.
func (s *StdLoggerAdaptee) PrintLine(level string, msg string, fields map[string]any) {
	// naive key=value formatting to keep example dependency-free
	formatted := msg
	if len(fields) > 0 {
		formatted += " | "
		for k, v := range fields {
			formatted += fmt.Sprintf("%s=%v ", k, v)
		}
	}
	s.logger.Printf("[%s] %s", level, formatted)
}