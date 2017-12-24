package main

import "fmt"

const fullName = "Joshua"

func main() {
	// Short variable declaration.
	name := "Josh"
	i, j := 0, 10
	fmt.Println("\nOriginal values: i =", i, " j =", j)
	i, j = j, i
	fmt.Println("\nAfter the swap: i =", i, " j =", j)

	fmt.Println("\nYou can call me", name)
	pointerUpdate(&name, fullName)
	fmt.Println("\nMy full name is", name)
}

func pointerUpdate(name *string, newName string) string {
	*name = newName
	fmt.Println("\nUpdating the name to", newName)
	return *name
}
