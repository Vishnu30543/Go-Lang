package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8)

	fmt.Println("Numbers:", numbers)
	fmt.Printf("Type of numbers: %T\n", numbers)
	fmt.Println("Slice length:", len(numbers))

	// Create a new slice with make
	fmt.Println("\nCreating a new slice with make:")

	num := make([]int, 0, 4)
	fmt.Println("New slice:", num)
	fmt.Printf("Type of num: %T\n", num)
	fmt.Println("Slice length:", len(num))
	fmt.Println("Slice capacity:", cap(num))

	// Append to the new slice
	num = append(num, 10, 20, 30, 40, 50)
	fmt.Println("Updated slice after appending:", num)
	fmt.Println("Slice length after appending:", len(num))
	fmt.Println("Slice capacity after appending:", cap(num))

	// 2 easy ways to create a slice
	slice1 := []string{}
	slice2 := make([]string, 0)		// capacity is optional, default is 0
	
	
}
