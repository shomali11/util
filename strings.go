package util

import (
	"strings"
)

// IsEmpty returns true if the string is empty
func IsEmpty(text string) bool {
	return len(text) == 0
}

// IsNotEmpty returns true if the string is not empty
func IsNotEmpty(text string) bool {
	return !IsEmpty(text)
}

// IsBlank returns true if the string is blank (all whitespace)
func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

// IsNotBlank returns true if the string is not blank
func IsNotBlank(text string) bool {
	return !IsBlank(text)
}

// Reverse reverses the input while respecting UTF8 encoding and combined characters
func Reverse(text string) string {
	textRunes := []rune(text)
	textRunesLength := len(textRunes)
	if textRunesLength <= 1 {
		return text
	}

	reverseRunes := make([]rune, textRunesLength)
	endIndex := textRunesLength - 1

	for i := 0; i < textRunesLength; {
		j := i + 1
		for j < textRunesLength && IsMark(textRunes[j]) {
			j++
		}

		for k := j - 1; k >= i; k-- {
			reverseRunes[endIndex] = textRunes[k]
			endIndex--
		}
		i = j
	}
	return string(reverseRunes)
}
