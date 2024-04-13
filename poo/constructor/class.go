package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Forma 1
	e := Employee{}
	fmt.Printf("%v\n", e)
	// Forma 2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	// Forma 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)

	// Forma 4
	e4 := NewEmployee(1, "Name", false)
	fmt.Printf("%v\n", *e4)
}
