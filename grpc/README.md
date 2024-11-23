# Pre-requisite

1. Protobuf
    - Windows:  
    Link: https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/
1. gRPC
    - ```
        $ go get -u google.golang.org/grpc
        ```
1. protoc-gen-go-grpc
    - ```
        $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        ```
1. protoc-gen-go
    - ```
        $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        ```

# Generate Protobuf
Masuk ke folder grpc, lalu jalankan perintah berikut:
``` 
$ protoc --go_out=. --go-grpc_out=. proto/*
```

# TLS Configuration Example

Server: https://github.com/olivere/grpc-example/blob/master/cmd/server/main.go 

Client: https://github.com/olivere/grpc-example/blob/master/cmd/client/main.go 