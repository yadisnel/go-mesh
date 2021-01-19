package mock

import (
	"context"

	proto "github.com/yadisnel/go-ms/v2examples/helloworld/proto"
	"github.com/yadisnel/go-ms/v2/client"
)

type mockGreeterService struct {
}

func (m *mockGreeterService) Hello(ctx context.Context, req *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
	return &proto.Response{
		Greeting: "Hello " + req.Name,
	}, nil
}

func NewGreeterService() proto.GreeterService {
	return new(mockGreeterService)
}
