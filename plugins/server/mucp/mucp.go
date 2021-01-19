// Package mucp provides an mucp server
package mucp

import (
	"github.com/yadisnel/go-ms/v2/cmd"
	"github.com/yadisnel/go-ms/v2/server"
)

func init() {
	cmd.DefaultServers["mucp"] = NewServer
}

// NewServer returns a micro server interface
func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}
