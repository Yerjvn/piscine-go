package piscine

func IsLower(s string) bool {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] < 97 || str[i] > 122 {
			return false
		}
	}
	return true
}
