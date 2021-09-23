package piscine

func IsPrime(nb int) bool {
	prime := false
	if nb <= 1 {
		prime = false
	} else if nb%2 == 0 {
		prime = false
	} else {
		prime = true
	}
	return prime
}
