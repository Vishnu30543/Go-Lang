package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// variable declarations
	var name string = "Vishnu"
	var age int = 30
	var isStudent bool = false

	// printing variable values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)

	// using short variable declaration
	var city = "Bangalore"
	var country = "India"

	fmt.Println("City:", city)
	fmt.Println("Country:", country)

	// constants
	const pi = 3.14
	const gravity = 9.8

	// public and private variables
	// In Go, identifiers that start with an uppercase letter are exported (public)
	// while those that start with a lowercase letter are unexported (private).
	var PublicVariable = "I am public"
	var privateVariable = "I am private"

	fmt.Println("Public Variable:", PublicVariable)
	fmt.Println("Private Variable:", privateVariable)

	// public and private functions calls
	fmt.Print("Public Function Output: ")
	PublicFunction()
	fmt.Println("Private Function Output:", privateFunction())

	// scenario
	name = "Alice"
	age = 25
	height := 5.6562

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
	fmt.Printf("Type of name: %T\n, Age: %T, Height: %T\n", name, age, height)
}



func PublicFunction() {
	fmt.Println("This is a public function")
}

func privateFunction() string {
	return "This is a private function"
}