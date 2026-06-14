package main

import (
	"fmt"
	"os"
	"strings"
)

func readBannerFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("error reading file: %q:\n %w", fileName, err)
	}
	return string(data), nil
}

func parseBannerFile(content string) (map[rune][]string, error) {

	for content == "" {
		return nil, fmt.Errorf("error: file is empty")
	}

	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")
	charMap := make(map[rune][]string)

	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines)%9 != 0 {
		return nil, fmt.Errorf("error; banner file is corrupt.")
	}

	for i := 0; i < len(lines)/9; i++ {
		char := rune(32 + i)
		start := i * 9
		charMap[char] = lines[start+1 : start+9]
	}
	return charMap, nil
}
