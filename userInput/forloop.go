package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	count := 0
	for{
		if count >= 5 {
			break
		}
		fmt.Println("Count:", count)
		count++
	}

	for count < 10 {
		fmt.Println("Count:", count)
		count++
	}

	// for range loop
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	data := "Hello, World!"
	for index, char := range data {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}