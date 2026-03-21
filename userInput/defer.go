package main

import "fmt"

func main() {
	fmt.Println("Top of the Program")

	// Defer statements are executed in LIFO (Last In, First Out) order
	defer fmt.Println("Deferred: This will be printed at the end of the main function")
	defer fmt.Println("Middle of the Program")
	
	fmt.Println("Bottom of the Program")
}