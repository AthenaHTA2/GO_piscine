package piscine

func CollatzCountdown(start int) int {
	counter := 0
	for start != 1 {
		if start <= 0 {
			return -1
		} else if start%2 == 0 {
			start = start / 2
			counter++
		} else {
			start = 3*start + 1
			counter++
		}
	}
	return counter
}
