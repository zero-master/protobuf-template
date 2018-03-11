1. Write a proto file `service.proto`

2. Generate interface

    `protoc -I service/ service/service.proto --go_out=plugins=grpc:service
`

3. Write `server.go`

4. Write `client.go`

5. `go run server.go` in one terminal

6. `go run client.go` in a different terminal