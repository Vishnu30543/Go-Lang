package	main

import "fmt"
import "strconv"

func main() {
	num := 123
	str := strconv.Itoa(num)
	fmt.Printf("Type: %T\n", str)
	fmt.Printf("String: %s\n", str)

	numberStr := "456"
	number, _ := strconv.Atoi(numberStr)
	fmt.Printf("\nType: %T\n", number)
	fmt.Printf("Number: %d\n", number)

	numString := "3.14"
	digit, _ := strconv.ParseFloat(numString, 64)
	fmt.Printf("\nType: %T\n", digit)
	fmt.Printf("Number: %f\n", digit)

}