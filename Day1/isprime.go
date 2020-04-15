package main

import "fmt"

func isprime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	number := 4
	fmt.Printf("Is %d a prime number: %t\n", number, isprime(number))
}
