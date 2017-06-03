package util

import "strings"

func IsEmpty(text string) bool {
	return len(text) == 0
}

func IsNotEmpty(text string) bool {
	return !IsEmpty(text)
}

func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

func IsNotBlank(text string) bool {
	return !IsBlank(text)
}
