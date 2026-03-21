package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name?")

	// Scan
	// var name string
	// fmt.Scanln(&name)
	// fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)

	// bufio
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Printf("Hello, %s! Welcome to Go programming.", name)
}