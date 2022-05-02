package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	res := ""
	arr := strings.Split(str, " ")

	for j, s := range arr {
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				res += strings.ToUpper(string(s[i]))
			} else {
				res += strings.ToLower(string(s[i]))
			}
		}
		if j != len(arr)-1 {
			res += " "
		}
	}

	return res
}

func main() {
	if strings.Compare(toWeirdCase("This is a test Looks like you passed"), "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd") == 0 {
		fmt.Println("True")
	}
}
