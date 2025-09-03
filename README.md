# Go Template

A clean architecture template for building scalable Go applications with **Gin**, **GORM**, **PostgreSQL**, **Docker**, and **Migrations**.  
This project is organized to enforce separation of concerns and maintainability.

---

## Features

- **Clean Architecture** — Domain-driven design with clear boundaries  
- **Gin Framework** — Fast and lightweight HTTP router  
- **GORM** — ORM for database access  
- **PostgreSQL** — Reliable SQL database  
- **Database Migrations** — Managed with `migrate` tool  
- **Dependency Injection (DI)** — Centralized wiring of components  
- **Docker Support** — For local development and deployment  
- **Utility Helpers** — Pagination, email validation, standardized responses  

---

## Project Structure

```
├── cmd/
│ └── server/                                   # Application entry point
│ └── main.go
├── docker/                                     # Docker & Compose setup
│ ├── Dockerfile
│ └── docker-compose.yml
├── docs/                                       # Auto-generated API docs (hidden from Git)
├── internal/
│ ├── domain/                                   # Core business models
│ │ └── user.go
│ ├── infrastructure/                           # Framework & driver code
│ │ ├── db/                                     # Database connection
│ │ │ └── db.go
│ │ ├── di/                                     # Dependency injection
│ │ │ └── app.go
│ │ ├── repository/                             # GORM repository implementations
│ │ │ └── gorm_user_repo.go
│ │ └── router/                                 # Gin router setup
│ │ └── router.go
│ ├── interface/                                # Adapters (controllers, repositories)
│ │ ├── controller/
│ │ │ ├── dto/                                  # Request/response DTOs
│ │ │ │ ├── auth_dto.go
│ │ │ │ └── pagination_dto.go
│ │ │ └── auth_controller.go
│ │ └── repository/                             # Repository interface contracts
│ │ └── user_repository.go
│ └── usecase/                                  # Business use cases
│ └── user_usecase.go
├── migrations/                                 # SQL migration files
│ ├── 000001_create_users_table.up.sql
│ └── 000001_create_users_table.down.sql
├── pkg/                                        # Shared utilities
│ ├── config/                                   # App configuration loader
│ │ └── config.go
│ ├── response/                                 # Standard API responses
│ │ └── response.go
│ └── utils/                                    # Utility functions
│ ├── is_valid_email.go
│ └── pagination.go
├── .dockerignore
├── .env                                        # Environment variables (hidden from Git)
├── .gitignore
├── .golangci.yml                               # Linter configuration
├── Makefile                                    # Migrateion script cli
├── README.md
├── go.mod
└── go.sum
```
