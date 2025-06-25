# Simple Bank API

A RESTful banking API built with Go, featuring secure authentication, account management, and money transfer capabilities.

[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-12+-blue.svg)](https://www.postgresql.org/)

## 🚀 Features

- **User Management**: Secure user registration and authentication
- **Account Operations**: Create and manage bank accounts
- **Money Transfers**: Secure transfer between accounts with transaction support
- **JWT Authentication**: Token-based authentication with configurable duration
- **Database Migrations**: Automated schema management with golang-migrate
- **Code Generation**: Type-safe database operations with SQLC
- **Comprehensive Testing**: Unit tests with high coverage
- **Docker Support**: Easy development setup with Docker

## 🛠️ Tech Stack

- **Language**: Go 1.23
- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL 12+
- **ORM**: SQLC (SQL Compiler)
- **Authentication**: JWT tokens
- **Validation**: Go validator
- **Configuration**: Viper
- **Testing**: Go testing + Mockgen
- **Migrations**: golang-migrate

## 📋 Prerequisites

- Go 1.23 or higher
- PostgreSQL 12 or higher
- Docker (optional, for easy setup)
- Make (for using Makefile commands)

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone <repository-url>
cd bank
```

### 2. Set Up Database

Using Docker (recommended):
```bash
# Start PostgreSQL container
make postgres

# Create database
make createdb
```

Or manually:
- Install PostgreSQL
- Create a database named `simple_bank`
- Update connection string in `app.env`

### 3. Run Database Migrations

```bash
make migrateup
```

### 4. Generate SQL Code

```bash
make sqlc
```

### 5. Start the Server

```bash
make server
```

The API will be available at `http://localhost:8080`

## 📁 Project Structure

```
bank/
├── api/                 # HTTP API handlers and middleware
│   ├── account.go      # Account-related endpoints
│   ├── transfer.go     # Transfer-related endpoints
│   ├── user.go         # User authentication endpoints
│   ├── server.go       # Server configuration
│   └── middleware.go   # Authentication middleware
├── db/                 # Database layer
│   ├── migration/      # Database migration files
│   ├── query/          # SQL queries for SQLC
│   ├── sqlc/           # Generated SQLC code
│   └── mock/           # Mock implementations for testing
├── util/               # Utility functions
├── token/              # JWT token management
├── val/                # Validation utilities
├── main.go            # Application entry point
├── app.env            # Environment configuration
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
├── sqlc.yaml          # SQLC configuration
└── Makefile           # Build and development commands
```

## 🔧 Configuration

Create an `app.env` file in the root directory:

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
TOKEN_SYMMETRIC_KEY=your-secret-key-here
ACCESS_TOKEN_DURATION=15m
```

## 📚 API Endpoints

### Authentication
- `POST /users` - Create a new user
- `POST /users/login` - User login

### Accounts
- `POST /accounts` - Create a new account
- `GET /accounts/:id` - Get account details
- `GET /accounts` - List user accounts

### Transfers
- `POST /transfers` - Create a money transfer
- `GET /transfers` - List transfers

## 🧪 Testing

Run all tests:
```bash
make test
```

Run tests with coverage:
```bash
go test -v -cover ./...
```

Generate mocks:
```bash
make mock
```

## 🛠️ Development Commands

| Command | Description |
|---------|-------------|
| `make postgres` | Start PostgreSQL container |
| `make createdb` | Create database |
| `make dropdb` | Drop database |
| `make migrateup` | Run all migrations |
| `make migratedown` | Rollback all migrations |
| `make migrateup1` | Run one migration |
| `make migratedown1` | Rollback one migration |
| `make sqlc` | Generate SQLC code |
| `make test` | Run tests |
| `make server` | Start development server |
| `make mock` | Generate mocks |

## 🔒 Security Features

- **Password Hashing**: bcrypt for secure password storage
- **JWT Tokens**: Secure authentication with configurable expiration
- **Input Validation**: Comprehensive request validation
- **SQL Injection Protection**: Parameterized queries via SQLC
- **CORS Support**: Configurable cross-origin resource sharing

## 📊 Database Schema

The application uses the following main tables:
- `users` - User accounts and authentication
- `accounts` - Bank accounts
- `transfers` - Money transfer records
- `entries` - Account transaction entries

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built following the [Tech School](https://github.com/techschool) tutorial
- Uses [SQLC](https://sqlc.dev/) for type-safe database operations
- Powered by [Gin](https://gin-gonic.com/) web framework
