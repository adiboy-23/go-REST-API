# Go REST API

A simple RESTful API built with Go that manages student records using SQLite as the database.

## Features

- CRUD operations for student records
- SQLite database integration
- Graceful shutdown
- Configuration management
- Request validation
- Structured logging
- Error handling

## Project Structure
```
.
├── cmd/
│   └── go-REST-API/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/                     # Configuration management
│   │   └── config.go
│   ├── http/
│   │   └── handlers/
│   │       └── student/            # HTTP request handlers
│   │           └── student.go
│   ├── storage/                    # Database interfaces and implementations
│   │   ├── storage.go             # Storage interface
│   │   └── sqlite/
│   │       └── sqlite.go          # SQLite implementation
│   ├── types/                      # Data models
│   │   └── types.go
│   └── utils/
│       └── response/               # Utility functions
│           └── response.go
├── config/
│   └── local.yaml                  # Configuration file
└── storage/                        # Database files
    └── storage.db
```


## API Endpoints

| Method | Endpoint              | Description          |
|--------|-----------------------|----------------------|
| POST   | `/api/students`       | Create a new student |
| GET    | `/api/students/{id}`  | Get student by ID    |
| GET    | `/api/students`       | Get all students     |


### Prerequisites

- Go 1.22 or higher
- SQLite3

### Installation

1. Clone the repository:
    ```
        git clone https://github.com/yourusername/go-REST-API.git
    ```
go mod download

3. Create a configuration file `config/local.yaml`:
    ```
        yaml
        env: "dev"
        storage_path: "storage/storage.db"
        http_server:
        address: "localhost:8082"
    ```
### Running the Application
```
        go run cmd/go-REST-API/main.go -config config/local.yaml
```

## API Usage Examples

### Create a Student
```bash
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 20
  }'
```

Expected Response:
```json
{
    "id": 1
}
```

### Get a Student by ID
```bash
curl http://localhost:8082/api/students/1
```

Expected Response:
```json
{
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 20
}
```

### List All Students
```bash
curl http://localhost:8082/api/students
```

Expected Response:
```json
[
    {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "age": 20
    },
    {
        "id": 2,
        "name": "Jane Smith",
        "email": "jane@example.com",
        "age": 22
    }
]
```

## Configuration

The application can be configured using:
- Environment variables (CONFIG_PATH)
- YAML configuration file
- Command line flags (-config)

### Environment Variables
```bash
export CONFIG_PATH=config/local.yaml
go run cmd/go-REST-API/main.go
```

### Command Line
```bash
go run cmd/go-REST-API/main.go -config config/local.yaml
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.


