package main

import "fmt"

func main() {
	// var arreglo [5]int
	var slice []int
	fmt.Println(slice)
	
	slice = append(slice, 10)
	fmt.Println(cap(slice))
	// arreglo := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice := arreglo[:]
	// fmt.Println(arreglo)
	// fmt.Println(slice)

	// slice = append(slice, 11)
	// fmt.Println(slice)

	fuente := []int{10, 20, 30}
	destino := make([]int, 2)

	fmt.Println(fuente)
	fmt.Println(destino)
	
	copy(destino, fuente)
	fmt.Println(destino)

	slice100 := make([]int, 5, 100)
	fmt.Println(slice100)
	fmt.Println(len(slice100))
	fmt.Println(cap(slice100))
}