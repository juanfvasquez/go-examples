package main 

import "fmt"

func main() {
	x := 10
	y := 10
	if x > y {
		fmt.Println("X es mayor")
	} else if y > x {
		fmt.Println("Y es mayor")
	} else {
		fmt.Println("X e Y son iguales ")
	}
}