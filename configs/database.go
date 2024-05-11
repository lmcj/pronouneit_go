package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	var err error

	dsn := "root:root@tcp(localhost:3306)/dbpronouneit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	log.Println("Conectado exitosamente a la DB...")

	return db
}
