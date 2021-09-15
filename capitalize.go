package piscine

func Capitalize(s string) string {
	Rune := []rune(s)

	for i, v := range Rune {
		if v >= 'A' && v <= 'Z' {
			Rune[i] = v + 32 // turn all into lower case
		}
	}

	for i, v := range Rune {
		if v >= 'a' && v <= 'z' {
			if i-1 < 0 {
				Rune[i] = v - 32 // Capitalize the very first character in the sentence
			} else if Rune[i-1] >= 'a' && Rune[i-1] <= 'z' || Rune[i-1] >= 'A' && Rune[i-1] <= 'Z' || Rune[i-1] >= '0' && Rune[i-1] <= '9' {
			} else {
				Rune[i] = v - 32 // capitalize if previous character was not lower case alphanumeric
			}
		}
	}
	return string(Rune)
}
