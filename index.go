package piscine

func Index(s string, toFind string) int {
	input1 := []rune(s)
	input2 := []rune(toFind)
	letterToFind := input2[0]
	length := len(input2)
	s_length := len(input1)

	if toFind == "" {
		return 0
	}
	if length <= 0 {
		return -1
	}

	for index, letter := range s {
		if letter == letterToFind {
			if length+index > s_length { // in this instance do nothing
			} else {
				if s[index:length+index] == toFind {
					return index
				}
			}
		}
	}
	return -1
}
