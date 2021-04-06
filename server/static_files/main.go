package main

import (
	"net/http"
	"log"
	"html/template"
)

const PUERTO = ":5000"

func main() {
	log.Println("Corriendo en el puerto", PUERTO)

	http.HandleFunc("/archivo", responderHTML)

	log.Fatal(http.ListenAndServe(PUERTO, nil))
}

func responderHTML(w http.ResponseWriter, r *http.Request) {
	plantilla, err := template.ParseFiles("./public/archivo1.html")
	if err != nil {
		log.Println("Error generando el template")
		panic(err)
	}
	plantilla.Execute(w, nil)
}