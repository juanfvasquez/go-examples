package main

import "fmt"

type minuto int

func (m minuto) aEntero() int {
	return int(m) * 100
}

type hora int 

func main() {
	var tiempo minuto = 3
	fmt.Println(tiempo)

	var convertido int
	convertido = tiempo.aEntero()
	fmt.Println(convertido)

	var horas hora
	fmt.Println(hora.aEntero())
}