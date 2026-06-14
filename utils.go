package main

func isValidArgsCount(args []string) bool {
	return len(args) >= 1 && len(args) <= 2
}

func isValidChar(r rune) bool {
	return r >= 32 && r <= 126
}
