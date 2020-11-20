package grpc

import (
	"crypto/tls"

	gc "github.com/yadisnel/go-ms/v1/client/grpc"
	gs "github.com/yadisnel/go-ms/v1/server/grpc"
	"github.com/yadisnel/go-ms/v1/service"
)

// WithTLS sets the TLS config for the service
func WithTLS(t *tls.Config) service.Option {
	return func(o *service.Options) {
		o.Client.Init(
			gc.AuthTLS(t),
		)
		o.Server.Init(
			gs.AuthTLS(t),
		)
	}
}
