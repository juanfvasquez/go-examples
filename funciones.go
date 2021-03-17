package main

import "fmt"

func main() {
	num, _, _ := retornoCompuesto()
	fmt.Println(num)
}

func saludar(nombre string) {
	fmt.Println("Hola " + nombre)
}

func retornoConTipo() int {
	return 123
}

func retornoCompuesto() (int, string, bool) {
	return 10, "Hola Mudo", false
}