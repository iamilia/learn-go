package main

import (
	"fmt"
	// "strconv"
)

func sumSlice(numbers []int) int {
	numb := 0
	for _, n := range numbers {
		numb = numb + n
	}
	return numb
}

func countWords(words []string) map[string]int {
	wordCounter := map[string]int{}
	for _, word := range words {
		wordCounter[word]++
	}
	return wordCounter
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(sumSlice(nums))
	words := []string{"go", "lua", "go", "ts", "go", "lua"}
	fmt.Println(countWords(words))
}
