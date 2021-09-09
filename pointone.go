package piscine

import "github.com/01-edu/z01"

func PointOne(n *int) {
	*n = 1
}

func main() {
	x := 0
	PointOne(&x)
	z01.PrintRune('x')
}
