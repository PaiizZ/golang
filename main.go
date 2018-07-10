package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("Hello")
// 	reverseArray()
// }

func reverseArray() {
	num := []int{5, 2, 6, 3, 1, 4}

	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}

	fmt.Println(num)
}
