package models

import "time"

type RetryConfig struct {
	MaxRetries    int
	MaxDuration   time.Duration
	RetryInterval time.Duration
}
