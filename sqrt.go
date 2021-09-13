package piscine

import (
	"fmt"
	"math"
)

func Sqrt(nb int) int {
	rehersal := math.Sqrt(float64(nb))
	modulo := rehersal - math.Round(rehersal)

	if modulo != 0 {
		return 0
	} else {
		fmt.Printf("rehersal is: %v\n", rehersal)
		fmt.Printf("modulo is: %v\n", modulo)
	}
	return int(rehersal)
}
