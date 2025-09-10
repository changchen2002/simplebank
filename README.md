# Simple Bank API

A modern, secure, and extensible banking API built with Go. Features robust authentication, account management, and money transfer capabilities, following best practices for backend development.

[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-12+-blue.svg)](https://www.postgresql.org/)

---

## Features

- **User Management**: Secure registration, login, and role-based access
- **Account Operations**: Create, view, and list bank accounts
- **Money Transfers**: Transactional transfers between accounts
- **JWT Authentication**: Secure, stateless authentication
- **Database Migrations**: Automated schema management
- **Type-Safe DB Access**: SQLC-generated Go code
- **Comprehensive Testing**: Unit and integration tests
- **Docker Support**: Easy local development

## Tech Stack

- **Language**: Go 1.23+
- **Framework**: Gin (HTTP API)
- **Database**: PostgreSQL 12+
- **ORM/Query**: SQLC
- **Authentication**: JWT, PASETO
- **Validation**: Go validator
- **Config**: Viper
- **Testing**: Go test, Mockgen
- **Migrations**: golang-migrate

## Prerequisites

- Go 1.23+
- PostgreSQL 12+
- Docker (optional)
- Make

## Quick Start

1. **Clone the Repository**
	```bash
	git clone <repo-url>
	cd simplebank
	```
2. **Start PostgreSQL (Docker recommended)**
	```bash
	make postgres
	make createdb
	```
3. **Run Migrations**
	```bash
	make migrateup
	```
4. **Generate SQLC Code**
	```bash
	make sqlc
	```
5. **Start the Server**
	```bash
	make server
	```
	The API runs at `http://localhost:8080`

## Project Structure

```
simplebank/
├── api/         # HTTP handlers & middleware
├── db/          # Migrations, queries, generated code
├── frontend/    # (Optional) Frontend code
├── gapi/        # gRPC API (if enabled)
├── mail/        # Email sending logic
├── token/       # JWT/PASETO token logic
├── util/        # Utilities
├── val/         # Validation helpers
├── worker/      # Background jobs
├── main.go      # Entry point
├── app.env      # Environment config
├── Makefile     # Build/dev commands
```

## Configuration

Create an `app.env` file:
```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
TOKEN_SYMMETRIC_KEY=your-secret-key-here
ACCESS_TOKEN_DURATION=15m
```

## API Endpoints (REST)

- `POST   /users`           - Register user
- `POST   /users/login`     - User login
- `POST   /accounts`        - Create account
- `GET    /accounts/:id`    - Get account
- `GET    /accounts`        - List accounts
- `POST   /transfers`       - Create transfer
- `GET    /transfers`       - List transfers

## Testing

Run all tests:
```bash
make test
```
With coverage:
```bash
go test -v -cover ./...
```
Generate mocks:
```bash
make mock
```

## Development Commands

| Command            | Description                  |
|--------------------|------------------------------|
| make postgres      | Start PostgreSQL container   |
| make createdb      | Create database              |
| make dropdb        | Drop database                |
| make migrateup     | Run all migrations           |
| make migratedown   | Rollback all migrations      |
| make sqlc          | Generate SQLC code           |
| make test          | Run tests                    |
| make server        | Start server                 |
| make mock          | Generate mocks               |

## Security

- **Password Hashing**: bcrypt
- **JWT/PASETO**: Secure token auth
- **Input Validation**: Strict request validation
- **SQL Injection Protection**: Parameterized queries
- **CORS**: Configurable

## Database Schema

- `users`      - User accounts
- `accounts`   - Bank accounts
- `transfers`  - Money transfers
- `entries`    - Transaction entries

## Contributing

1. Fork the repo
2. Create a branch (`git checkout -b feature/xyz`)
3. Commit (`git commit -m 'Add xyz'`)
4. Push (`git push origin feature/xyz`)
5. Open a PR

## License

MIT. See [LICENSE](LICENSE).

## Acknowledgments

- Uses [SQLC](https://sqlc.dev/) and [Gin](https://gin-gonic.com/)
