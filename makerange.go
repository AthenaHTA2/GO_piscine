package piscine

func MakeRange(min, max int) []int {
	size := max - min
	var result []int

	if max <= min {
		return result
	} else {
		newSlice := make([]int, size)
		for i := 0; i < size; i++ {
			newSlice[i] = min + i
		}
		return newSlice
	}
}
