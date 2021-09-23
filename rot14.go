package piscine

func Rot14(s string) string {
	// rString := []rune(s)
	var rot14String string

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			if r+14 > 'z' {
				rot14String = rot14String + string(r-12)
			} else {
				rot14String = rot14String + string(r+14)
			}
		} else if r >= 'A' && r <= 'Z' {
			if r+14 > 'Z' {
				rot14String = rot14String + string(r-12)
			} else {
				rot14String = rot14String + string(r+14)
			}
		} else {
			rot14String = rot14String + string(r)
		}
	}
	return rot14String
}
