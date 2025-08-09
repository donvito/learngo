# Adapter Pattern (Go) — Realistic Logger Example

Adapt the standard library `log.Logger` (adaptee) to an application-defined structured `Logger` interface (target). This mirrors real-world cases where your app expects an interface but you need to plug in a third-party/stdlib implementation with a different API.

## Files
- `logger.go`: The `Logger` interface used by application code.
- `stdlog_adaptee.go`: Wraps `log.Logger` and exposes an incompatible method `PrintLine`.
- `stdlog_adapter.go`: Implements `Logger` by delegating to the adaptee and handling level/fields.
- `main.go`: Demonstrates usage.

## Run
```bash
go run /workspace/adapter_example_go
```

## Output (example)
```
2025/01/01 12:00:00 stdlog_adaptee.go:22: [INFO] starting server | addr=:8080 mode=dev 
...
```