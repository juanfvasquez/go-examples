package main

import (
	"strconv"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(numero int) string {
	if numero <= 0 {
		return "No es un número válido"
	}
	if numero % 3 == 0 && numero % 5 == 0 {
		return "fizzbuzz"
	}
	if numero % 3 == 0 {
		return "fizz"
	}
	if numero % 5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(numero)
}