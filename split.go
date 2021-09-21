package piscine

func Split(s, sep string) []string {
	var word string
	var result []string

	for i := 1; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep[0:] {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		}
		if s[i:i+len(sep)] != sep[0:] {
			if i <= len(s)-len(sep) {
				word = word + s[i:i+len(sep)]
			} else if i <= len(s) {
				word = word + string(s[i])
			}
		}

		result = append(result, word)
	}
	return result
}
