package piscine

func Sqrt(nb int) int {
	var root int

	for i := 1; i <= nb; i++ {

		compare := i * i
		modulo := compare % 1

		if compare == nb && modulo != 0 {
			root = i
		} else {
			root = 0
		}
	}
	return root
}
