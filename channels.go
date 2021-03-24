package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal := make(chan int, 10)
	salida := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go escuchar(canal, salida, wg)
	publicar(canal, salida, wg)
	salida <- struct{}{}
	wg.Wait()
}

func publicar(canal chan<- int, salida chan struct{}, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second) // Bloquea por 2 seg
	fmt.Println("Publicando...")
	canal <- 10
	canal <- 20
	canal <- 30
	canal <- 40
	canal <- 50
	canal <- 60
	canal <- 70
	canal <- 80
	canal <- 90
	canal <- 100
	// fmt.Println("Saliendo...")
	// salida <- struct{}{}
	wg.Done()
}

func escuchar(canal <-chan int, salida chan struct{}, wg *sync.WaitGroup) {
	fmt.Println("Escuchando...")
	ciclo := true
	for ciclo {
		select {
			case dato := <- canal:
				fmt.Println("Recibido: ", dato)
			case <- salida:
				fmt.Println("Salida...")
				ciclo = false
		}
	}
	wg.Done()
}