export PATH=$PATH:/usr/local/go/bin

export PATH=$PATH:$(go env GOPATH)/bin

protoc --go_out=pkg/ --go_opt=paths=source_relative --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative proto/fibonacci.proto

что бы подключиться к GRPC серверу из корневой папки go-fibonacci-api

go run pkg/cmd/grpcserver/main.go

в втором окне Клиент

go run pkg/cmd/grpcclient/main.go 3 5

что бы подключиться к HTTP серверу из корневой папки go-fibonacci-api

go run pkg/cmd/httpserver/main.go

http://localhost:8000/api/fibonacci-calc

POST запрос с json

{ "x": 3, "y": 40 }