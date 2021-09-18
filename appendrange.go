package piscine

func AppendRange(min, max int) []int {
	var betweenMinAndMax []int

	if min >= max {
		return betweenMinAndMax
	} else {
		for i := min; i < max; i++ {
			betweenMinAndMax = append(betweenMinAndMax, i)
		}
	}
	return betweenMinAndMax
}
