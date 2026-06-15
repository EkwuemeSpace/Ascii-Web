package main

import (
	"fmt"
	"strings"
)

func renderString(charMap map[rune][]string, input string) (string, error) {
	var result strings.Builder

	for _, ch := range input {
		if !isValidChar(ch) {
			return "", fmt.Errorf("error: invalid character detected %q", ch)
		}
	}
	segment := strings.Split(input, "\n")

	if len(segment) > 1 && segment[0] == "" && segment[len(segment)-1] == "" {
		segment = segment[:len(segment)-1]
	}

	for _, seg := range segment {
		if seg == "" {
			result.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range seg {
				art, ok := charMap[ch]
				if ok {
					result.WriteString(art[row])
				}
			}
			result.WriteString("\n")

		}
	}
	return result.String(), nil
}
