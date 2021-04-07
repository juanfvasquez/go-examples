package database

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	// "gorm.io/driver/postgres"

	"../models"
)

var db *gorm.DB
var err error

func GetConn() *gorm.DB {
	if db != nil {
		return db
	}
	// conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := gorm.Open(postgres.Open(conexion), &gorm.Config{})
	db, err = gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		log.Println("Error en la conexión...")
		panic(err)
	}
	return db
}

func Migrar() {
	db.AutoMigrate(&models.Curso{})
}

func GetCursos() (listaCursos []models.Curso) {
	_ = GetConn()
	db.Find(&listaCursos)
	return
}

func CrearCurso(nombre, descripcion string) {
	_ = GetConn()
	curso := models.Curso{Nombre: nombre, Descripcion: descripcion}
	db.Create(&curso)
}

func VerCurso(id uint) (curso models.Curso) {
	_ = GetConn()
	db.Find(&curso, id)
	return
}