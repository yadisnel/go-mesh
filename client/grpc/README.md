# GRPC Client

The grpc client is a [goms.Client](https://godoc.org/github.com/yadisnel/go-ms/client#Client) compatible client.

## Overview

The client makes use of the [google.golang.org/grpc](google.golang.org/grpc) framework for the underlying communication mechanism.

## Usage

Specify the client to your go-ms service

```go
import (
	"github.com/yadisnel/go-ms"
	"github.com/yadisnel/go-ms-plugins/client/grpc"
)

func main() {
	service := goms.NewService(
		goms.Name("greeter"),
		goms.Client(grpc.NewClient()),
	)
}
```
