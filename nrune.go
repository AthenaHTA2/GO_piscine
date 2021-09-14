package piscine

func NRune(s string, n int) rune {
	sRune := []rune(s)
	rlength := len(sRune)
	if n <= 0 { // First error trap
		return 0
	} else if n > rlength { // second error trap
		return 0
	} else {
		return sRune[n-1]
	}
}
