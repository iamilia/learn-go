// Lesson 3: functions and multiple return values.
package main

import (
	"fmt"
	"strconv"
)

// fizzBuzz returns the FizzBuzz word for n instead of printing it.
func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

// minMax returns the smallest and the largest of three numbers.
func minMax(a, b, c int) (lo, hi int) {
	lo, hi = a, a
	if b < lo {
		lo = b
	}
	if c < lo {
		lo = c
	}
	if b > hi {
		hi = b
	}
	if c > hi {
		hi = c
	}
	return lo, hi
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(fizzBuzz(i))
	}

	lo, hi := minMax(3, 7, 5)
	fmt.Println("Min and Max of 3, 7, 5:", lo, hi)
}
