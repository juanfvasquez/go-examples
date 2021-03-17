package main

import (
	"fmt"
	"errors"
)

func main() {
	resultado, err := operar(-1, 2, "resta")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}

func operar(x, y int, operacion string) (resultado int, err error) {
	if x < 0 {
		err = errors.New("El numero X es menor a 0")
		return 0, err
	}
	switch {
		case operacion == "suma":
			resultado, err = x + y, nil
		case operacion == "resta":
			resultado, err = x - y, nil
		case operacion == "multiplicacion":
			resultado, err = x * y, nil
		case operacion == "division":
			resultado, err = x / y, nil
		default:
			resultado, err = x, nil
	}
	return
}