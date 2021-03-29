package main

import (
	"fmt"
	// "strings"
	"time"
)

func main() {
	fmt.Println("Inicio")
	for i := 1; i <= 20; i++ {
		fmt.Println("Lanzando goroutine ", i)
		go func(num int) {
			time.Sleep(2 * time.Second)
			fmt.Printf("Go routine # %d\n", num)
		}(i)
	}
	fmt.Println("Fin")
	var espera string
	fmt.Scanln(&espera)
}


// func main() {
// 	fmt.Println("Imprimiendo nombre...")
// 	go imprimirNombre("Pedro")
// 	fmt.Println("Fin de impresión")
// 	var espera string
// 	fmt.Scanln(&espera)
// }

// func imprimirNombre(nombre string) {
// 	letras := strings.Split(nombre, "")
// 	for _, letra := range letras {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println(letra)
// 	}
// }