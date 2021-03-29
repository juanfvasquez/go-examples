package main

import "fmt"

func main() {
	monedas := make(map[string]string)
	monedas["colombia"] = "Peso colombiano"
	monedas["EEUU"] = "Dolar"
	monedas["España"] = "Euro"
	fmt.Println(monedas)

	// delete(monedas, "España")
	// fmt.Println(len(monedas))
	for key, value := range monedas {
		fmt.Println("Clave: ", key, "Valor: ", value)
	}

	arreglo := []int{10, 20, 30, 40, 50}
	for i, _ := range arreglo {
		fmt.Println("Elemento:", i)
	}
}