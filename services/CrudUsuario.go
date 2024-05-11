package services

import (
	"log"

	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func CreateUsuario(db *gorm.DB, usuario models.Usuario) (string, bool) {
	usuario.Contrasenia = Hash256(usuario.Contrasenia)

	usuarioCreate := db.Create(&usuario)
	if usuarioCreate.Error != nil {
		log.Println("Error al crear el usuario")
		return "Error al crear el usuario", false
	}
	log.Println("Usuario creado exitosamente...")
	return "Usuario creado exitosamente...", true
}

func GetUsuarios(db *gorm.DB) ([]models.Usuario, bool) {
	var usuarios []models.Usuario
	usuariosList := db.Find(&usuarios)
	if usuariosList.Error != nil {
		log.Println("Error al obtener los usuarios")
		return []models.Usuario{}, false
	}

	log.Println("Usuarios encontrados exitosamente...")
	return usuarios, true
}

func GetUsuarioByID(db *gorm.DB, id string) (models.Usuario, bool) {
	var usuario models.Usuario
	usuarioGet := db.First(&usuario, id)
	if usuarioGet.Error != nil {
		log.Println("Error al obtener el usuario")
		return models.Usuario{}, false
	}

	log.Println("Usuario encontrado exitosamente...")
	return usuario, true
}

func UpdateUsuario(db *gorm.DB, usuario models.Usuario) (string, bool) {
	usuarioUpdate := db.Save(&usuario)
	if usuarioUpdate.Error != nil {
		log.Println("Error al actualizar el usuario:", usuarioUpdate.Error)
		return "Error al actualizar el usuario", false
	}
	log.Println("Usuario actualizado exitosamente:", usuario.ID)
	return "Usuario actualizado exitosamente", true
}

func DeleteUsuario(db *gorm.DB, id int) (string, bool) {
	var usuario models.Usuario
	var aux models.Usuario = models.Usuario{}
	db.First(&usuario, id)
	if usuario.ID == (aux.ID) {
		log.Println("Usuario no encontrado...")
		return "Usuario no encontrado...", false
	}

	usuarioDelete := db.Delete(&usuario)
	if usuarioDelete.Error != nil {
		log.Println("Error al eliminar el usuario")
		return "Error al eliminar el usuario", false
	}
	log.Println("Usuario eliminado exitosamente...")
	return "Usuario eliminado exitosamente...", true
}

func GetEjerciciosRealizadosPorUsuario(db *gorm.DB, userID uint) ([]models.EjercicioRealizado, error) {
	var ejerciciosRealizados []models.EjercicioRealizado

	if err := db.Preload("Usuario").Preload("Ejercicio").Preload("Ejercicio.Nivel").Preload("Ejercicio.Tipo").Where("usuario_id = ?", userID).Find(&ejerciciosRealizados).Error; err != nil {
		log.Println("Error al obtener los ejercicios realizados por el usuario:", err)
		return nil, err
	}

	return ejerciciosRealizados, nil
}
