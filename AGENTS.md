## Cursor Cloud specific instructions

### Repository structure

This is a **learning monorepo** of independent Go example projects. There is **no root `go.mod`** — each subdirectory is a standalone program. Seven subdirectories have their own `go.mod`; the rest are legacy GOPATH-style single-file programs.

### Go modules vs legacy directories

| Type | Directories | How to build/run |
|---|---|---|
| **go.mod projects** | `http-server`, `mongo-microservice`, `rest-kafka-mongo-microservice`, `htmx`, `ai/go-groq`, `connect-to-db-postgres`, `docker` | `cd <dir> && go build ./...` or `go run main.go` |
| **Legacy (no go.mod)** | All others (`helloworld`, `channels`, `closures`, `testing`, `benchmark`, etc.) | `cd <dir> && GO111MODULE=off go run main.go` or `GO111MODULE=off go build .` |

### Running tests

- Unit tests: `cd testing && GO111MODULE=off go test -v .`
- Benchmarks: `cd benchmark && GO111MODULE=off go test -bench=. .`

### Linting

- `go vet ./...` inside any go.mod directory
- `staticcheck ./...` (installed in `~/go/bin`) inside any go.mod directory
- For legacy directories: `cd <dir> && GO111MODULE=off go vet .`

### Running the HTTP server (primary demo app)

```
cd http-server && go run main.go
```
Listens on `:8000`. Verify with `curl http://localhost:8000/` (returns "Hello World").

### External-service-dependent projects (optional)

These projects require external services and are **not** needed for core development:

- `mongo-microservice`, `rest-kafka-mongo-microservice` — need MongoDB (`:27017`); docker-compose at `mongo-microservice/mongodb/docker-compose.yml`
- `kafka/producer`, `kafka/consumer`, `rest-kafka-mongo-microservice` — need Apache Kafka (`:9092`)
- `connect-to-db-postgres` — needs PostgreSQL (`:5432`); docker-compose at `.devcontainer/docker-compose.yml`
- `docker/*` — needs Docker daemon
- `ai/go-groq`, `htmx` — need `GROQ_API_KEY` env var

### Gotchas

- `pointers/basic/` has multiple files with `main` redeclared — this is a pre-existing repo issue. Build individual files with `GO111MODULE=off go run main.go`.
- `~/go/bin` must be on `PATH` to use `staticcheck`. It is added via `~/.bashrc`.
