package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var rSlice []rune
	var line []rune

	for i := 0; i < len(a); i++ {
		if i < len(a) {
			rSlice = []rune(a[i])
			line = append(rSlice, '\n')
			for i := range line {
				z01.PrintRune(line[i])
			}
		} else {
			rSlice = []rune(a[i])
			line = rSlice
			for _, c := range line {
				z01.PrintRune(line[c])
			}
		}
	}
}
