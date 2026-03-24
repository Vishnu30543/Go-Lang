package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is started\n", id)
	// Some work here
	fmt.Printf("Worker %d is done\n", id)
}

func main() {
	fmt.Println("Learning WaitGroup in Golang...")

	var wg sync.WaitGroup

	// Starting multiple worker goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all worker goroutines to finish
	fmt.Println("All workers are done. Exiting main function.")
}