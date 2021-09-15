package piscine

func IsNumeric(s string) bool {
	for _, l := range s {
		if l < '0' || l > '9' {
			return false
		}
	}
	return true
}
