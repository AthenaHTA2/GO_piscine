package piscine

func RecursiveFactorial(nb int) int {
	factorial := 1

	if nb < 0 || nb > 25 {
		factorial = 0
	}

	if nb == 0 {
		factorial = 1
	} else {
		factorial = nb * RecursiveFactorial(nb-1)
	}
	return factorial
}
