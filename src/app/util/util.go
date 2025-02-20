package util

import (
	"regexp"
)

func ContainsNumbers(s string) bool {
	match, _ := regexp.MatchString(`[0-9]`, s)
	return match
}
