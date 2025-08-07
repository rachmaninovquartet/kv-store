# Key-Value Store

A key-value store with multiple language implementations and storage backends.

## 🏗️ Architecture

This project consists of two complete implementations of a key-value store:

- **Python Implementation** (`py_code/`) - FastAPI-based REST API
- **Go Implementation** (`go_code/`) - Go-based REST API

Each implementation supports multiple storage backends:
- **Redis** - Persistent, distributed storage
- **In-Memory** - Fast, ephemeral storage (default)

## See the individual README in each of the *_code/ dirs for specific instructions



## 📁 Project Structure

```
censys-test/
├── README.md                 # This file
├── py_code/                  # Python implementation
│   ├── docker-compose.yml    # Docker orchestration
│   ├── DOCKER_README.md      # Docker documentation
│   ├── server/               # Main API server
│   │   ├── Dockerfile
│   │   ├── server.py         # FastAPI application
│   │   ├── api/routes.py     # API endpoints
│   │   ├── services/         # Business logic
│   │   ├── storage/          # Storage backends
│   │   └── models/           # Data models
│   └── test_client/          # Test automation
│       ├── Dockerfile
│       ├── client.py         # Test client app
│       └── api/routes.py     # Test endpoints
└── go_code/                  # Go implementation
    ├── main.go               # Main application
    ├── handlers/             # HTTP handlers
    ├── storage/              # Storage backends
    └── models/               # Data structures
```