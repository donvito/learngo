# Adapter Pattern (Go) — Realistic Logger Example

Adapt the standard library `log.Logger` (adaptee) to an application-defined structured `Logger` interface (target). This mirrors real-world cases where your app expects an interface but you need to plug in a third-party/stdlib implementation with a different API.

## Intent
- **Problem**: Application code wants structured, leveled logging via a simple `Logger` interface, but the available logger (stdlib `log.Logger`) exposes a different API (`Print*`/`Printf` without levels or fields).
- **Solution**: Introduce an Adapter that makes `log.Logger` look like your `Logger` interface.

## Structure (roles → files)
- **Target** → `Logger` interface: `logger.go`
- **Adaptee** → wrapper over `log.Logger` exposing an incompatible method: `stdlog_adaptee.go`
- **Adapter** → implements `Logger` by delegating to the Adaptee: `stdlog_adapter.go`
- **Client** → application code that depends only on `Logger`: `main.go`

## How the call flows
1. `main.go` creates a `*log.Logger` and wraps it in `StdLoggerAdaptee`.
2. `StdLoggerAdapter` is created around the adaptee and returned as a `Logger`.
3. Client code calls `logger.Info("msg", fields)` on the `Logger` interface.
4. Adapter translates the call into `adaptee.PrintLine(level, msg, fields)`.
5. Adaptee formats the message/fields and calls the underlying `log.Logger.Printf`.

## Code tour

### Target: what the app wants
```go
// logger.go
package main

// Logger is the interface our application expects: structured logging with levels.
type Logger interface {
    Debug(msg string, fields map[string]any)
    Info(msg string, fields map[string]any)
    Warn(msg string, fields map[string]any)
    Error(msg string, fields map[string]any)
}
```

### Adaptee: what we actually have
```go
// stdlog_adaptee.go
package main

import (
    "fmt"
    "log"
)

// StdLoggerAdaptee represents the concrete type we cannot change (std log.Logger)
type StdLoggerAdaptee struct { logger *log.Logger }

func NewStdLoggerAdaptee(l *log.Logger) *StdLoggerAdaptee { return &StdLoggerAdaptee{logger: l} }

// PrintLine is the incompatible API we need to adapt to our Logger interface.
func (s *StdLoggerAdaptee) PrintLine(level string, msg string, fields map[string]any) {
    formatted := msg
    if len(fields) > 0 {
        formatted += " | "
        for k, v := range fields {
            formatted += fmt.Sprintf("%s=%v ", k, v)
        }
    }
    s.logger.Printf("[%s] %s", level, formatted)
}
```

### Adapter: the bridge between the two
```go
// stdlog_adapter.go
package main

// StdLoggerAdapter adapts StdLoggerAdaptee to the Logger interface.
type StdLoggerAdapter struct { std *StdLoggerAdaptee }

func NewStdLoggerAdapter(std *StdLoggerAdaptee) *StdLoggerAdapter { return &StdLoggerAdapter{std: std} }

func (a *StdLoggerAdapter) Debug(msg string, fields map[string]any) { a.std.PrintLine("DEBUG", msg, fields) }
func (a *StdLoggerAdapter) Info(msg string, fields map[string]any)  { a.std.PrintLine("INFO", msg, fields) }
func (a *StdLoggerAdapter) Warn(msg string, fields map[string]any)  { a.std.PrintLine("WARN", msg, fields) }
func (a *StdLoggerAdapter) Error(msg string, fields map[string]any) { a.std.PrintLine("ERROR", msg, fields) }
```

### Client: uses only the Target interface
```go
// main.go
package main

import (
    "log"
    "os"
)

func main() {
    std := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)  // concrete logger we cannot change
    adaptee := NewStdLoggerAdaptee(std)                           // wrap as Adaptee
    logger := NewStdLoggerAdapter(adaptee)                        // adapt to Target interface

    // Application code depends only on Logger
    logger.Info("starting server", map[string]any{"addr": ":8080", "mode": "dev"})
    logger.Debug("cache primed", map[string]any{"entries": 128})
    logger.Warn("rate limit approaching", map[string]any{"remaining": 3})
    logger.Error("database connection failed", map[string]any{"retry_in": "2s", "attempt": 1})
}
```

## Why Adapter (and not something else)?
- **Adapter vs. Facade**: Facade provides a simplified API to a subsystem; Adapter translates between two existing interfaces.
- **Adapter vs. Wrapper**: A wrapper is a generic term; an Adapter specifically matches one interface to another.
- **Adapter vs. Bridge**: Bridge decouples abstraction from implementation to vary independently; Adapter is for compatibility.

## Design notes
- **Decoupling**: The app never imports `log` directly; it relies on `Logger`. This makes it trivial to swap in `zap`, `zerolog`, or a no-op logger via another Adapter.
- **Formatting**: The example uses a simple `key=value` format to avoid extra deps. Swap the formatting inside `StdLoggerAdaptee.PrintLine` as needed.
- **Thread-safety**: `log.Logger` is safe for concurrent use. The adapter maintains no mutable shared state, so it’s goroutine-safe under the same guarantees.
- **Performance**: Field formatting is O(n) over the map. For high-throughput paths, consider a `strings.Builder` or pre-allocated buffers and avoid map iteration order assumptions.
- **Error handling**: Logging shouldn’t fail the app; if formatting fails, prefer a best-effort representation.

## Extending the example
- **Map log levels**: If the adaptee has different levels (e.g., `zapcore.Level`), implement a mapping in the adapter.
- **Contextual fields**: Add `With(fields)` to `Logger` and return a child adapter that merges base fields with per-call fields.
- **Sampling**: Implement sampling in the adapter before delegating to the adaptee to reduce log volume.
- **Multiple backends**: Create a composite adapter that fans out to multiple adaptees (e.g., stdout + file + remote).

## When to use this pattern
- You can’t change the third-party/stdlib type, but you control your app’s interface.
- You need to swap implementations without touching business logic.
- You must support multiple backends behind a stable interface.

## Run
```bash
go run /workspace/adapter_example_go
```

## Example output
```
2025/08/09 19:25:17 stdlog_adaptee.go:27: [INFO] starting server | addr=:8080 mode=dev 
2025/08/09 19:25:17 stdlog_adaptee.go:27: [DEBUG] cache primed | entries=128 
2025/08/09 19:25:17 stdlog_adaptee.go:27: [WARN] rate limit approaching | remaining=3 
2025/08/09 19:25:17 stdlog_adaptee.go:27: [ERROR] database connection failed | retry_in=2s attempt=1 
```