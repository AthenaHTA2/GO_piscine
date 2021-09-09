package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	MutableString := []rune(s)
	for _, letter := range MutableString {
		z01.PrintRune(letter)
	}
}

func main() {
	piscine.PrintStr("s")
}
