package main

import "fmt"

func main() {
	var arr [5]int

	fmt.Println("Array:", arr)
	// Prints [0 0 0 0 0]

	var names [5]string
	
	fmt.Println("Names:", names)
	// Prints [     ] (5 empty strings for 5 elements)
	fmt.Printf("Names with quotes (quoted strings): %q\n", names)
	// Prints ["" "" "" "" ""]


	fmt.Printf("Names with spaces (space-separated strings): %s\n", names)
	// Prints [     ] (5 spaces for 5 empty strings)
	fmt.Printf("Names with default formatting: %v\n", names)
	// Prints [     ] (5 spaces for 5 empty strings)
	fmt.Printf("Names with Go-syntax representation: %#v\n", names)
	// Prints ["" "" "" "" ""] (5 empty strings without spaces)
	fmt.Printf("Names with quoted Go-syntax representation: %#q\n", names)
	// Prints [`` `` `` `` ``] (5 empty strings without spaces)
}