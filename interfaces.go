package main

import "fmt"

type Animal interface {
	Mover() string
	Comer(comida string) string
}

type Perro struct {
	nombre string
	patas string
}

type Serpiente struct {
	nombre string
}

func (p Perro) Mover() string {
	return "Caminando... moviendo " + p.patas + " patas"
}

func (s Serpiente) Mover() string {
	return "Reptando..."
}

func iniciar(animal Animal) {
	fmt.Println(animal.Mover())
}

func main() {
	perro := Perro{"ajndsfkj", "4"}
	serpiente := Serpiente{"mnkjb"}

	iniciar(perro)
	iniciar(serpiente)
}