package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	// Open the file for reading
	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")

	// Read the file content using a function
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println("File content:\n", string(content))

}