package main

import "fmt"

func main() {
	fmt.Println("We are learning about functions in Go!")
	simpleFunction()

	result := add(5, 3)
	fmt.Printf("The result of adding 5 and 3 is: %d\n", result)
}



func simpleFunction() {
	fmt.Println("This is a simple function.")
}

func add(a, b int) int {
	return a + b
}