# Go Fiber Boilerplate

This project is a boilerplate template for building web applications using Go and the Fiber framework. It includes setup for database integration with PostgreSQL and provides a basic structure for implementing CRUD operations for user management.

## Prerequisites

- Docker
- Docker Compose

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/karabayyazilim/go-fiber-boilerplate.git
    cd go-fiber-boilerplate
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up your environment variables. Create a `.env` file in the root directory and add the following:
    ```dotenv
    APP_PORT=8000

    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=go_fiber
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_SSL_MODE=disable
    DB_TIMEZONE=Europe/Istanbul
    ```

## Running the Application

1. Start the PostgreSQL database using Docker Compose:
    ```sh
    docker-compose up -d
    ```

2. Run the application:
    ```sh
    go run src/cmd/app/main.go
    ```

3. The application will be available at `http://localhost:8000`.

## Project Structure

- `src/internal/config`: Contains configuration files for the database.
- `src/internal/models`: Contains the data models.
- `src/internal/services`: Contains the service layer for business logic.
- `src/pkg/paginate`: Contains pagination logic.

## API Endpoints

### User Endpoints

- **GET /users**: List all users with pagination.
- **POST /users**: Create a new user.
- **GET /users/:id**: Get a user by ID.
- **PUT /users/:id**: Update a user by ID.
- **DELETE /users/:id**: Delete a user by ID.
