package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 {
		return 0
	} else if len(toFind) == 0 {
		return 0
	}
	r := 0
	for i := range s {
		if s[i] == toFind[r] {
			if r == len(toFind)-1 {
				return i - (len(toFind) - 1)
			}
			r++
		}
	}
	return -1
}
