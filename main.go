package main

import (
	"fmt"
	"strconv"
	"time"
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
	fmt.Println("MAP")
	fmt.Println(m["Key2"])

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

	// Main solo tiene la funcion de crear la subrutina
	// Pero no espera a que termine de ejecutar.
	// Para eso crearemos los channels

	//c := make(chan int)
	//go doSomething(c)
	//<-c

	// Apuntadores
	g := 25
	fmt.Println(g)
	h := &g // Apuntamos a su ubicación física
	fmt.Println(h)
	fmt.Println(*h) // El asterisco transforma el valor de memoria en su valor
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	// La rutina main se bloqueará hasta que este canal
	// Devuelva un mensaje
	c <- 1
}
