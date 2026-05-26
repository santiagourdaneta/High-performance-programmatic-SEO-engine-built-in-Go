# Nexus-SEO Engine

High-performance, minimalist programmatic SEO engine built entirely in Go. Designed to run on bare-metal architectures with zero external runtime overhead, delivering fast indexable static layouts for search engine optimization.

## Features

* **Zero Node.js Dependency:** Built on top of Go's native HTTP primitives.
* **Automated Data Structures:** Injects real-time Schema.org JSON-LD definitions directly into target nodes.
* **Multi-Stage Delivery:** Minimalist footprint ready for low-spec servers and instant edge deployment.

## Local Development

Since the environment bypasses virtualization toolchains (such as Docker) for local runs, the workflows operate directly on native system runtimes via standard binaries.

### Prerequisites

* Go 1.22+
* `golangci-lint`
* `pre-commit`

### Setup & Activation

1. Initialize git hooks configuration:
   ```bash
   pre-commit install

Run the local development server:

make dev

Available Workflows

make fmt: Standardizes code formatting.

make lint: Runs strict structural diagnostic tools.

make test: Executes thread-safe race condition sweeps and testing coverage.

make build: Generates a high-performance production binary inside ./bin.

CI/CD Pipeline

Code quality is enforced via automated GitHub Actions on every integration cycle:

Formatter structural verification.

Static semantic analysis (golangci-lint).

Isolated test execution suites.

Distroless-ready production artifact assembly evaluation.