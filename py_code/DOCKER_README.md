# Docker Setup for Censys Test

This project includes Docker containers for the server, test client, and Redis database.

## 🐳 Services

- **Redis**: Database for key-value storage
- **Server**: Main API server (port 8000)
- **Test Client**: Test client service (port 8002)

## 🚀 Quick Start

### Build and run all services:
```bash
cd py_code
docker-compose up --build
```

### Run in background:
```bash
cd py_code
docker-compose up -d --build
```

### Stop all services:
```bash
cd py_code
docker-compose down
```

## 📋 Service URLs

- **Main Server**: http://localhost:8000
- **Test Client**: http://localhost:8002
- **Redis**: localhost:6379

## 🔧 Individual Service Commands

### Build specific service:
```bash
cd py_code
docker-compose build server
docker-compose build test-client
```

### Run specific service:
```bash
cd py_code
docker-compose up server
docker-compose up test-client
```

### View logs:
```bash
cd py_code
docker-compose logs server
docker-compose logs test-client
docker-compose logs redis
```

## 🧪 Testing

### Test the main server:
```bash
curl "http://localhost:8000/"
curl -X POST "http://localhost:8000/set" \
  -H "Content-Type: application/json" \
  -d '{"key": "test", "value": "hello"}'
```

### Test the test client:
```bash
curl "http://localhost:8002/test_deletion"
curl "http://localhost:8002/test_overwrite"
```

## 🔍 Docker Network

All services communicate through the `censys-network` bridge network:
- Server can reach Redis at `redis:6379`
- Test client can reach server at `server:8000`

## 📁 File Structure

```
censys-test/
├── py_code/
│   ├── docker-compose.yml
│   ├── DOCKER_README.md
│   ├── server/
│   │   ├── Dockerfile
│   │   ├── .dockerignore
│   │   └── requirements.txt
│   └── test_client/
│       ├── Dockerfile
│       ├── .dockerignore
│       └── requirements.txt
└── go_code/
```
