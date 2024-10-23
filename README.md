# Scoreplay API

## Description
This project is a Domain Driven Design (DDD) implementation of a REST API built in **Go** using net/http, GORM, and PostgreSQL. It provides a basic service for managing **Tags** and **Media**

## Features
- **Tag Management**: Create, list, and search tags by name.
- **Image Management**: Upload images, assign tags, search images by tag, and retrieve image download links.
- **File Upload**: Handles image file uploads via multipart form and stores files in the /uploads directory.
- **Swagger Documentation**: Automatically generated Swagger documentation for API endpoints.
- **Unit Testing**: Comprehensive unit tests for handlers, services, and repositories using httptest and testify.

## Technology Stack
- Go 1.22+
- PostgreSQL
- GORM (ORM)
- Swagger(Swaggo) for API documentation
- testify for unit testing

## Setup

### Prerequisites
1. GO
2. Docker(docker compose) or local installation of PostgreSQL and PGAdmin
3. Git
4. Swaggo

### Steps
1. Clone the repository
```bash
git clone https://github.com/dhelic98/scoreplay-api
cd scoreplay-api
```
2. Setup environment variables(.env.template for reference)
```
PORT=INTEGER
ENV=development | production | integration
HOST= 
DB_CONNECTION_STRING=
FILE_HOST_URL=
CURRENT_API_VERSION=
```
3. Install dependencies
```
go mod tidy
```
4. Start docker containers **if using docker**
```
docker-compose -f ./docker-compose.dev.yml up -d
```
5. Install swaggo
```
go install github.com/swaggo/swag/cmd/swag@latest
```
6. Build and Run the application
```
go build -o scoreplay ./cmd/server  && ./scoreplay
```

Server is not up and running on ``http://localhost:4200``

## Testing

Run tests
```
go test ./... -coverprofile=coverage.out

```

Check out test coverage
```
 go tool cover -html=coverage.out
```

## Swagger 

To generate swagger definitions we used swaggo

Run
```
swag init -g cmd/server/main.go 
```

## Author
Dzenan Helic