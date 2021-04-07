package main

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"

	"./database"
)

// type Persona struct {
// 	gorm.Model
// 	Nombre string
// 	Direccion string
// }

const PORT = ":5000"

func main() {
	router := mux.NewRouter()
	log.Println("Escuchando en el puerto", PORT)
	router.HandleFunc("/cursos", listarCursos).Methods("GET")
	router.HandleFunc("/curso/{curso}", verCurso).Methods("GET")
	router.HandleFunc("/curso", crearCurso).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func listarCursos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	cursos := database.GetCursos()
	json.NewEncoder(w).Encode(cursos)
}

func verCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, _ := strconv.Atoi(vars["curso"])
	curso := database.VerCurso(uint(idCurso))
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(curso)
}

func crearCurso(w http.ResponseWriter, r *http.Request) {
	database.CrearCurso("Curso de Go", "Introducción al Golang")
	fmt.Fprint(w, "Ok")
}