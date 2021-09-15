package piscine

func ToLower(s string) string {
	Rune := []rune(s)
	for index, val := range Rune {
		for char := 'A'; char <= 'Z'; char++ {
			if val == char {
				Rune[index] = val + 32
			}
		}
	}
	return string(Rune)
}
