package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	e := make([]int, max-min)
	for i := 0; i < (max - min); i++ {
		e[i] = min + i
	}
	return e
}
