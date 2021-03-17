package main

import "fmt"

func main() {
	// num := 10
	// fmt.Println(num)
	// cambiarNumero(num)
	// fmt.Printf("Después de la función queda con el valor de %d\n", num)

	// var puntero *int
	// puntero = &num
	// fmt.Println(puntero)
	// fmt.Println(*puntero)
	
	// *puntero = 100
	// fmt.Println(num)

	num2 := 20
	fmt.Printf("Antes de la función cambiarNumero: %d\n", num2)
	cambiarNumero(&num2)
	fmt.Printf("Después de la función cambiarNumero: %d\n", num2)
}

func cambiarNumero(num *int) {
	*num = *num * 1000
	fmt.Printf("En la función cambiarNumero tiene el valor de %d\n", *num)
}