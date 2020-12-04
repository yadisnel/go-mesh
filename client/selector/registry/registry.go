// Package registry uses the go-ms registry for selection
package registry

import (
	"github.com/yadisnel/go-ms/v2/client/selector"
)

// NewSelector returns a new registry selector
func NewSelector(opts ...selector.Option) selector.Selector {
	return selector.NewSelector(opts...)
}
