package piscine

func ToUpper(s string) string {
	Rune := []rune(s)
	for index, val := range Rune {
		for char := 'a'; char <= 'z'; char++ {
			if val == char {
				Rune[index] = val - 32
			}
		}
	}
	return string(Rune)
}
