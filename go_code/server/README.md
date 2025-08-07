First install the Go dependencies

```
go mod tidy
```

The KV service offers 2 storage options, in memory and redis. Default is in memory, to use redis, start like this:

```
STORAGE_TYPE=redis go run .
```

otherwise start like:

```
go run .
```

and to test from cli:

set a value

```
curl -X POST "http://localhost:8000/set" \
  -H "Content-Type: application/json" \
  -d '{"key": "testkey", "value": "testvalue"}'
```

get a value

```
curl "http://localhost:8000/get/testkey"
```

delete a key

```
curl -X DELETE "http://localhost:8000/delete/testkey"
```

verify deletion

```
curl "http://localhost:8000/exists/testkey"
``` 