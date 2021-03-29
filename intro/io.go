package main

import "fmt"

func main() {
	var nombre string
	fmt.Println("Ingresa tu nombre:")
	fmt.Scanf("%s", &nombre)
	saludar(nombre)

	var edad int
	fmt.Println("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)

	fmt.Printf("Tu edad es %d", edad)
}

func saludar(nombre string) {
	fmt.Println("Hola " + nombre)
}