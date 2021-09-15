package piscine

func IsPrintable(s string) bool {
	for _, c := range s {
		if c < ' ' {
			return false
		}
	}
	return true
}
