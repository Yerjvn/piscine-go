package piscine

func BasicAtoi2(s string) int {
	ar := []rune(s)
	r := len(s)
	d := 0
	for i := 0; i < r; i++ {
		if ar[i] < '0' || ar[i] > '9' {
			return 0
		} else {
			d *= 10
			d += int(ar[i]) - '0'
		}
	}
	return d
}
