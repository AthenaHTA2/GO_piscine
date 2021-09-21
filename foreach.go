package piscine

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

//Luis's code:
//func ForEach(f func(int), a []int) {
//	for i := 0; i < len(a); i++ {
//		f(a[i])
//	}
//}
