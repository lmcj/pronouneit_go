package main

import (
	"log"

	"github.com/lmcj/pronouneit_go.git/configs"
	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func main() {

	var db *gorm.DB = configs.ConnectToDB()
	db.AutoMigrate(
		&models.Usuario{},
		&models.EjercicioRealizado{},
		&models.Ejercicio{},
		&models.Nivel{},
		&models.TipoPalabra{})
	log.Println("Migración exitosa...")
}
