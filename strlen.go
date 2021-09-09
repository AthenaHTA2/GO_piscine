package piscine

func strLen(s string) int {
	RunesInStr := []rune(s)
	lenght := '0'
	for _, letter := range RunesInStr {
		lenght = lenght + letter
	}
	return int(lenght)
}
