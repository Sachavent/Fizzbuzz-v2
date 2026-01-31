# Fizzbuzz-v2 🧩

### Context 🧭

Fizzbuzz API that generates the result and displays the most called route.

### Prerequisites ✅

Go 1.24 installed.

### Available commands 📋

- `make help` displays all commands.

### Run the project 🚀

- `make run` or `make run-dev` for live reload.

### Quickstart ⚡

- `make run`
- `curl "http://localhost:8080/fizzbuzz/result?int1=3&int2=5&limit=15&str1=fizz&str2=buzz"`

### Endpoints 📡

- `GET /health`
- `GET /fizzbuzz/result`
- `GET /fizzbuzz/stats`
- `GET /swagger/*`

### Swagger documentation 🧪

- `make swagger` then the documentation is available at `http://localhost:8080/swagger/index.html`.

### Lint 🧹

- `make lint` runs golangci-lint (auto-installs if missing).
- `make lint-fix` runs golangci-lint with `--fix`.

### Tests 🧪

- `make test-unit` runs unit tests.
- `make test-coverage` runs unit tests and generates the coverage report (`coverage.html`).
- `make test-integration` runs integration tests.

### Default port 🔌

The API listens on `http://localhost:8080`.
