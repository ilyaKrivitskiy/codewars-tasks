package main

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	res := ""

	res += "(" +
		strconv.FormatUint(uint64(numbers[0]), 10) +
		strconv.FormatUint(uint64(numbers[1]), 10) +
		strconv.FormatUint(uint64(numbers[2]), 10) +
		") " +
		strconv.FormatUint(uint64(numbers[3]), 10) +
		strconv.FormatUint(uint64(numbers[4]), 10) +
		strconv.FormatUint(uint64(numbers[5]), 10) +
		"-" +
		strconv.FormatUint(uint64(numbers[6]), 10) +
		strconv.FormatUint(uint64(numbers[7]), 10) +
		strconv.FormatUint(uint64(numbers[8]), 10) +
		strconv.FormatUint(uint64(numbers[9]), 10)

	return res
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
