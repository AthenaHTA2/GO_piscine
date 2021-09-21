package piscine

func forEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
