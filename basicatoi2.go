package piscine

func BasicAtoi2(s string) int {
	result := 0

	for _, v := range s {
		if int(v) < 48 || int(v) > 57 {
			result = 0
			return result
		} else {
			number := int(v) - 48
			result = (result * 10) + number
		}
	}
	return result
}
