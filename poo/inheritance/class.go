package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Rodo"
	ftEmployee.age = 32
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	//GetMessage(ftEmployee)
}
