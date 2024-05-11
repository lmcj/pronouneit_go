package services

import (
	"log"

	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func CreateEjercicio(db *gorm.DB, ejercicio models.Ejercicio) (string, bool) {
	ejercicioCreate := db.Create(&ejercicio)
	if ejercicioCreate.Error != nil {
		log.Println("Error al crear el ejercicio")
		return "Error al crear el ejercicio", false
	}

	log.Println("Ejercicio creado exitosamente...")
	return "Ejercicio creado exitosamente...", true
}

func GetEjercicios(db *gorm.DB) ([]models.Ejercicio, bool) {
	var ejercicios []models.Ejercicio
	if err := db.Preload("Nivel").Preload("Tipo").Find(&ejercicios).Error; err != nil {
		log.Println("Error al obtener los ejercicios")
		return []models.Ejercicio{}, false
	}

	log.Println("Ejercicios encontrados exitosamente...")
	return ejercicios, true
}

func GetEjercicioByID(db *gorm.DB, id uint) (models.Ejercicio, bool) {
	var ejercicio models.Ejercicio
	if err := db.Preload("Nivel").Preload("Tipo").First(&ejercicio, id).Error; err != nil {
		log.Println("Error al obtener el ejercicio")
		return models.Ejercicio{}, false
	}

	log.Println("Ejercicio encontrado exitosamente...")
	return ejercicio, true
}

func UpdateEjercicio(db *gorm.DB, ejercicio models.Ejercicio) (string, bool) {
	if err := db.Model(&models.Ejercicio{}).Where("id = ?", ejercicio.ID).Updates(&ejercicio).Error; err != nil {
		log.Println("Error al actualizar el ejercicio:", err)
		return "Error al actualizar el ejercicio", false
	}

	log.Println("Ejercicio actualizado exitosamente...")
	return "Ejercicio actualizado exitosamente", true
}

func DeleteEjercicio(db *gorm.DB, id uint) (string, bool) {
	var ejercicio models.Ejercicio
	result := db.First(&ejercicio, id)
	if result.Error != nil {
		log.Println("Error al encontrar el ejercicio")
		return "Ejercicio no encontrado", false
	}

	if err := db.Delete(&ejercicio).Error; err != nil {
		log.Println("Error al eliminar el ejercicio")
		return "Error al eliminar el ejercicio", false
	}

	log.Println("Ejercicio eliminado exitosamente...")
	return "Ejercicio eliminado exitosamente...", true
}

func GetEjerciciosPorNivel(db *gorm.DB, nivelID uint) ([]models.Ejercicio, error) {
	var ejercicios []models.Ejercicio

	if err := db.Preload("Nivel").Preload("Tipo").Where("nivel_id = ?", nivelID).Find(&ejercicios).Error; err != nil {
		log.Println("Error al obtener los ejercicios por nivel:", err)
		return nil, err
	}

	return ejercicios, nil
}
