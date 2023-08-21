package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTraceID() string {
	timestamp := time.Now().UnixNano()
	randomNumber := rand.Intn(1000000)
	traceID := fmt.Sprintf("%d-%d", timestamp, randomNumber)
	return traceID
}
