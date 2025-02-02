## Playing with Go Chi

This repository is a playground for me to experiment with Go by creating an REST API for a forum website 


## Local Development

You must have `go`, `docker` and `make` installed.

To run the application:

```bash
make up
```

we have a volume set up between our project code and the container. Any change to the code is reflected in the container, and we use CompileDaemon to watch our changes.

### generate ent
we use ent framework as a our ORM. \
Before running the app for the first time, we need to:
```bash
go generate ./repositories
```
everytime we add/delete/update a file inside repositories/schema, we should run this command.

## API Testing (TODO:)
API test requests are located in `test/api/api.http`. Use the VS Code REST Client extension to execute these tests.


## Project Architecture

This project follows a 3-Layer Architecture pattern with the following components:


### 1. Handler Layer (Presentation Layer)
- Location: `./handlers/`
- Responsibilities:
  - HTTP request/response handling
  - Request/input parsing and validation
  - Route management
  - HTTP error responses
  - HTTP-specific concerns (headers, status codes, etc.)
- Should not contain:
  - Business logic
  - Data access logic

### 2. Service Layer (Business Layer)
- Location: `./services/`
- Responsibilities:
  - Business logic implementation
  - Business rules validation
  - Orchestrating multiple repositories
  - Transaction management
  - Integration with external services
  - Domain error handling
- Should not contain:
  - HTTP-specific logic
  - Data access logic

### 3. Repository Layer (Data Access Layer)
- Location: `./repository/`
- Responsibilities:
  - Database operations
  - Data persistence logic
  - Query building
  - Data mapping
  - Database error handling
- Should not contain:
  - Business logic
  - HTTP-specific logic
