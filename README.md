# Inventory Management System (IMS)

## Project Structure

Below is the structure of the `go-ims` project, reflecting its actual layout:

```
go-ims/
├── cmd/                # Application entry points
├── config/             # Configuration logic
├── conn/               # Database connection setup
├── consts/             # Application contestants
├── controllers/        # HTTP request handlers (controllers)
├── docs/               # Postman collection and database dump
├── domain/             # Repository and service interface definitions
├── env/                # Environment-specific configuration files
├── logger/             # Logger utility
├── middlewares/        # Middleware functions for request processing
├── models/             # Data models and structures
├── repositories/       # Database interaction logic
├── routes/             # API route definitions
├── server/             # Server setup and initialization
├── services/           # Core business logic
├── types/              # Shared type definitions
├── utils/              # Shared utility functions
├── worker/             # Worker pool and tasks
├── go.mod              # Go module definition
├── go.sum              # Dependency checksum file
├── main.go             # Main go file
└── README.md           # Project documentation
```

## Tech Stack

- Language : Golang
- Database : Mysql
- Cache: Redis
- Asynq Queue: Asynq
- Config Management: Consul
- Library
  - [Cobra](https://github.com/spf13/cobra) - Framework for building CLI applications
  - [Viper](https://github.com/spf13/viper) - Library for managing configuration files and environment variables
  - [GORM](https://github.com/go-gorm/gorm) - ORM library for interacting with relational databases
  - [Echo](https://github.com/labstack/echo) - Web framework used for building RESTful APIs and handling HTTP routing
  - [Ozzo-Validation](https://github.com/go-ozzo/ozzo-validation) - Library used for validating input data, ensuring data integrity and consistency
  - [Asynq](https://github.com/hibiken/asynq) - Library for managing background tasks and distributed task queues

## Install dependencies

```bash
go mod tidy
go mod vendor
```

## Run consul service (optional)

```bash
consul agent -config-file="C:\Users\Abid\consul\config\consul.hcl"
```

## Build project

```bash
go build -o app .
```

## Run project server

```bash
./app serve
```

## Run project worker

```bash
./app worker
```
