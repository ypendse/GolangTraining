package main

import (
	"testing"
)

func TestStudentInfo(t *testing.T) {
	s1 := Student{
		name: "lala",
		age:  18,
		ContactInfo: ContactInfo{
			phone: "1234", email: "s@sample.com",
		},
	}
	ans := s1.age
	if ans != 18 {
		t.Errorf("Student age = %d; want 18", ans)
	}

	ansEmail := s1.email
	if ansEmail == "" {
		t.Error("Student info is nil")
	}
}
