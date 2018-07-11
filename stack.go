package main

import (
	"errors"
)

type stack []int

// func main() {
// 	var s stack
// 	s = s.push(5)
// 	s = s.push(4)
// 	s = s.push(3)
// 	s = s.push(2)
// 	s = s.push(1)
// 	s, _ = s.pop()
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// }

func (s stack) push(a int) stack {
	return append(s, a)
}

func (s stack) pop() (stack, error) {
	if len(s) == 0 {
		return s, errors.New("Empty Stack")
	}
	s = s[:len(s)-1]
	return s, nil
}
