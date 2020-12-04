// Package mucp provides an mucp server
package mucp

import (
	"github.com/yadisnel/go-ms/v2/server"
)

// NewServer returns a go-ms server interface
func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}
