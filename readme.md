# To-Work API

A Go-based Backend API for the To-Work application, featuring both REST and GraphQL interfaces.

## 🚀 Overview

This project is a clean-architecture inspired backend service built with Go. It utilizes PostgreSQL for persistent storage and Redis for caching/session management. Database interactions are handled using `sqlc` for type-safe SQL queries.

## 🏗 Project Structure

```text
.
├── .vscode/            # VS Code configuration
├── cmd/                # Application entry points
│   ├── graphql/        # GraphQL server
│   └── rest/           # REST API server (Gin)
│        ├── handler/    # Route handlers
│        ├── middleware/ # Custom middlewares
│        ├── request/    # Request DTOs
│        ├── response/   # Response DTOs
│        ├── server.go   # Server initialization
│        ├── router.go   # Router setup
│        └── main.go     # Application entry point
├── config/             # Configuration logic and DB initialization
├── db/                 # Database schema and queries
│   ├── queries/        # SQL queries for sqlc
│   └── schema/         # SQL migrations/schema
├── docker/             # Dockerfiles for infrastructure
├── internal/           # Private application code
│   ├── shared/         # Shared utilities
│   ├── sqlc/           # Generated code by sqlc
│   └── user/           # User feature logic
│       ├── domain/     # Domain models
│       ├── enum/       # Enumerations
│       └── service/    # Business logic services
├── pkg/                # Public library code
├── .env                # Environment variables (gitignored)
├── .env.example        # Example environment variables
├── docker-compose.yml  # Local infrastructure setup
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
└── sqlc.yaml           # sqlc configuration
```

## 🛠 Setup & Installation

### Prerequisites

- **Go**: 1.24+
- **Docker & Docker Compose**: For running PostgreSQL and Redis
- **sqlc**: For SQL code generation (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)

### Steps

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd to-work-api
   ```

2. **Environment Configuration**
   Copy the example environment file and update the values as needed:

   ```bash
   cp .env.example .env
   ```

3. **Start Infrastructure**
   Launch PostgreSQL and Redis using Docker Compose:

   ```bash
   docker-compose up -d
   ```

4. **Generate SQL Code**
   Generate type-safe Go code from your SQL queries:

   ```bash
   sqlc generate
   ```

5. **Install Dependencies**
   ```bash
   go mod tidy
   ```

## 🚀 Running the Application

### REST API

```bash
go run cmd/rest/main.go
```

The server will be available at `http://localhost:8080`.

### GraphQL API (Placeholder)

```bash
go run cmd/graphql/main.go
```

## 🧹 Code Quality & Formatting

This project uses specific tools to maintain code quality:

- **Formatting**: `gofumpt` is used for stricter Go formatting.
  ```bash
  # Inside VS Code, it's configured to run on save.
  # Manually:
  go install mvdan.cc/gofumpt@latest
  gofumpt -l -w .
  ```
- **Linting**: `golangci-lint` is used for static analysis.
  ```bash
  golangci-lint run
  ```

### VS Code Settings

VS Code is pre-configured in `.vscode/settings.json` to use:

- `gofumpt` for formatting
- `golangci-lint` for linting
- `gopls` as the language server
