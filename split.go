package piscine

func Split(s, sep string) []string {
	var word string
	var result []string
	var strLength int = len([]rune(s)) - 1
	var sepLength int = len([]rune(sep))

	for i := 0; i <= strLength-sepLength; i++ {
		if s[i:i+sepLength] == sep[0:] {
			if word != "" {
				result = append(result, word)
				word = ""
				i += sepLength - 1
			}
		} else {

			word = word + string(s[i])
		}
		if i == strLength-sepLength {
			word = word + string(s[i+1:])
			result = append(result, word)
		}
	}
	return result
}
