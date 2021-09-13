package piscine

func IterativeFactorial(nb int) int {
	factorial := 1

	if nb < 0 || nb > 25 {
		factorial = 0
	} else if nb == 0 {
		factorial = 1
	} else {
		for i := 1; i <= nb; i++ {
			factorial = factorial * i
		}
	}
	return factorial
}
