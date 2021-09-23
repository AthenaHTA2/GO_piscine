package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, v := range s {
		number := int(v) - 48
		result = (result * 10) + number
	}
	return result
}
