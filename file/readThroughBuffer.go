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

	// Read the file content using a buffer
	buffer := make([]byte, 1024) // Create a buffer of size 1024 bytes

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		fmt.Print(string(buffer[:n]))
	}

}