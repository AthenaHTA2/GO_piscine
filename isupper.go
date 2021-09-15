package piscine

func IsUpper(s string) bool {
	for _, l := range s {
		if l < 'A' || l > 'Z' {
			return false
		}
	}
	return true
}

// Returns True True, as does not recognise characters not in scope e.g. "!"
//	for _, b := range Test {
//	if b >= 65 && b <= 90 {
//		return true
//	}
//}
//return false
