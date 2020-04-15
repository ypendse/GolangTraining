package main

import "fmt"

func main() {
	lines := 5
	value := 1

	for i := 0; i < lines; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", value)
			value = value + 1
		}
		fmt.Print("\n")
	}
}
