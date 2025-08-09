package main

// StdLoggerAdapter adapts StdLoggerAdaptee to the Logger interface.
type StdLoggerAdapter struct {
	std *StdLoggerAdaptee
}

func NewStdLoggerAdapter(std *StdLoggerAdaptee) *StdLoggerAdapter {
	return &StdLoggerAdapter{std: std}
}

func (a *StdLoggerAdapter) Debug(msg string, fields map[string]any) {
	a.std.PrintLine("DEBUG", msg, fields)
}

func (a *StdLoggerAdapter) Info(msg string, fields map[string]any) {
	a.std.PrintLine("INFO", msg, fields)
}

func (a *StdLoggerAdapter) Warn(msg string, fields map[string]any) {
	a.std.PrintLine("WARN", msg, fields)
}

func (a *StdLoggerAdapter) Error(msg string, fields map[string]any) {
	a.std.PrintLine("ERROR", msg, fields)
}