package main

import (
	"fmt"
	"os"
)

func main() {

	// Create a new file
	file, err := os.Create("example.txt")

	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	fmt.Println("File created successfully")

	// Write some content to the file
	content := "Hello, this is a sample file created using Go!"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Println("Content written to file successfully")
}