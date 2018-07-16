package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	strs := make(map[string]int)
	for i := 0; i < len(strings.Fields(s)); i++ {
		if _, ok := strs[arr[i]]; ok {
			strs[arr[i]]++
		} else {
			strs[arr[i]] = 1
		}
	}
	return strs
}

func WordCountTon(s string) map[string]int {
	strs := make(map[string]int)
	for _, str := range strings.Fields(s) {
		strs[str]++
	}
	return strs
}

// func main() {
// 	fmt.Println(WordCountTon("I LOVE U"))
// 	fmt.Println(WordCount("I LOVE U"))
// }
