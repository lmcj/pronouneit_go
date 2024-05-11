package services

import (
	"log"

	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func CreateTipoPalabra(db *gorm.DB, tipoPalabra models.TipoPalabra) (string, bool) {
	tipoPalabraCreate := db.Create(&tipoPalabra)
	if tipoPalabraCreate.Error != nil {
		log.Println("Error al crear el tipo de palabra")
		return "Error al crear el tipo de palabra", false
	}

	log.Println("Tipo de palabra creado exitosamente...")
	return "Tipo de palabra creado exitosamente...", true
}

func GetTiposPalabra(db *gorm.DB) ([]models.TipoPalabra, bool) {
	var tiposPalabra []models.TipoPalabra
	if err := db.Find(&tiposPalabra).Error; err != nil {
		log.Println("Error al obtener los tipos de palabra")
		return []models.TipoPalabra{}, false
	}

	log.Println("Tipos de palabra encontrados exitosamente...")
	return tiposPalabra, true
}

func GetTipoPalabraByID(db *gorm.DB, id uint) (models.TipoPalabra, bool) {
	var tipoPalabra models.TipoPalabra
	if err := db.First(&tipoPalabra, id).Error; err != nil {
		log.Println("Error al obtener el tipo de palabra")
		return models.TipoPalabra{}, false
	}

	log.Println("Tipo de palabra encontrado exitosamente...")
	return tipoPalabra, true
}

func UpdateTipoPalabra(db *gorm.DB, tipoPalabra models.TipoPalabra) (string, bool) {
	if err := db.Model(&models.TipoPalabra{}).Where("id = ?", tipoPalabra.ID).Updates(&tipoPalabra).Error; err != nil {
		log.Println("Error al actualizar el tipo de palabra:", err)
		return "Error al actualizar el tipo de palabra", false
	}

	log.Println("Tipo de palabra actualizado exitosamente...")
	return "Tipo de palabra actualizado exitosamente", true
}

func DeleteTipoPalabra(db *gorm.DB, id uint) (string, bool) {
	var tipoPalabra models.TipoPalabra
	result := db.First(&tipoPalabra, id)
	if result.Error != nil {
		log.Println("Error al encontrar el tipo de palabra")
		return "Tipo de palabra no encontrado", false
	}

	if err := db.Delete(&tipoPalabra).Error; err != nil {
		log.Println("Error al eliminar el tipo de palabra")
		return "Error al eliminar el tipo de palabra", false
	}

	log.Println("Tipo de palabra eliminado exitosamente...")
	return "Tipo de palabra eliminado exitosamente...", true
}
