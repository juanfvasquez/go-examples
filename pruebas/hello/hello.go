package main

import "fmt"

func main() {
	fmt.Println(saludar("Juan"))
}

func saludar(nombre string) string {
	if nombre == "" {
		nombre = "mundo!"
	}
	return "Hola " + nombre
}