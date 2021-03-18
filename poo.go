package main

import (
	"strconv"
	"fmt"
)

var var1 bool

type Persona struct {
	nombre, apellido string
	edad int
}

type empleado struct {
	idEmpleado string
	fecha string
}

func (e *empleado) registrarIngreso() {
	e.fecha = "12/12/2020"
}

func (p *Persona) toString() string {
	return "Nombre: " + p.nombre + " Apellido: " + p.apellido + ", Edad: " + strconv.Itoa(p.edad)
}

type Usuario struct {
	Persona
	empleado
	email string
	password string
}

type User struct {
	nombre string
}

func (u *Usuario) imprimirDatos() string {
	return "Nombre: " + u.nombre + " Apellido: " + u.apellido + ", Fecha: " + u.fecha 
}

func (u User) imprimirDatos() string {
	return "imprimir"
}

func (u *Usuario) inciarSesion() bool {
	//Validar email y password
	return true
}

func main() {
		usuario := Usuario{Persona{nombre: "Juan"}, empleado{"KJANSKD", "10/10/2020"}, "juan@mail.com", ""}
		fmt.Println(usuario)
		fmt.Println(usuario.nombre)
		fmt.Println(usuario.idEmpleado)
		usuario.registrarIngreso()
		fmt.Println(usuario.toString())
		fmt.Println(usuario.imprimirDatos())
}