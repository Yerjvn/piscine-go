package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	for index := range str {
		if index+1 == n {
			return str[index]
		}
	}
	return 0
}
