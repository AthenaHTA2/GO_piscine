package piscine

func Sqrt(nb int) int {
	var root int // Anthony suggestion: this variable can be dropped

	for i := 1; i <= nb; i++ {

		compare := i * i
		modulo := nb % i

		if compare == nb && modulo == 0 {
			root = i
			return root // Anthony: can say return i instead of return root
		} else { // Anthony explanation: this step is redundant, if make change in line 18
			root = 0
		}
	}
	return root // Change suggested by Anthony: root 0
}
