package piscine

func Max(a []int) int {
	max := a[0]
	for _, val := range a {
		if max < val {
			max = val
		}
	}
	return max
}
