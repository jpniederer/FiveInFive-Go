package main

import (
	"fmt"
)

var x = 42
var y = 2

func main() {
	words := []string{"hi", "hello", "howdy"}
	sameWords := LogThenReturn(words)
	fmt.Println("Previous line should read:", sameWords)

	// Map
	person := make(map[string]int)
	person["Josh"] = 6
	person["Lauren"] = 24
	fmt.Println("The map value: ", person)

	// Build slice of numbers
	numbers := make([]int, 10)

	for i := range numbers {
		numbers[i] = i + 1
	}

	fmt.Println("My list of numbers: ", numbers)

	fmt.Println("10^10 =", Power(10, 10))
}

// Simple function
func LogThenReturn(ws []string) []string {
	fmt.Println(ws)
	return ws
}

// Recursion
func Power(x int, y int) int {
	return powerIter(x, y, 1)
}

// Package Scoped via lowercase.
func powerIter(x int, y int, total int) int {
	if y <= 0 {
		return total
	}

	return powerIter(x, y-1, total*x)
}
