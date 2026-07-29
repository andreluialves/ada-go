package username

import "strings"

func normalizeUsername(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
