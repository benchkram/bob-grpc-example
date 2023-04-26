# bob-grpc example

A simple example on how to generate gRPC code with bob.  

## Prerequisites
- go: https://go.dev/doc/install
- bob + nix: https://bob.build/docs/getting-started/installation

### Generate code from proto definition
```shell
bob build proto
```

### Run server
```shell
go run server/main.go
```

### Run client
```shell
go run client/main.go
```
