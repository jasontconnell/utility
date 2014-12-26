package utility

import (
	"strings"
	"regexp"
)

func MakeKey(str string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9\-\\ ]+`)
	spacereg := regexp.MustCompile(`[ ]+`)
	str = reg.ReplaceAllString(str, "")
	str = strings.Trim(str, " ")
	str = spacereg.ReplaceAllString(str, " ")
	str = spacereg.ReplaceAllString(str, "-")
	str = strings.ToLower(str)
	return str
}
