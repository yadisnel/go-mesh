package client

import (
	"context"
	"time"

	"github.com/yadisnel/go-ms/v1/util/backoff"
)

type BackoffFunc func(ctx context.Context, req Request, attempts int) (time.Duration, error)

func exponentialBackoff(ctx context.Context, req Request, attempts int) (time.Duration, error) {
	return backoff.Do(attempts), nil
}
