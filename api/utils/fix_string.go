package utils

import "strings"

func MakeStringUniformForSignUp(str string, protectSpace bool) string {
	if !protectSpace {
		str = strings.ReplaceAll(str, " ", "")
	}
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	return str
}
