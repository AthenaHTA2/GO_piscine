package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb < 0 {
		return false
	}
	for i := 1; i < 20; i++ {
		if nb == i*i {
			return false
		}
	}
	if nb%2 != 0 || nb == 2 {
		return true
	}
	return false
}
