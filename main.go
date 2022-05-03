package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {

	arr := strings.Split(ip, ".")

	if len(arr) != 4 {
		return false
	}

	for _, s := range arr {
		dig, err := strconv.Atoi(s)
		if err != nil || dig < 0 || dig > 255 || (s[0] == '0' && len(s) != 1) {
			return false
		}

	}

	return true
}

func main() {
	fmt.Println(Is_valid_ip("123.045.067.089"))
}
