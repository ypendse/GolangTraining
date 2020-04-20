package main

import "fmt"

func printStudents(students map[int]string) {
	for roll, name := range students {
		fmt.Printf("[%d] %s\n", roll, name)
	}
}

func main() {
	fmt.Println("-------------------------------------------------------------------------")
	students := map[int]string{
		1: "ABC",
		2: "BCD",
		5: "XYZ",
	}

	fmt.Println("Student map")
	printStudents(students)
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("Changing name of roll no: 5")
	students[5] = "New name"
	fmt.Println("Student map after update")
	printStudents(students)
	fmt.Println("-------------------------------------------------------------------------")
	delete(students, 5)
	fmt.Println("Deleting student")
	fmt.Println("Student map after delete")
	printStudents(students)

}
