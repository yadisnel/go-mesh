package grpc

import (
	"github.com/yadisnel/go-ms/v2/errors"
	"google.golang.org/grpc/status"
)

func gomsError(err error) error {
	// no error
	switch err {
	case nil:
		return nil
	}

	if verr, ok := err.(*errors.Error); ok {
		return verr
	}

	// grpc error
	s, ok := status.FromError(err)
	if !ok {
		return err
	}

	// return first error from details
	if details := s.Details(); len(details) > 0 {
		return gomsError(details[0].(error))
	}

	// try to decode go-ms *errors.Error
	if e := errors.Parse(s.Message()); e.Code > 0 {
		return e // actually a go-ms error
	}

	// fallback
	return errors.InternalServerError("go.ms.client", s.Message())
}
