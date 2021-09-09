package main

import "fmt"

func PointOne(n *int) {
	*n = 1
}

func main() {
	x := 0

	PointOne(&x)
	fmt.Println(x)
}
