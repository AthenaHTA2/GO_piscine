package piscine

func UltimateDivMod(a *int, b *int) {
	var x int
	var y int
	x = *a / *b
	y = *a % *b
	*a = x
	*b = y
}
