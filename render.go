package main

import (
	"fmt"
	"io"
	"strings"
)

func renderString(w io.Writer, charMap map[rune][]string, input string) error {
	var result strings.Builder

	for _, ch := range input {
		if !isValidChar(ch) {
			return fmt.Errorf("error: invalid character detected %q", ch)
		}
	}
	segment := strings.Split(input, `\n`)
	if len(segment) > 1 && segment[0] == "" && segment[len(segment)-1] == "" {
		segment = segment[:len(segment)-1]
	}

	for _, seg := range segment {
		if seg == "" {
			fmt.Fprintln(w)
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range seg {
				art, ok := charMap[ch]
				if ok {
					result.WriteString(art[row])
				}
			}
			fmt.Fprintln(w, result.String())
			result.Reset()
		}
	}
	return nil
}
