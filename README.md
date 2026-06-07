# RBK Go Final Project

This is a Go-based project implementing a CRM-like system with CRUD operations for multiple entities using CQRS and GORM.

## Components

- `api/`: Swagger documentation.
- `cmd/`: Entry points for the application.
- `internal/`: Main application logic.
  - `application/`: CQRS Handlers and application services.
  - `container/`: Dependency injection container.
  - `delivery/`: HTTP delivery layer (Gin).
  - `domain/`: Core business logic and entities.
  - `infrastructure/`: Database persistence and external adapters.
