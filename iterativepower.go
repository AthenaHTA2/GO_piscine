package piscine

func IterativePower(nb int, power int) int {
	iPower := 1

	if power < 0 {
		iPower = 0
	} else {
		for i := 1; i <= power; i++ {
			iPower = nb * IterativePower(nb, power-1)
		}
	}
	return iPower
}
