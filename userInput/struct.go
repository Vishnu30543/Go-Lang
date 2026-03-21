package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1st method
	// student := Person{Name: "Vishnu", Age: 20}

	// 2nd method
	var student Person
	student.Name = "Vishnu"
	student.Age = 20

	// 3rd method
	// student := new(Person)
	// student.Name = "Vishnu"
	// student.Age = 20

	fmt.Println("Student:", student)
	fmt.Printf("Type of student: %T\n", student)
}