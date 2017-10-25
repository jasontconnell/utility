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

func HasString(list []string, find string) bool {
	for _, i := range list {
		if i == find { return true }
	}
	return false
}