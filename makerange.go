package piscine

func MakeRange(min, max int) []int {
	size := max - min
	betweenMinAndMax := make([]int, 0)

	if max <= min {
		negative := make([]int, 0)
		betweenMinAndMax = negative

	} else {
		positive := make([]int, size)
		betweenMinAndMax = positive

		for i := 0; i < size; i++ {
			betweenMinAndMax[i] = i + (size)
		}
	}
	return betweenMinAndMax
}
