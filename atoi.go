package piscine

func Atoi(s string) int {
	result := 0
	sign := 1

	for _, v := range s {
		if int(v) == 45 {
			sign = -1
		} else if int(v) == 43 {
			sign = 1
		}
		if int(v) < 48 && int(v) != 45 && int(v) != 43 || int(v) > 57 {
			result = 0
			return result
		} else {
			if int(v) != 43 {
				number := int(v) - 48
				result = (result * 10) + number
				result = sign * result
			}
		}
	}
	return result
}
