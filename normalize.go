package sodium

import (
	"strings"
	"unicode"
)

func NormalizeUsername(raw string) string {
	s := strings.TrimSpace(raw)
	s = strings.TrimPrefix(s, "@")
	return strings.ToLower(s)
}

func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}
