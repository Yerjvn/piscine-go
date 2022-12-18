package piscine

func LastRune(s string) rune {
	str := []rune(s)
	ln := len(str) - 1

	return str[ln]
}
