package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	MutableString := []rune(s)
	for _, letter := range MutableString {
		z01.PrintRune(letter)
	}
}
