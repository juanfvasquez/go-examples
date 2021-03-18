package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
)

func main() {
	fmt.Println("Inicio del main")
	ejecutarFuncion()
	fmt.Println("Fin del main")
}

func verDefer() {
	fmt.Println("Inicio de método")
	defer fmt.Println("Esto es ejecutado con DEFER")
	fmt.Println("Este es el fin del método")
}

func leerArchivo() {
	archivo, err := os.Open("./holaaa.txt")
	defer func() {
		fmt.Println("Cerrando archivo...")
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if err != nil {
		fmt.Println("Algo salió mal...")
		panic(err)
	}
	scanner := bufio.NewScanner(archivo)
	for i := 1; scanner.Scan(); {
		linea := scanner.Text()
		fmt.Printf("Línea %d -> %s\n", i, linea)
		i++
		panic(errors.New("Simulando error"))
	}
}

func ejecutarFuncion() {
	fmt.Println("Inicio de tercera función")
	leerArchivo()
	fmt.Println("Fin de la tercera función")
}