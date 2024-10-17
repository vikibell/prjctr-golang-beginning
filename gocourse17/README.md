# Getting started

1. In order to run the server, open new terminal window and run command:
   `go run server/server.go`
2. In order to test how it works you can use client, for this run command in new terminal window:
   `go run client/client.go`

If you want to recompile proto file, go to the api directory and run command:
`protoc --go_out=../grpcapi --go_opt=paths=source_relative --go-grpc_out=../grpcapi --go-grpc_opt=paths=source_relative api.proto`

See details here: https://grpc.io/docs/languages/go/quickstart/
