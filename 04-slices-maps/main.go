// Lesson 4: slices, maps and range.
package main

import "fmt"

// sumSlice adds up every number in the slice.
func sumSlice(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// countWords counts how many times each word appears.
// A missing key reads as 0, so ++ works without checking first.
func countWords(words []string) map[string]int {
	counts := map[string]int{}
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(sumSlice(nums))

	words := []string{"go", "lua", "go", "ts", "go", "lua"}
	fmt.Println(countWords(words))
}
