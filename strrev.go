package piscine

func StrRev(a string) string {
	RunA := []rune(a)
	for i, j := 0, len(RunA)-1; i < j; i, j = i+1, j-1 {
		RunA[i], RunA[j] = RunA[j], RunA[i]
	}
	return string(RunA)
}

// func StrRev(a string) string{ //my old code
//	length := StrLen(a)
//	RunA := []rune(a)
//	for _, letter := range RunA {
//		return string(RunA[length])
//	}
//	length = length -1
