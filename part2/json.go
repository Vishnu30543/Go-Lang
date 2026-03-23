package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Learning JSON in Golang...")

	person := Stu{ ID: 1, Name: "Alice", Age: 30}

	fmt.Printf("\nType of person struct: %T\n", person)
	fmt.Println("\nPerson struct:", person)

	// Convert struct to JSON Encoding [Marshalling]
	// Gives slice of bytes & error
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Printf("\nType of JSON data: %T\n", jsonData)
	fmt.Println("\nJSON Data:", string(jsonData))

	// Decoding JSON to struct [Unmarshalling]
	var persondata Stu
	err = json.Unmarshal(jsonData, &persondata)
	if err != nil {
		fmt.Println("Error converting from JSON:", err)
		return
	}

	fmt.Printf("\nType of persondata struct: %T\n", persondata)
	fmt.Println("\nPersondata struct:", persondata)

}