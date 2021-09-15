package piscine

func TrimAtoi(s string) int {
	rune := []rune(s)
	var inter []int
	neg := false
	for i, val := range s {
		if val == 45 {
			for j := 0; j < i; j++ {
				if rune[j] >= 48 && rune[j] <= 57 {
					neg = false
				} else {
					neg = true
				}
			}
		}
	}
	for _, val := range s {
		if val >= 48 && val <= 57 {
			inter = append(inter, int(val)-48)
		}
	}
	var num int
	for _, i := range inter {
		num = (num * 10) + i
	}
	if neg {
		num = (num * -1)
	}
	return num
}
