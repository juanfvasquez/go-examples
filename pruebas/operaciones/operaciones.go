package main

import "errors"

func main() {
	
}

func sumar(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No pueden ser menores a cero")
	} else {
		resultado = x + y
	}
	return
}

func restar(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No puede ser menores a cero")
	} else if y > x {
		err = errors.New("Y no puede ser menor a X")
	} else {
		resultado = x - y
	}
	return
}

func multiplicar(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No puede ser menores a cero")
	} else {
		resultado = x * y
	}
	return
}

func dividir(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No puede ser menores a cero")
	} else if y == 0 {
		err = errors.New("Error: División por cero")
	} else {
		resultado = x / y
	}
	return
}