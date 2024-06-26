package services

import (
	"log"
	"strings"

	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func CreateEjercicioRealizado(db *gorm.DB, ejercicioRealizado models.EjercicioRealizado) (string, bool) {

	var ejercicio models.Ejercicio

	ejercicioRealizado.Resultado = TranscribeBase64Audio(ejercicioRealizado.Resultado)

	if ejercicioRealizado.Resultado == "" {
		log.Println("Error al transcribir el audio")
		return "Error al transcribir el audio", false
	}

	log.Println("ejercicio realizado: " + ejercicioRealizado.Resultado)

	db.First(&ejercicio, ejercicioRealizado.EjercicioID)

	if strings.EqualFold(ejercicioRealizado.Resultado, ejercicio.Contenido) {
		ejercicioRealizado.Aprobado = true
	}

	ejercicioRealizadoCreate := db.Create(&ejercicioRealizado)
	if ejercicioRealizadoCreate.Error != nil {
		log.Println("Error al crear el ejercicio realizado")
		return "Error al crear el ejercicio realizado", false
	}

	var usuario models.Usuario
	db.First(&usuario, ejercicioRealizado.UsuarioID)

	if ejercicioRealizado.Aprobado {
		usuario.Racha++
		racha := usuario.Racha
		valorBase := 10.0
		factor := 1.0 + 0.2*float64(racha-1)
		usuario.XP += int(valorBase * factor)

		usuario.UltimoResultado = true

		if usuario.XP >= 100 {
			usuario.Nivel++
			usuario.XP = 0
		}

		db.Save(&usuario)
	} else {
		usuario.Racha = 1
		usuario.UltimoResultado = false

		db.Save(&usuario)
	}

	log.Println("Ejercicio realizado creado exitosamente...")
	return "Ejercicio realizado creado exitosamente...", true
}

func GetEjerciciosRealizadosByUsuarioID(db *gorm.DB, usuarioID uint) ([]models.EjercicioRealizado, bool) {
	var ejerciciosRealizados []models.EjercicioRealizado
	if err := db.Where("usuario_id = ?", usuarioID).Preload("Ejercicio").Find(&ejerciciosRealizados).Error; err != nil {
		log.Println("Error al obtener los ejercicios realizados")
		return []models.EjercicioRealizado{}, false
	}

	log.Println("Ejercicios realizados encontrados exitosamente...")
	return ejerciciosRealizados, true
}

func GetUltimoEjercicioRealizadoByUsuarioID(db *gorm.DB, usuarioID uint) (models.EjercicioRealizado, bool) {
	var ejercicioRealizado models.EjercicioRealizado
	if err := db.Where("usuario_id = ?", usuarioID).Order("created_at desc").Preload("Ejercicio").First(&ejercicioRealizado).Error; err != nil {
		log.Println("Error al obtener el último ejercicio realizado")
		return models.EjercicioRealizado{}, false
	}

	log.Println("Último ejercicio realizado encontrado exitosamente...")
	return ejercicioRealizado, true
}

func GetEjercicioRealizadoByID(db *gorm.DB, id uint) (models.EjercicioRealizado, bool) {
	var ejercicioRealizado models.EjercicioRealizado
	if err := db.Preload("Usuario").Preload("Ejercicio").First(&ejercicioRealizado, id).Error; err != nil {
		log.Println("Error al obtener el ejercicio realizado")
		return models.EjercicioRealizado{}, false
	}

	log.Println("Ejercicio realizado encontrado exitosamente...")
	return ejercicioRealizado, true
}

func UpdateEjercicioRealizado(db *gorm.DB, ejercicioRealizado models.EjercicioRealizado) (string, bool) {
	if err := db.Model(&models.EjercicioRealizado{}).Where("id = ?", ejercicioRealizado.ID).Updates(&ejercicioRealizado).Error; err != nil {
		log.Println("Error al actualizar el ejercicio realizado:", err)
		return "Error al actualizar el ejercicio realizado", false
	}

	log.Println("Ejercicio realizado actualizado exitosamente...")
	return "Ejercicio realizado actualizado exitosamente", true
}

func DeleteEjercicioRealizado(db *gorm.DB, id uint) (string, bool) {
	var ejercicioRealizado models.EjercicioRealizado
	result := db.First(&ejercicioRealizado, id)
	if result.Error != nil {
		log.Println("Error al encontrar el ejercicio realizado")
		return "Ejercicio realizado no encontrado", false
	}

	if err := db.Delete(&ejercicioRealizado).Error; err != nil {
		log.Println("Error al eliminar el ejercicio realizado")
		return "Error al eliminar el ejercicio realizado", false
	}

	log.Println("Ejercicio realizado eliminado exitosamente...")
	return "Ejercicio realizado eliminado exitosamente...", true
}
