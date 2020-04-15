package main

import "fmt"

func fact(num int) int {
	value := 1
	for i := 1; i < num; i++ {
		value *= num
	}

	return value
}

func main() {
	number := 5
	fmt.Printf("Factorial of %d is %d\n", number, fact(number))
}
