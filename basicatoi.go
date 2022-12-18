package piscine

func BasicAtoi(s string) int {
	arrayStr := []rune(s)
	r := len(s)
	d := 0
	for i := 0; i < r; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			return d
		} else {
			d *= 10
			d += int(arrayStr[i]) - '0'
		}
	}
	return d
}
