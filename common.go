package utility

import (
	"strings"
	"regexp"
)

func MakeKey(str string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9\-\\]+`)
	str = reg.ReplaceAllString(str, "-")
	str = strings.ToLower(str)
	return str
}