// Package mucp provides an mucp client
package mucp

import (
	"github.com/yadisnel/go-ms/v2/client"
)

// NewClient returns a new go-ms client interface
func NewClient(opts ...client.Option) client.Client {
	return client.NewClient(opts...)
}
