package utils

import "regexp"

func IsNonRetryableError(err error) bool {
	// Define regular expressions to match error messages
	patterns := []string{
		`connection error: desc = "transport: Error while dialing: dial tcp :4000: connect: connection refused"`,
		`connection error: desc = "transport: Error while dialing: dial tcp :5050: connect: connection refused"`,
	}

	// Convert the error to a string
	errMsg := err.Error()

	for _, pattern := range patterns {
		matched, err := regexp.MatchString(pattern, errMsg)
		if err == nil && matched {
			return false // This error is retryable
		}
	}

	return true
}
