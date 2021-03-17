package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var nombre string
	nombre = "Juan"
	fmt.Println(nombre)

	var x int
	fmt.Println(x)
	
	var bandera bool
	fmt.Println(bandera)
	
	// int -> 0, bool false, string ""

	x = 5
	// var strX string
	// strX = strconv.Itoa(x)
	// fmt.Println("El valor de X es : " + strX)

	// %.2f, %t, %s, %d
	fmt.Printf("El valor de X es %d y el nombre es %s\n", x, nombre)

	// var num int
	// var apellido string

	// num = 1
	// apellido = "Vasquez"

	// var nombre string
	// nombre = "Juan"

	num := 1
	apellido := "Vasquez"

	num = 2

	const constante 

	fmt.Printf("El num es %d y el apellido es %s", num, apellido)
}