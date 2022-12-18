package piscine

func TrimAtoi(s string) int {
	var array []int
	result := 0
	minIndx := -1
	firstIndx := 0
	index := 0
	count := 0
	for _, runs := range s {
		if runs == '-' {
			minIndx = index
		}
		if isDigit(runs) {
			if firstIndx == 0 {
				firstIndx = index
			}
			array = append(array, int(runs-'0'))
		}
		index++
	}

	for c := range array {
		count = c + 1
	}

	for i := 0; i < count; i++ {
		result = result*10 + array[i]
	}
	if minIndx < firstIndx && minIndx != -1 {
		result = result * -1
	}

	return result
}

func isDigit(digit rune) bool {
	for a := '0'; a <= '9'; a++ {
		if digit == a {
			return true
		}
	}
	return false
}
