package database

import (
	"log"
	"fmt"

	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"

	"../models"
)

const (
	user     = "postgres"
	password = "123456"
	dbname   = "golang_test"
	host     = "localhost"
	port     = "5432"
)

var db *gorm.DB
var err error

func GetConn() *gorm.DB {
	if db != nil {
		return db
	}
	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(conexion), &gorm.Config{})
	// db, err = gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		log.Println("Error en la conexión...")
		panic(err)
	}
	Migrar()
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

func CrearCurso(nombre, descripcion string) (curso models.Curso) {
	_ = GetConn()
	curso = models.Curso{Nombre: nombre, Descripcion: descripcion}
	db.Create(&curso)
	db.Last(&curso)
	return
}

func VerCurso(id uint) (curso models.Curso) {
	_ = GetConn()
	db.Find(&curso, id)
	return
}

func ActualizarCurso(id uint, nombre, descripcion string) (curso models.Curso) {
	_ = GetConn()
	db.Find(&curso, id)
	// db.Model(&curso).Update("Nombre", nombre) // Actualizar una propiedad a la vez
	// db.Model(&curso).Update("Descripcion", descripcion) // Actualizar una propiedad a la vez
	// db.Model(&curso).Updates(models.Curso{Nombre: nombre, Descripcion: descripcion}) // Actualizar todas las propiedades
	db.Model(&curso).Updates(map[string]interface{}{"Nombre": nombre, "Descripcion": descripcion}) // Actualizar solo las propiedades que están en el mapa
	return
}

func EliminarCurso(id uint) {
	_ = GetConn()
	db.Delete(&models.Curso{}, id)
}