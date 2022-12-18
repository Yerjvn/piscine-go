package piscine

func Rot14(s string) string {
	str := []rune(s)
	for i := range str {
		if str[i] >= 'a' && str[i] <= 'l' || str[i] >= 'A' && str[i] <= 'L' {
			str[i] += 14
		} else if str[i] >= 'm' && str[i] <= 'z' || str[i] >= 'M' && str[i] <= 'Z' {
			str[i] -= 12
		}
	}
	return string(str)
}
