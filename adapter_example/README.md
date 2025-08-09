# Adapter Pattern (Python) Example

This example shows how an `Adapter` makes an incompatible `Adaptee` usable via the `Target` interface expected by the client.

## Files
- `target.py`: Defines the `Target` interface.
- `adaptee.py`: Provides an incompatible API via `specific_request()`.
- `adapter.py`: Implements `Target` and translates calls to the `Adaptee`.
- `client.py`: Demonstrates the client working with the `Adapter`.

## Run
```bash
python3 client.py
```