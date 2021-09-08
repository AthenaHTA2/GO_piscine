package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for x := '0'; x <= '9'; x++ {
		for y := '0'; y <= '9'; y++ {
			for z := '0'; z <= '9'; z++ {
				if x < y && y < z && x < z && x != y && y != z && x != '7' {
					z01.PrintRune(rune(x))
					z01.PrintRune(rune(y))
					z01.PrintRune(rune(z))
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else if x == '7' && y == '8' && z == '9' {
					z01.PrintRune(rune(x))
					z01.PrintRune(rune(y))
					z01.PrintRune(rune(z))
					z01.PrintRune('\n')
				}
			}
		}
	}
}
