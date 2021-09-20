package main

import (
	"os"

	"github.com/01-edu/z01"
)

var (
	evenMsg string = "I have an even number of arguments"
	oddMsg  string = "I have an odd number of arguments"
	mod     int
	rString []rune
)

func printStr(s string) {
	rString = []rune(s)
	for _, r := range rString {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	mod = (len(os.Args) - 1) % 2
	if mod == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	if isEven(len(os.Args) - 1) {
		printStr(evenMsg)
	} else {
		printStr(oddMsg)
	}
}
