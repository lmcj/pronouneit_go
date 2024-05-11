package services

import (
	"log"

	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func CreateNivel(db *gorm.DB, nivel models.Nivel) (string, bool) {
	nivelCreate := db.Create(&nivel)
	if nivelCreate.Error != nil {
		log.Println("Error al crear el nivel")
		return "Error al crear el nivel", false
	}

	log.Println("Nivel creado exitosamente...")
	return "Nivel creado exitosamente...", true
}

func GetNiveles(db *gorm.DB) ([]models.Nivel, bool) {
	var niveles []models.Nivel
	if err := db.Find(&niveles).Error; err != nil {
		log.Println("Error al obtener los niveles")
		return []models.Nivel{}, false
	}

	log.Println("Niveles encontrados exitosamente...")
	return niveles, true
}

func GetNivelByID(db *gorm.DB, id uint) (models.Nivel, bool) {
	var nivel models.Nivel
	if err := db.First(&nivel, id).Error; err != nil {
		log.Println("Error al obtener el nivel")
		return models.Nivel{}, false
	}

	log.Println("Nivel encontrado exitosamente...")
	return nivel, true
}

func UpdateNivel(db *gorm.DB, nivel models.Nivel) (string, bool) {
	if err := db.Model(&models.Nivel{}).Where("id = ?", nivel.ID).Updates(&nivel).Error; err != nil {
		log.Println("Error al actualizar el nivel:", err)
		return "Error al actualizar el nivel", false
	}

	log.Println("Nivel actualizado exitosamente...")
	return "Nivel actualizado exitosamente", true
}

func DeleteNivel(db *gorm.DB, id uint) (string, bool) {
	var nivel models.Nivel
	result := db.First(&nivel, id)
	if result.Error != nil {
		log.Println("Error al encontrar el nivel")
		return "Nivel no encontrado", false
	}

	if err := db.Delete(&nivel).Error; err != nil {
		log.Println("Error al eliminar el nivel")
		return "Error al eliminar el nivel", false
	}

	log.Println("Nivel eliminado exitosamente...")
	return "Nivel eliminado exitosamente...", true
}

func GetNivelMaximo(db *gorm.DB) (int, error) {
	var nivelMaximo int

	result := db.Model(&models.Nivel{}).Select("MAX(nivel)").Scan(&nivelMaximo)
	if result.Error != nil {
		log.Println("Error al obtener el nivel máximo:", result.Error)
		return 0, result.Error
	}

	return nivelMaximo, nil
}
