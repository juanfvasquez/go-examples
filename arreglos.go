package main

import "fmt"

func main() {
	var arreglo [5]int
	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = i * 10
	}
	fmt.Println(arreglo)

	var arreglo2 [3][2]int
	fmt.Println(arreglo2)

	fmt.Printf("Posicion 1, 1 : %d\n", arreglo2[1][1])

	arreglo3 := [3]string{"A", "B", "C"}
	arreglo3[1] = "3"
	fmt.Println(arreglo3)
}