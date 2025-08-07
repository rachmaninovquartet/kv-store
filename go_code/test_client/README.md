First install the Go dependencies

```
go mod tidy
```

Start the client like this:

```
SERVER_URL=http://localhost:8000 go run .
```

and to test from cli:

set a value

```
curl -X POST "http://localhost:8000/set" \
  -H "Content-Type: application/json" \
  -d '{"key": "testkey", "value": "testvalue"}'
```

Test deletion workflow

```
curl "http://localhost:8002/test_deletion"
```

Test overwrite workflow

```
curl "http://localhost:8002/test_overwrite"
```

Test getting a specific key

```
curl "http://localhost:8002/test_get/mykey"
``` 