package piscine

func SplitWhiteSpaces(s string) []string {
	sSlice := []rune(s)
	var newSlice []string
	result := ""
	for i := 0; i < len(sSlice); i++ {
		if sSlice[i] == '\t' || sSlice[i] == '\n' || sSlice[i] == ' ' {
			if result != "" {
				newSlice = append(newSlice, result)
				result = ""
			}
		} else {
			result = result + string(sSlice[i])
		}
		if i == len(sSlice)-1 {
			newSlice = append(newSlice, result)
		}
	}
	return newSlice
}

// My solution before asking Stephen for help

// package piscine

//func SplitWhiteSpaces(s string)[]string {
//	countWords := 1
//	sSlice := []rune(s)
//	var newSlice []string
//
//	for i,c := range sSlice {
//
//		if c == "\t" || c == "\n" || c == " " {
//			countWords = countWords + 1
//		}
//
//		for i,c := range sSlice {
//		if c != "\t" || c != "\n" || c != " " {
//			newSlice = append(newSlice, c)
//
//		}
//	}
//}
//}
