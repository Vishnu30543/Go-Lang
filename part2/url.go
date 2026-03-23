package main

import (
	"fmt"
	"net/http"
)

func main() {
	myURL := "https://jsonplaceholder.typicode.com/posts/1"

	fmt.Printf("Type of URL: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer parsedURL.Body.Close()

	fmt.Printf("Type of parsed URL: %T\n", parsedURL)

	fmt.Println("Schema:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.RawQuery)

	// Modifying the URL
	parsedURL.Path = "/posts/2"

	newurl := parsedURL.String()
	fmt.Println("Modified URL:", newurl)

}