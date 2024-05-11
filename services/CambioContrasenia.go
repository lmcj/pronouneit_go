package services

import (
	"log"

	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func CambiarContrasenia(db *gorm.DB, id uint, nuevaContrasenia string) error {
	var usuario models.Usuario
	if err := db.First(&usuario, id).Error; err != nil {
		return err
	}

	usuario.Contrasenia = Hash256(nuevaContrasenia)
	if err := db.Save(&usuario).Error; err != nil {
		log.Println("Error al actualizar la contraseña del usuario:", err)
		return err
	}

	log.Println("Contraseña actualizada exitosamente para el usuario:", usuario.ID)
	return nil
}

func CompararContrasenia(contraseniaActual, contraseniaNueva string) bool {
	return contraseniaActual == Hash256(contraseniaNueva)
}
