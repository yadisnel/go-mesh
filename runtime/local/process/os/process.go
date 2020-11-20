// Package os runs processes locally
package os

import (
	"github.com/yadisnel/go-ms/v1/runtime/local/process"
)

type Process struct{}

func NewProcess(opts ...process.Option) process.Process {
	return &Process{}
}
