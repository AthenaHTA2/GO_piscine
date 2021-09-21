package piscine

func Map(f func(int) bool, a []int) []int {
	b := make([]bool, len(a))

	for i, v := range a {
		b[i] = f(v)
	}
	return b
}
