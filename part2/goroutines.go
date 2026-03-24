package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Begin of Hello!")
	time.Sleep(500 * time.Millisecond) // Simulating some work with sleep
	fmt.Println("End of Hello!")
}

func sayHi() {
	fmt.Println("Begin of Hi!")
	time.Sleep(500 * time.Millisecond) // Simulating some work with sleep
	fmt.Println("End of Hi!")
}

func main() {
	fmt.Println("Learning Goroutines in Golang...")
	go sayHello() // This will run concurrently with the main function
	go sayHi()    // This will also run concurrently with the main function

	fmt.Println("This is the main function.")
	// Adding a sleep to allow the goroutine to finish before the program exits
	time.Sleep(1 * time.Second)


}