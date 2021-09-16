package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fileName := os.Args[0][2:]

	for _, c := range fileName {
		z01.PrintRune(c)
	}

	z01.PrintRune('\n')

	// from the internet, not working:
	// the first argumen is indexed with 0 and is always the program name
	// myProgramName := os.Args[0]
	// this will display the program name
	// fmt.Println(myProgramName)
}
