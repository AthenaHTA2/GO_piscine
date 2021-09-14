package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
	//	aRune := []rune(a)
	//	bRune := []rune(b)
	//	same := 0
	//	twoSame := 1
	//	fourSame := -1
	//	result := 0
	//	counter := 0
	//
	//	if len(aRune) == len(bRune) {
	//		for i := 0; i < len(aRune); i++ {
	//			if aRune[i] == bRune[i] {
	//				counter = counter + 1
	//			}
	//		}
	//		if counter == len(aRune) {
	//			result = same
	//		} else if counter == 4 {
	//			result = fourSame
	//		} else if counter == 2 {
	///			result = twoSame
	//	}
	//
	//	}
	//	return result
}
