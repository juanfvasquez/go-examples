package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./hola.txt")
	if err != nil {
		fmt.Println("Algo salió mal...")
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}