package piscine

func Abort(a, b, c, d, e int) int {
	intSlice := [5]int{a, b, c, d, e}
	var store int

	for i := 0; i <= 3; i++ {
		if intSlice[i] > intSlice[i+1] {
			store = intSlice[i+1]
			intSlice[i+1] = intSlice[i]
			intSlice[i] = store
		}
	}
	return intSlice[2]
}
