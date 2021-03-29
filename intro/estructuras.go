package main

import (
	"fmt"
	"strconv"
)

type Humano struct {
	nombre string
	edad int
	mapa map[string]int
}

func (humano Humano) imprimir() string {
	return "Nombre: " + humano.nombre + ", edad: " + strconv.Itoa(humano.edad)
}

func (humano *Humano) cambiarValores(nombre string, edad int) {
	humano.nombre = nombre
	humano.edad = edad
}

type Estudiante struct {
	nombre string
	calificaciones []float32
	promedio
}

func (estudiante Estudiante) calcularProm() float32 {
	for _, val := range estudiante.calificaciones {
		// calculo
	}
	estudiante.promedio = promedio
}

func main() {
	// persona := Humano{"Juan", 30, mapa}
	persona := Humano{nombre: "Juan", edad: 30}
	fmt.Println(persona.imprimir())
	persona2 := new(Humano)
	fmt.Println(persona2.imprimir())
	
	persona.nombre = "David"
	persona2.nombre = "Pedro"
	persona2.edad = 20
	fmt.Println(persona2.imprimir())

	persona.cambiarValores("Pablo", 32)
	fmt.Println(persona.imprimir())

	persona2.cambiarValores("José", 37)
	fmt.Println(persona2.imprimir())
	// grupo := new(Grupo)
	// fmt.Println(grupo)

	persona.mapa["algo"] = 12
}