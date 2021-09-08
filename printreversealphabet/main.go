package main

import "github.com/01-edu/z01"

func main() {
	zRune := 'z'
	for zRune >= 'a' {
		z01.PrintRune(zRune)
		zRune--
	}
	z01.PrintRune('\n')
}
