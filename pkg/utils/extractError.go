package utils

import "strings"

func ExtractErrorMessage(fullErrorMessage string) string {

	parts := strings.Split(fullErrorMessage, "desc =")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1])
	}
	return fullErrorMessage
}
