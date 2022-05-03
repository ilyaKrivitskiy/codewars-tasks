package main

import (
	"fmt"
)

func Race(v1, v2, g int) [3]int {
	res := [3]int{}

	if v1 > v2 {
		return [3]int{-1, -1, -1}
	}

	t := float32(g) / (float32(v2) - float32(v1))
	res[0] = g / (v2 - v1)
	t -= float32(res[0])
	t *= 60
	min := int(t)
	res[1] = min
	t -= float32(min)
	res[2] = int(t * 60)

	return res
}

func main() {
	fmt.Println(Race(820, 81, 550))

}
