package main

import "fmt"

func main() {
	fmt.Println("Error Handling in Go")

	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	ans, _ := divide(10, 2)

	fmt.Printf("Result of division: %d\n", ans)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}