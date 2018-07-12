package main

import (
	"fmt"
	"strings"
)

func main() {
	wc := NewWordCount()
	wc = wc.AddWord("I a a a")
	fmt.Println(wc.MaxWordCount())
	wc.PrintAll()
}

type WordCount struct {
	Count  int
	Length int
}

type WordsCount map[string]WordCount

func NewWordCount() WordsCount {
	return make(WordsCount)
}

func (wc WordsCount) AddWord(s string) WordsCount {
	for _, str := range strings.Fields(s) {
		wc[str] = WordCount{
			Count:  wc[str].Count + 1,
			Length: len(str),
		}
	}
	return wc
}

func (wc WordsCount) MaxWordCount() (string, int) {
	var key string
	var maxCount int
	for value := range wc {
		if wc[value].Count > maxCount || maxCount == 0 {
			key = value
			maxCount = wc[value].Count
		}
	}
	return key, maxCount
}

func (wc WordsCount) MinWordCount() (string, int) {
	return "", 0
}

func (wc WordsCount) LongestWord() (string, int) {
	return "", 0
}

func (wc WordsCount) ShortestWord() (string, int) {
	return "", 0
}

func (wc WordsCount) PrintAll() {
	for key, value := range wc {
		fmt.Println("Key:", key, "Value:", value)
	}
}
