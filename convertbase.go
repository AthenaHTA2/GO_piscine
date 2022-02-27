package piscine

import (
	"strconv"
)

func ConvertBase(nbr, convertFrom, convertTo string) string {

	var rNbr []rune = []rune(nbr)
	var multiply int = 1
	var power int = 0
	step := 0
	calc := 0
	var base int

	mBin := make(map[rune]int)
	mBin['0'] = 0
	mBin['1'] = 1

	mHex := make(map[rune]int)
	mHex['0'] = 0
	mHex['1'] = 1
	mHex['2'] = 2
	mHex['3'] = 3
	mHex['4'] = 4
	mHex['5'] = 5
	mHex['6'] = 6
	mHex['7'] = 7
	mHex['8'] = 8
	mHex['9'] = 9
	mHex['A'] = 10
	mHex['B'] = 11
	mHex['C'] = 12
	mHex['D'] = 13
	mHex['E'] = 14
	mHex['F'] = 15

	mDec := make(map[rune]int)
	mDec['0'] = 0
	mDec['1'] = 1
	mDec['2'] = 2
	mDec['3'] = 3
	mDec['4'] = 4
	mDec['5'] = 5
	mDec['6'] = 6
	mDec['7'] = 7
	mDec['8'] = 8
	mDec['9'] = 9

	if "nbr" == "" || "convertFrom" == "" || "convertTo" == "" {
	}
	if convertFrom == "01" {
		base = 2
	} else if convertFrom == "16" {
		base = 16
	} else {
		base = 10
	}
	switch base {
	case 2:
		for i := len(rNbr) - 1; i >= 0; i-- {
			if power == 0 {
				multiply = 1
				step = mBin[rNbr[i]] * multiply
				calc = calc + step
				power = power + 1
			} else {
				multiply = multiply * base
				step = mBin[rNbr[i]] * multiply
				calc = calc + step
				power = power + 1
			}
		}
	case 16:
		for i := len(rNbr) - 1; i >= 0; i-- {
			if power == 0 {
				multiply = 1
				step = mHex[rNbr[i]] * multiply
				calc = calc + step
				power = power + 1
			} else {
				multiply = multiply * base
				step = mHex[rNbr[i]] * multiply
				calc = calc + step
				power = power + 1
			}
		}
	case 10:
		for i := len(rNbr) - 1; i >= 0; i-- {
			if power == 0 {
				multiply = 1
				step = mDec[rNbr[i]] * multiply
				calc = calc + step
				power = power + 1
			} else {
				multiply = multiply * base
				step = mDec[rNbr[i]] * multiply
				calc = calc + step
				power = power + 1
			}
		}
		// fmt.Printf("convertFrom %T \nconvertFrom %v --> \n", convertFrom, convertFrom)
		// fmt.Print("base --> ", base, "\n")
		// fmt.Print("multiply --> ", multiply, "\n")
		// fmt.Print("number length --> ", len(rNbr), "\n")
		// fmt.Print("rune number --> ", string(rNbr), "\n")
	}
	return strconv.Itoa(calc)
}
