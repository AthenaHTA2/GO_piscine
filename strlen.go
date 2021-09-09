package piscine

func StrLen(s string) int {
	RunesInStr := []rune(s)
	length := 0
	for range RunesInStr {
		length = length + 1
	}
	return length
}
