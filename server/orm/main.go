package main

import (
	"log"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"

	"./database"
	"./models"
	"./respuestas"
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
	router.HandleFunc("/curso/{curso}", actualizarCurso).Methods("PUT")
	router.HandleFunc("/curso/{curso}", eliminarCurso).Methods("DELETE")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func listarCursos(w http.ResponseWriter, r *http.Request) {
	cursos := database.GetCursos()
	respuestas.ResponderJSON(cursos, w)
}

func verCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])
	if responderSiHayError(err, "ID no válido", -1, w) {
		return
	}
	curso := database.VerCurso(uint(idCurso))
	respuestas.ResponderJSON(curso, w)
}

func crearCurso(w http.ResponseWriter, r *http.Request) {
	var body models.Curso
	err := json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body...")
	curso := database.CrearCurso(body.Nombre, body.Descripcion)
	respuestas.ResponderJSON(curso, w)
}

func actualizarCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])
	if responderSiHayError(err, "ID no válido", -1, w) {
		return
	}
	var body models.Curso
	err = json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body...")
	curso := database.ActualizarCurso(uint(idCurso), body.Nombre, body.Descripcion)
	respuestas.ResponderJSON(curso, w)
}

func eliminarCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])
	if responderSiHayError(err, "ID no válido", -1, w) {
		return
	}
	database.EliminarCurso(uint(idCurso))
	respuestas.ResponderJSON(models.Respuesta{Mensaje: "Ok", Codigo: 1}, w)
}

func validarError(err error, mensaje string) {
	if err != nil {
		log.Println(mensaje)
		panic(err)
	}
}

func responderSiHayError(err error, mensaje string, codigo int, w http.ResponseWriter) bool{
	if err != nil {
		respuestas.ResponderJSON(
			models.Respuesta{Mensaje: mensaje, Codigo: codigo},
			w,
		)
		return true
	}
	return false
}