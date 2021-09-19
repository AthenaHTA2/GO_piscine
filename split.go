package piscine

func Split(s, sep string) []string {
	rSlice := []rune(s)
	rSep := []rune(sep)
	word := []rune()
	var newSlice []string
	var start rune
	var end rune
	result := ""
	for i := 0; i < len(rSlice); i++ {
		if rSlice[i:i+len(rSep)-1] == rSep {
			start = rSlice[i+len(rSep)]
		}
		for j := start; j < len(rSlice); j++ {
			if rSlice[j:j+len(rSep)-1] == rSep {
				end = rSlice[j-1]
			}
		}
		word = rSlice[start:end]
		if result != "" {
			newSlice = append(newSlice, string(word))
			result = ""
		} else {
			result = result + string(rSlice[i+len(sep)])
		}
		if i == len(rSlice)-1 {
			newSlice = append(newSlice, result)
		}
	}
	return newSlice
}
