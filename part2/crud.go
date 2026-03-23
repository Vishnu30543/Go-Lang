package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", res.StatusCode)
		return
	}

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Retrieved Todo:", todo)
}

func performPostRequest() {
	todo := Todo{ UserID: 1, Title: "Learn Golang", Completed: false }

	// Convert struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// convert jsonData to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/todos"

	// Send POST request
	res, err := http.Post(myurl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response from POST request:", string(data))

	fmt.Println("POST request status:", res.Status)
}

func performUpdateRequest() {
	// This is a placeholder for the update request function

	todo := Todo{ UserID: 1, Title: "Learn Golang - Updated", Completed: true }

	// Convert struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// convert jsonData to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	// Create UPDATE/PUT request

	req , err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send PUT request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response from PUT request:", string(data))
	fmt.Println("PUT request status:", res.Status)
}

func performDeleteRequest() {
	// This is a placeholder for the delete request function

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making DELETE request:", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("DELETE request status:", res.Status)

}

func main() {
	fmt.Println("Learning CRUD operations in Golang...")

	// performGetRequest()

	// performPostRequest()

	// performUpdateRequest()

	 performDeleteRequest()
}
