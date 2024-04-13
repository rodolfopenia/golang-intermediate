package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7.0

	fmt.Println(x)
	fmt.Println(y)

	//Manejo de errores
	myValue, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Map
	m := make(map[string]int)
	m["Key"] = 6

	fmt.Println(m["Key"])

	// Slice
	s := []int{1, 2, 3}
	// Recorriendo con range
	for index, value := range s {
		fmt.Println(index, value)
	}
	// Agregar valor nuevo a slice
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index, value)
	}
}
