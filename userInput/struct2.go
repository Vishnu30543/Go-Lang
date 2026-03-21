package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	City    string
	Country string
}

type Employee struct {
	Person
	Contact
	Address
	Position string
}

func main() {
	// 1st method
	var emp1 Employee
	emp1.Name = "Alice"
	emp1.Age = 30
	emp1.Email = "alice@example.com"
	emp1.Phone = "123-456-7890"
	emp1.City = "New York"
	emp1.Country = "USA"
	emp1.Position = "Software Engineer"
	fmt.Printf("Employee: %+v\n", emp1)

	// 2nd method
	emp2 := Employee{
		Person: Person{
			Name: "Bob",
			Age:  25,
		},
		Contact: Contact{
			Email: "bob@example.com",
			Phone: "098-765-4321",
		},
		Address: Address{
			City:    "Los Angeles",
			Country: "USA",
		},
		Position: "Product Manager",
	}
	fmt.Printf("\nEmployee: %+v\n", emp2)

	// 3rd method
	var emp3 = new(Employee)
	emp3.Name = "Charlie"
	emp3.Age = 28
	emp3.Email = "charlie@example.com"
	emp3.Phone = "555-555-5555"
	emp3.City = "Chicago"
	emp3.Country = "USA"
	emp3.Position = "Designer"
	fmt.Printf("\nEmployee: %+v\n", emp3)

}