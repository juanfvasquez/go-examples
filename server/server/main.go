package main

import (
	"log"
	"net/http"
	"io"
	"fmt"
	"encoding/json"
)

type Curso struct {
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Autor Datos `json:"autor"`
}

type Datos struct {
	Nombre string `json:"nombre_autor"`
	Apellido string `json:"apellido_autor"`
	Edad int `json:"edad_autor"`
}

const PUERTO = ":5000"

func main() {
	log.Println("Corriendo en el puerto", PUERTO)
	
	http.HandleFunc("/redirect", redireccion)
	http.HandleFunc("/hola", hola)
	http.HandleFunc("/not-found", noEncontrado)
	http.HandleFunc("/error", errorPeticion)
	http.HandleFunc("/parametros", parametros)
	http.HandleFunc("/enviar-json", enviarJson)

	log.Fatal(http.ListenAndServe(PUERTO, nil))
}

func hola(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola Mundo!")
	fmt.Fprint(w, "Hola con FMT")
}

func redireccion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirigiendo a hola")
	http.Redirect(w, r, "/hola", 301)
}

func noEncontrado(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func errorPeticion(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error en el servidor", 502)
}

func parametros(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL)
	// log.Println(r.URL.RawQuery)
	mapaParametros := r.URL.Query()
	log.Println(mapaParametros)
	// log.Println(mapaParametros["idProducto"][0])
	fmt.Fprintln(w, "Recibiendo parámetros")
}

func enviarJson(w http.ResponseWriter, r *http.Request) {
	datos := Datos{Nombre: "Juan", Apellido: "Vásquez", Edad: 28}
	curso1 := Curso{1, "Curso de Go", "Curso de introducción a Go (Golang)", datos}
	curso2 := Curso{2, "Curso de Flutter", "Curso de introducción a Flutter", datos}
	curso3 := Curso{3, "Curso de Angular", "Curso de introducción a Angular", datos}
	cursos := []Curso{curso1, curso2, curso3}
	json.NewEncoder(w).Encode(cursos)
}
