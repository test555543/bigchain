package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	num1 := 10
	num2 := 20
	result := add(num1, num2)
	fmt.Println("Result is:", result)
}
