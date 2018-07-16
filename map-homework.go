package main

// func main() {
// 	wc := NewWordCount()
// 	wc = wc.AddWord("Oh wait til’ I do what I do Hit you with that ddu-du ddu-du du Hit you with that ddu-du ddu-du du")
// 	fmt.Println(wc.MaxWordCount())
// 	fmt.Println(wc.MinWordCount())
// 	fmt.Println(wc.LongestWord())
// 	fmt.Println(wc.ShortestWord())
// 	// wc.PrintAll()
// }

// type WordCount struct {
// 	Count  int
// 	Length int
// }

// type WordsCount map[string]WordCount

// func NewWordCount() WordsCount {
// 	return make(WordsCount)
// }

// func (wc WordsCount) AddWord(s string) WordsCount {
// 	for _, str := range strings.Fields(s) {
// 		wc[str] = WordCount{
// 			Count:  wc[str].Count + 1,
// 			Length: len(str),
// 		}
// 	}
// 	return wc
// }

// func (wc WordsCount) MaxWordCount() (string, int) {
// 	var key string
// 	var maxCount int
// 	for value := range wc {
// 		if wc[value].Count > maxCount || maxCount == 0 {
// 			key = value
// 			maxCount = wc[value].Count
// 		}
// 	}
// 	return key, maxCount
// }

// func (wc WordsCount) MinWordCount() (string, int) {
// 	var key string
// 	var minCount int
// 	for value := range wc {
// 		if wc[value].Count < minCount || minCount == 0 {
// 			key = value
// 			minCount = wc[value].Count
// 		}
// 	}
// 	return key, minCount
// }

// func (wc WordsCount) LongestWord() (string, int) {
// 	var key string
// 	var longest int
// 	for value := range wc {
// 		if wc[value].Length > longest || longest == 0 {
// 			key = value
// 			longest = wc[value].Length
// 		}
// 	}
// 	return key, longest
// }

// func (wc WordsCount) ShortestWord() (string, int) {
// 	var key string
// 	var shortest int
// 	for value := range wc {
// 		if wc[value].Length < shortest || shortest == 0 {
// 			key = value
// 			shortest = wc[value].Length
// 		}
// 	}
// 	return key, shortest
// }

// func (wc WordsCount) PrintAll() {
// 	for key, value := range wc {
// 		fmt.Println("Key:", key, "Value:", value)
// 	}
// }
