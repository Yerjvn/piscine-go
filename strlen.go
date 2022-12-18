package piscine

func StrLen(s string) int {
	count := 0
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		count++
	}
	return count
}
