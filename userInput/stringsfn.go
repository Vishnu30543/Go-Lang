package main

import "fmt"
import "strings"

func main() {
	data := "apple,banana,orange"
	parts := strings.Split(data, ",")
	fmt.Printf("Parts: %v\n", parts)

	str := "one two three two two one"
	count := strings.Count(str, "two")
	fmt.Printf("Occurrences of 'two': %d\n", count)

	str = "          Hello, World!     "
	trimmed := strings.TrimSpace(str)
	fmt.Printf("Trimmed string: '%s'\n", trimmed)
}