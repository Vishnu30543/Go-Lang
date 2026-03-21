package main

import "fmt"

func main() {
	// student <-> marks
	studentGrades := make(map[string]int)

	studentGrades["Vishnu"] = 100
	studentGrades["Alice"] = 85
	studentGrades["Bob"] = 90
	studentGrades["Charlie"] = 78

	fmt.Println("\nStudent Grades:", studentGrades)
	fmt.Printf("Type of studentGrades: %T\n", studentGrades)

	studentGrades["Alice"] = 95
	fmt.Println("\nGrade updated for Alice:", studentGrades["Alice"])
	fmt.Println("\nUpdated Student Grades:", studentGrades)

	for student, grade := range studentGrades {
		fmt.Printf("Student: %s, Grade: %d\n", student, grade)
	}

	//delete a key-value pair
	delete(studentGrades, "Alice")
	fmt.Println("\nStudent Grades after deleting Alice:", studentGrades)

	// Check if a key exists in the map
	if grade, exists := studentGrades["Alice"]; exists {
		fmt.Printf("\nAlice's grade exists: %d\n", grade)
	} else {
		fmt.Println("\nAlice's grade does not exist")
	}

	// Initialize a map with values
	countryCapitals := map[string]string{
		"India":    "New Delhi",
		"USA":      "Washington, D.C.",
		"France":   "Paris",
		"Japan":    "Tokyo",
	}
	fmt.Println("\nCountry Capitals:", countryCapitals)
}