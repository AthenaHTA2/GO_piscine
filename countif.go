package piscine

func CountIf(f func(string) bool, tab []string) int {
	a := 0

	for i := range tab {
		f(tab[i])
		if f(tab[i]) {
			a++
		}
	}
	return a
}
