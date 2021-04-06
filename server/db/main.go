package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	user = "postgres"
	password = "123456"
	dbname = "golang_test"
	host = "localhost"
	port = "5432"
)

type Categoria struct {
	id int
	nombre string
	descripcion string
}

var db *sql.DB
var err error

func main() {
	listarCategorias()
}

func conectar() {
	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", conexion)
	if err != nil {
		log.Println("Error en la conexión")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("Error en ping")
		panic(err)
	}
	log.Println("Conexión abierta")
}

func listarCategorias() {
	conectar()
	defer func () {
		db.Close()
		log.Println("Conexión Cerrada")
	}()

	query := "SELECT * FROM listarCategorias()"
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error en la consulta")
		panic(err)
	}
	listaCat := []Categoria{}
	for rows.Next() {
		cat := Categoria{}
		rows.Scan(
			&cat.id,
			&cat.nombre,
			&cat.descripcion,
		)
		listaCat = append(listaCat, cat)
	}
	log.Println(listaCat)
}