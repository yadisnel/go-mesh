package subscriber

import (
	"context"

	"github.com/yadisnel/go-ms/v2/util/log"

	example "github.com/yadisnel/go-ms/v2examples/template/fnc/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
