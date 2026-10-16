// Lesson 1: variables, basic types and formatted printing.
package main

import "fmt"

func main() {
	name := "Ilia"
	birthYear := 2005
	height := 1.60

	age := 2026 - birthYear

	fmt.Printf("My name is %s, I am %d years old and my height is %.2f meters.\n", name, age, height)
}
