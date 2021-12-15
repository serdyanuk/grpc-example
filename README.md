# Run server
```
$ go run server/main.go
```
# Run client
```
$ go run client/main.go
```
# Generate server and client
```
$ protoc --go_out=. --go-grpc_out=.  products.proto
```
# Prerequisites
<https://grpc.io/docs/languages/go/quickstart/#prerequisites>