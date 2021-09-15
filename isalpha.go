package piscine

func IsAlpha(s string) bool {
	for _, l := range s {
		if l != ' ' && l < '0' || l > '9' && l < 'A' || l > 'Z' && l < 'a' || l > 'z' {
			return false
		}
	}
	return true
}
