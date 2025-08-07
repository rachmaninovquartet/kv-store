## 🚀 Quick Start

### Python Implementation

```bash
# in  py_code/

# Run with Docker (recommended)
docker-compose up -d --build

# Or run locally in two terminals
conda activate censys-env
STORAGE_TYPE=redis python server/server.py # omit storage type to run in memory kv store

conda activate censys-env
SERVER_URL=http://localhost:8000 python test_client/client.py 
```

## 📋 SERVER API Endpoints

Both implementations provide the same REST API:

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/set` | Set a key-value pair |
| `GET` | `/get/{key}` | Get value by key |
| `DELETE` | `/delete/{key}` | Delete a key |
| `GET` | `/exists/{key}` | Check if key exists |

### Example Usage

```bash
# Set a key-value pair
curl -X POST "http://localhost:8000/set" \
  -H "Content-Type: application/json" \
  -d '{"key": "mykey", "value": "myvalue", "ttl": 3600}'

# Get a value
curl "http://localhost:8000/get/mykey"

# Delete a key
curl -X DELETE "http://localhost:8000/delete/mykey"
```

## 🐳 Docker Setup

### Python Services

The Python implementation includes three Docker services:

- **Redis** (port 6379) - Database
- **Server** (port 8000) - Main API
- **Test Client** (port 8002) - Test automation

```bash
# in py_code/
docker-compose up -d --build

# verify everything works
chmod +x test_all.sh
sh test_all.sh

```

### Test Client Endpoints

| Method |	Endpoint |	Description |
|--------|-----------|--------------|
| GET |	/test_deletion | Test deletion workflow (set → check → delete) |	
| GET	|	/test_overwrite	|	Test overwrite workflow (set → overwrite → verify) |	
| GET	|	/test_get/{key}	|	Test getting a specific key from server |	

The Python implementation includes a test client that automates testing:

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
- `REDIS_HOST` - Redis host (default: `localhost`)
- `REDIS_PORT` - Redis port (default: `6379`)


## 🧪 Testing

### Manual Testing

```bash
# Test the main server
curl "http://localhost:8000/"
curl -X POST "http://localhost:8000/set" \
  -H "Content-Type: application/json" \
  -d '{"key": "test", "value": "hello"}'
curl "http://localhost:8000/get/test"
```

### API Documentation

FastAPI automatically generates interactive API documentation:

- **Swagger UI**: http://localhost:8000/docs
- **ReDoc**: http://localhost:8000/redoc

## 🛠️ Development

### Prerequisites

- **Python**: 3.11+ with conda environment
- **Docker**: 20.10+
- **Redis**: 7.0+ (optional, for Redis backend)
