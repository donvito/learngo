# Code Review and Test Results

Date: 2026-05-12  
Branch: `cursor/review-results-b76c`

## Scope

Reviewed the Go examples in this repository and ran the available Go test commands across:

- The repository root
- Standalone non-module test directories
- Nested Go modules:
  - `http-server`
  - `htmx`
  - `ai/go-groq`
  - `connect-to-db-postgres`

The repository does not have a root `go.mod`; several example directories are standalone Go programs, while a few subdirectories define their own modules.

## Test Results

| Command | Directory | Result | Notes |
| --- | --- | --- | --- |
| `go version && go test ./...` | `/workspace` | Failed | Go version was `go1.22.2 linux/amd64`. Test command failed because the repository root is not a Go module: `pattern ./...: directory prefix . does not contain main module or its selected dependencies`. |
| `go test` | `/workspace/testing` | Failed | Failed in module mode because the repo has no root module: `go: cannot find main module, but found .git/config in /workspace`. |
| `go test` | `/workspace/benchmark` | Failed | Failed in module mode for the same no-root-module reason. |
| `GO111MODULE=off go test` | `/workspace/testing` | Passed | `PASS`, `ok _/workspace/testing 0.002s`. |
| `GO111MODULE=off go test` | `/workspace/benchmark` | Passed | No test functions were present, but package compiled: `testing: warning: no tests to run`, `PASS`. |
| `go test ./...` | `/workspace/http-server` | Passed | Package compiled; no test files. |
| `go test ./...` | `/workspace/htmx` | Passed | Package compiled; no test files. |
| `go test ./...` | `/workspace/ai/go-groq` | Passed | Package compiled; no test files. |
| `go test ./...` | `/workspace/connect-to-db-postgres` | Passed | Package compiled; no test files. Downloaded `github.com/lib/pq v1.10.9`. |
| `GO111MODULE=off go test ./...` | `/workspace` | Failed | Failed on packages with external dependencies that are not available in GOPATH mode, including Docker, Kafka, MongoDB, Gorilla mux, and PostgreSQL packages. |

## Findings

### Critical

No critical findings were identified during this review.

### High

1. `rest-kafka-mongo-microservice/rest-to-kafka/rest-kafka-sample.go:64-86` creates a Kafka producer per POST and never closes it.
   - Each request calls `kafka.NewProducer`, but the producer is not closed.
   - This can leak broker connections and native resources under repeated requests.

2. `docker/api/main.go:44-49`, `docker/api/main.go:68-73`, `docker/api/main.go:90-96`, and `docker/api/main.go:113-119` build HTML by string concatenation with Docker API data.
   - Image, container, network, and node metadata is written directly into HTML without escaping.
   - If any metadata is attacker-controlled, these handlers can return unsafe HTML.

3. `docker/apijson/main.go:63-68`, `docker/apijson/main.go:85-91`, and `docker/apijson/main.go:108-114` repeat the same unsafe HTML string concatenation pattern.
   - Use `html/template`, JSON responses, or explicit HTML escaping for values rendered into HTML.

4. `htmx/main.go:61-66` panics when `GROQ_API_KEY` is missing.
   - The panic occurs from request handling via `/translate`, so a missing environment variable can crash the server instead of returning an HTTP error.

### Medium

5. `htmx/groq_client.go:131-140` and `ai/go-groq/groq_client.go:131-140` can leak response bodies on non-200 responses.
   - Both functions return on `resp.StatusCode != 200` before registering the `defer resp.Body.Close()`.
   - This prevents connection reuse and can leak resources during retries or repeated failures.

6. `mongo-microservice/cmd/main.go:19-24` defines `Job.Tile` instead of `Job.Title`.
   - Other Mongo/REST examples use `Title`, so inserted document fields may not align across examples.

7. `mongo-microservice/cmd/main.go:36-41` opens an mgo session and never closes it.
   - The example should defer `session.Close()` after a successful connection.

8. `mongo-microservice/rest/connect-mongo-rest-sample.go:70-122` panics inside HTTP handlers.
   - Handler-level panics on JSON, body, or database errors can terminate requests abruptly and make the service brittle.
   - Returning structured HTTP errors would make the sample safer and easier to debug.

9. `json/main.go:12-25` prints unmarshaled data before checking `err` and then uses unchecked type assertions.
    - If the JSON shape changes or unmarshalling fails, the program can panic.
    - Check `err` before using `f`, and use checked assertions before indexing `Sample`.

10. `kafka/producer/main.go:11-40` flushes the producer but never closes it.
    - Add `defer p.Close()` after producer creation to release librdkafka resources.

11. `connect-to-db-postgres/main.go:11-16` hard-codes database credentials.
    - This is acceptable for a toy local example, but should be clearly documented or moved to environment variables before being reused.

12. `connect-to-db-postgres/main.go:42-43` prints user passwords from the database.
    - Even in examples, printing password fields normalizes unsafe handling of sensitive data.

13. `docker/cmd/main.go:47-52` does not close the container logs reader.
    - `ContainerLogs` returns an `io.ReadCloser`; close it after copying logs to avoid leaking file descriptors.

14. `rest-kafka-mongo-microservice/kafka-to-mongo/kafka-mongo-sample.go:35-39` and `mongo-microservice/multistage/main.go:39-47` keep Mongo sessions open for the life of the process without a shutdown path.
    - Long-running examples should close sessions during shutdown and avoid panics for recoverable processing errors.

### Lower Severity / Maintainability

15. The repository has mixed module and non-module layouts.
    - Root-level `go test ./...` cannot run because there is no root module.
    - GOPATH mode covers some standalone examples but fails for examples with external dependencies.
    - A root `go.work` or per-example `go.mod` files would make test/build behavior clearer.

16. Several examples use deprecated or older APIs.
    - Docker samples use `golang.org/x/net/context` instead of the standard `context` package.
    - Docker samples use `client.NewEnvClient`, which has been superseded by `client.NewClientWithOpts`.
    - Mongo samples use `gopkg.in/mgo.v2`, which is unmaintained.

17. Many server examples call `panic` in handlers or request paths.
    - Examples include Mongo, Docker, and translation handlers.
    - Prefer returning `http.Error` and logging server-side details.

## Recommendations

1. Add a root `go.work` file or convert examples into clearly scoped Go modules so CI can run deterministic package tests.
2. Add small tests for the packages that already expose testable functions, especially HTTP handlers and client error handling.
3. Close external resources consistently: HTTP response bodies, Docker clients, Mongo sessions, Kafka producers, and log readers.
4. Avoid `panic` in HTTP handlers; return appropriate status codes and error messages.
5. Escape or template HTML responses that contain values from Docker, database, or request data.
6. Treat secrets and password fields as sensitive even in examples; prefer environment variables and avoid printing password values.
