package main

import "fmt"

func MoveZeros(arr []int) []int {
	c := 0

	mas := make([]int, 0, cap(arr))

	for i, x := range arr {
		if x == 0 {
			c++
			continue
		}
		mas = append(mas, arr[i])
	}
	if c != 0 {
		for i := 0; i < c; i++ {
			mas = append(mas, 0)
		}
	}

	return mas
}

func main() {
	arr0 := []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}
	arr := []int{5, 0, 6, 7, 8, 1, 1, 3, 1, 9, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(MoveZeros(arr0))
	fmt.Println(MoveZeros(arr))
}
