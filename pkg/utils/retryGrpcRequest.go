package utils

import (
	"context"
	"errors"
	"gateway/pkg/common/models"
	"time"
)

func RetryOperation(ctx context.Context, retryConfig models.RetryConfig, operation func() (interface{}, error)) (interface{}, error) {
	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}

		result, err := operation()
		if err == nil {
			return result, nil
		}

		if IsNonRetryableError(err) {
			return nil, err
		}
	}

	return nil, errors.New("exceeded maximum retries")
}
