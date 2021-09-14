package piscine

// import 		"github.com/01-edu/z01"

func AlphaCount(s string) int {
	//	alphabet := []byte{
	//		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	//	}
	// input := []string{"s"}

	var numberOfLetters int = 0
	// var inputLength int = len(s)

	for _, letter := range s {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			numberOfLetters++
		}
	}
	return numberOfLetters
}
