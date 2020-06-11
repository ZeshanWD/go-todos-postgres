package helpers

import (
	"strings"
)

func IsValidDescription(description string) bool {
	desc := strings.TrimSpace(description)
	if len(desc) == 0 {
		return false
	}

	return true
}
