package optional

import (
	"strconv"
	"strings"
)

func MapStringsToOptionalInts(values []string) []*int {
	result := make([]*int, len(values))

	for i, value := range values {
		if n, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			result[i] = &n
		} else {
			result[i] = nil
		}
	}

	return result
}
