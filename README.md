# axis-allies-backend
Axis and Allies Classic Golang Backend Servers

## run the server
```
go run main.go
```

# automated testing

## web driver integration tests
```
go clean -testcache
go test ./integration/...
```

## package unit tests
```
go clean -testcache
go test ./pkg/...
```