# Golang E-Commerce API

A simple e-commerce API built with Go, featuring JWT-based authentication, PostgreSQL, and Swagger documentation.

## Features

- **Registration & Login**  
  - Stores user passwords with bcrypt hashing  
  - Sets JWT token in an `access_token` cookie
- **Role-Based Access Control**
    - Protected seller management endpoints
    - Token-based authentication for API access
- **Swagger Documentation**  
  - Accessible at [http://localhost:8081/swagger/](http://localhost:8081/swagger/)
- **Containerized Deployment**
  - Fully Dockerized setup for easy deployment
  - Supports both local and production environments
- **Testing**
    - **Unit Tests** for isolated function validation
    - **Mock-Based Tests** for repository and service layer
    - **Integration Tests** to validate API behavior

## Quick Start

1. Clone the Repository
    ```bash
    git clone https://github.com/horoshi10v/tested.git
    cd 01-server
    ```
2. **Necessary Step:** change environment variables in `.env`.
3. **Build & Run:**
   ```bash
   docker compose up --build
   ```
4. **Run Tests**:
   ```bash
   docker compose exec app go test ./tests/... -v
   ```
5. **Swagger UI**:
    - Open [http://localhost:8081/swagger/](http://localhost:8081/swagger/) to explore the API.
### Note

- **Persistent Storage**: By default, a named Docker volume is used. Check `docker-compose.yml` for an alternative local directory binding.

