package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	count := 5
	temp := 0

	for i := 0; i < count-1; i++ {
		temp = i
		for j := i + 1; j < count; j++ {
			if arr[j] < arr[temp] {
				temp = j
			}
		}
		arr[i], arr[temp] = arr[temp], arr[i]
	}
	return arr[2]
}
