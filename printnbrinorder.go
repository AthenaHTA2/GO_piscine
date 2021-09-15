package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	aRune := "0123456789"
	intS := make([]int, 0)

	for {
		mod := n % 10
		if n > 0 {
			intS = append(intS, mod)
		} else {
			intS = append(intS, (mod * (-1)))
		}

		n = n / 10
		if n == 0 {
			break
		}
	}
	SortIntegerTable2(intS)

	for i := 0; i < len(intS); i++ {
		z01.PrintRune(rune(aRune[intS[i]]))
	}
}

func SortIntegerTable2(table []int) {
	for i := 0; i < len(table); i++ {
		temp := 0
		for j := (i + 1); j < len(table); j++ {
			if table[i] <= table[j] {
				continue
			} else {
				temp = table[i]
				table[i] = table[j]
				table[j] = temp

			}
		}
	}
}
