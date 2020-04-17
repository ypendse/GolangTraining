package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ContactInfo ...
type ContactInfo struct {
	phone string
	email string
}

// Student ...
type Student struct {
	name string
	age  int
	ContactInfo
}

func (s Student) print() {
	fmt.Printf("%#v\n", s)
}

func (s *Student) updateAge(newAge int) {
	s.age = newAge
}

func (s Student) saveToFile(filename string) {
	lines := fmt.Sprintf("%#v\n", s)
	err := ioutil.WriteFile(filename, []byte(lines), 0666)
	if err != nil {
		os.Exit(1)
	}

}

func main() {
	s1 := Student{
		name: "lala",
		age:  18,
		ContactInfo: ContactInfo{
			phone: "1234", email: "s@sample.com",
		},
	}

	fmt.Println("--------Before update---------")
	s1.print()
	s1.updateAge(20)
	fmt.Println("--------After update---------")
	s1.print()
	fmt.Println("-----------------------------")
	fmt.Println("--------Saving to file-------")
	s1.saveToFile("./s1_data.txt")
	fmt.Println("-----------------------------")

}
