package piscine

func AlphaCount(s string) int {
	str := []rune(s)
	counter := 0
	for i := 0; i < len(str); i++ {
		if (str[i] >= 65 && str[i] <= 90) || (str[i] >= 97 && str[i] <= 122) {
			counter++
		}
	}
	return counter
}
