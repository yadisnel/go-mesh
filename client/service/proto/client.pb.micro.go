// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: client/service/proto/client.proto

package go_micro_client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/yadisnel/go-ms/v2/api"
	client "github.com/yadisnel/go-ms/v2/client"
	server "github.com/yadisnel/go-ms/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Client service

func NewClientEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Client service

type ClientService interface {
	// Call allows a single request to be made
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// Stream is a bidirectional stream
	Stream(ctx context.Context, opts ...client.CallOption) (Client_StreamService, error)
	// Publish publishes a message and returns an empty Message
	Publish(ctx context.Context, in *Message, opts ...client.CallOption) (*Message, error)
}

type clientService struct {
	c    client.Client
	name string
}

func NewClientService(name string, c client.Client) ClientService {
	return &clientService{
		c:    c,
		name: name,
	}
}

func (c *clientService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Client.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientService) Stream(ctx context.Context, opts ...client.CallOption) (Client_StreamService, error) {
	req := c.c.NewRequest(c.name, "Client.Stream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &clientServiceStream{stream}, nil
}

type Client_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type clientServiceStream struct {
	stream client.Stream
}

func (x *clientServiceStream) Close() error {
	return x.stream.Close()
}

func (x *clientServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *clientServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *clientServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *clientServiceStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *clientServiceStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientService) Publish(ctx context.Context, in *Message, opts ...client.CallOption) (*Message, error) {
	req := c.c.NewRequest(c.name, "Client.Publish", in)
	out := new(Message)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Client service

type ClientHandler interface {
	// Call allows a single request to be made
	Call(context.Context, *Request, *Response) error
	// Stream is a bidirectional stream
	Stream(context.Context, Client_StreamStream) error
	// Publish publishes a message and returns an empty Message
	Publish(context.Context, *Message, *Message) error
}

func RegisterClientHandler(s server.Server, hdlr ClientHandler, opts ...server.HandlerOption) error {
	type client interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		Publish(ctx context.Context, in *Message, out *Message) error
	}
	type Client struct {
		client
	}
	h := &clientHandler{hdlr}
	return s.Handle(s.NewHandler(&Client{h}, opts...))
}

type clientHandler struct {
	ClientHandler
}

func (h *clientHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.ClientHandler.Call(ctx, in, out)
}

func (h *clientHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.ClientHandler.Stream(ctx, &clientStreamStream{stream})
}

type Client_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type clientStreamStream struct {
	stream server.Stream
}

func (x *clientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *clientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *clientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *clientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *clientStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *clientStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *clientHandler) Publish(ctx context.Context, in *Message, out *Message) error {
	return h.ClientHandler.Publish(ctx, in, out)
}
