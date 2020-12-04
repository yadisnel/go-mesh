# GRPC Server

The grpc server is a [go-ms.Server](https://godoc.org/github.com/yadisnel/go-ms/server#Server) compatible server.

## Overview

The server makes use of the [google.golang.org/grpc](google.golang.org/grpc) framework for the underlying server 
but continues to use go-ms handler signatures and protoc-gen-go-ms generated code.

## Usage

Specify the server to your go-ms service

```go
import (
        "github.com/yadisnel/go-ms"
        "github.com/yadisnel/go-ms-plugins/server/grpc"
)

func main() {
        service := goms.NewService(
                // This needs to be first as it replaces the underlying server
                // which causes any configuration set before it
                // to be discarded
                goms.Server(grpc.NewServer()),
                goms.Name("greeter"),
        )
}
```
**NOTE**: Setting the gRPC server and/or client causes the underlying the server/client to be replaced which causes any previous configuration set on that server/client to be discarded. It is therefore recommended to set gRPC server/client before any other configuration