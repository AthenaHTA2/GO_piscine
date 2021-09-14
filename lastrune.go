package piscine

func LastRune(s string) rune {
	sRune := []rune(s)
	lastR := len(sRune)
	return sRune[lastR-1]
}
