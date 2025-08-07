# Go Key-Value Store

A Gin-based key-value store server with support for multiple storage backends.

## Quick Start

### Install Dependencies

```bash
go mod tidy
```

### Run the Server

The KV service offers 2 storage options: **in-memory** and **redis**.

**Using Redis (recommended):**
```bash
STORAGE_TYPE=redis go run server/main.go
```

**Using In-Memory (default):**
```bash
cd server 
go run .
```

## API Testing

### Set a Value

```bash
curl -X POST "http://localhost:8000/set" \
  -H "Content-Type: application/json" \
  -d '{"key": "testkey", "value": "testvalue"}'
```

### Get a Value

```bash
curl "http://localhost:8000/get/testkey"
```

### Delete a Key

```bash
curl -X DELETE "http://localhost:8000/delete/testkey"
```

### Verify Deletion

```bash
curl "http://localhost:8000/exists/testkey"
```

## Docker Setup

### Run with Docker

```bash
cd go_code
docker-compose up -d --build

# test end to end 
chmod +x test_all.sh
sh test_all.sh
```

### Test Client

The Go implementation includes a test client that automates testing:

```bash
# Test deletion workflow
curl "http://localhost:8002/test_deletion"

# Test overwrite workflow
curl "http://localhost:8002/test_overwrite"

# Test getting a specific key
curl "http://localhost:8002/test_get/mykey"
```

## 🔧 Configuration

### Environment Variables

- `STORAGE_TYPE` - Storage backend (`redis` or `memory`)
- `REDIS_ADDR` - Redis address (default: `localhost:6379`)
- `PORT` - Server port (default: `8000`)

## 📁 Project Structure

```
go_code/
├── server/
│   ├── main.go           # Main application
│   ├── api/routes.go         # API routes
│   ├── init.go           # Service initialization
|   ├── models/responses.go   # Responses
│   ├── services/         # Business logic
│   ├── storage/          # Storage backends
│   └── Dockerfile
├── test_client/
│   ├── main.go           # Test client app
│   ├── api/routes.go         # Test endpoints
|   ├── models/responses.go   # Responses
│   └── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
``` 