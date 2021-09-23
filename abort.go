package piscine

func Abort(a, b, c, d, e int) int {
	intSlice := []int{a, b, c, d, e}

	for i := 1; i < len(intSlice); i++ {
		temp := 0
		for j := i + 1; j < len(intSlice); j++ {
			if intSlice[i] <= intSlice[j] {
				continue
			} else {
				temp = intSlice[i]
				intSlice[i] = intSlice[j]
				intSlice[j] = temp
			}
		}
	}

	return intSlice[2]
}
