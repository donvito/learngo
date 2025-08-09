# Adapter Pattern (Go) Example

This example shows how an `Adapter` makes an incompatible `Adaptee` usable via the `Target` interface expected by the client.

## Files
- `target.go`: Defines the `Target` interface.
- `adaptee.go`: Provides an incompatible API via `SpecificRequest()`.
- `adapter.go`: Implements `Target` and translates calls to the `Adaptee`.
- `main.go`: Demonstrates the client working with the `Adapter`.

## Run
```bash
# From repo root
go run /workspace/adapter_example_go
```