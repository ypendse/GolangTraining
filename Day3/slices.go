package main

import "fmt"

func updateFruits(fruits []string, newValue string, pos int) {
	fruits[pos] = newValue
}

func main() {
	fmt.Println("-------------------------------------------------------------------------")
	fruits := []string{"Apple", "Mango", "Kiwi", "Pineapple", "Coconut", "Watermelon"}
	fmt.Printf("Fruits slice: %s\n", fruits)
	fmt.Printf("Length of fruits slice: %d\n", len(fruits))
	updateFruits(fruits, "Grapes", 4)
	fmt.Printf("Fruits after update: %s\n", fruits)
	fmt.Println("-------------------------------------------------------------------------")

	newFruits := fruits
	fmt.Printf("newFruits before update: %s\n", newFruits)
	updateFruits(newFruits, "Gauva", 5)
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Printf("newFruits after update: %s\n", newFruits)
	fmt.Printf("Fruits after update: %s\n", fruits)
	fmt.Println("-------------------------------------------------------------------------")

	if len(newFruits) == len(fruits) {
		areSame := true
		for i := 0; i < len(fruits) && areSame; i++ {
			if !(fruits[i] == newFruits[i]) {
				areSame = false
			}
		}
		if areSame {
			fmt.Println("fruits and newFruits are same")
		} else {
			fmt.Println("fruits and newFruits are not same")
		}
	} else {
		fmt.Println("fruits and newFruits are of differnt size")
	}
}
