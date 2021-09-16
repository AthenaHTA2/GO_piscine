package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// command line arguments all stored in cmdArgs
	cmdArgs := os.Args

	argsNum := len(cmdArgs) - 1

	for i := argsNum; i >= 1; i-- {
		argumentX := os.Args[i]

		for _, c := range argumentX {
			z01.PrintRune(rune(c))
		}
		z01.PrintRune('\n')
	}
}
