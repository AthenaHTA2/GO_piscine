package piscine

func TrimAtoi(s string) int {
	sS := []rune(s)
	intS := make([]int, 0)
	isNeg := 0
	multiplier := 1
	result := 0

	for i := 0; i < len(sS); i++ {
		if (sS[i] == 45 && isNeg == 0) && len(intS) == 0 {
			isNeg = 1
		} else if sS[i] >= 48 && sS[i] <= 57 {
			asciiValue := int(sS[i] - 48)
			intS = append(intS, asciiValue)
		}
	}
	for i := (len(intS) - 1); i >= 0; i-- {
		result += multiplier * intS[i]
		multiplier *= 10
	}
	if isNeg == 1 {
		result *= (-1)
	}
	return result
}
