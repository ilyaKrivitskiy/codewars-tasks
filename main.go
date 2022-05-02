package main

import (
	"fmt"
)

func validateLetter(a string) int {
	switch string(a) {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func canBeDiff(a, b string) bool {
	if (a == "I" && b == "V") || (a == "I" && b == "X") || (a == "X" && b == "L") || (a == "X" && b == "C") ||
		(a == "C" && b == "D") || (a == "C" && b == "M") {
		return true
	} else {
		return false
	}

}

func Decode(roman string) int {
	strlen := len(roman)

	cur := validateLetter(string(roman[0]))

	if strlen == 1 {
		return cur
	}

	for i := 0; i < strlen-1; i++ {
		a, b := validateLetter(string(roman[i])), validateLetter(string(roman[i+1]))

		if a >= b {
			cur += b
		} else {
			if canBeDiff(string(roman[i]), string(roman[i+1])) {
				cur = b - 2*a + cur
			} else {
				cur += b
			}
		}
	}

	return cur
}

func main() {
	fmt.Println(Decode("M"))
}
