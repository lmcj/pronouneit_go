package services

import (
	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func GetNextExercise(db *gorm.DB, userID string) (models.Ejercicio, error) {
	var usuario models.Usuario

	if err := db.First(&usuario, "id = ?", userID).Error; err != nil {
		return models.Ejercicio{}, err
	}

	var ejercicio models.Ejercicio
	var tipoEjercicio int

	if usuario.UltimoResultado {
		tipoEjercicio = (usuario.UltimoTipoEjercicio % 3) + 1
	} else {
		tipoEjercicio = usuario.UltimoTipoEjercicio
	}

	if err := db.Where("nivel_id = ? AND tipo_id = ?", usuario.Nivel, tipoEjercicio).Order("RAND()").First(&ejercicio).Error; err != nil {
		return models.Ejercicio{}, err
	}

	usuario.UltimoTipoEjercicio = tipoEjercicio
	if err := db.Save(&usuario).Error; err != nil {
		return models.Ejercicio{}, err
	}

	return ejercicio, nil
}

func UpdateUserResult(db *gorm.DB, userID string, correcto bool) error {
	var usuario models.Usuario
	if err := db.First(&usuario, "id = ?", userID).Error; err != nil {
		return err
	}

	usuario.UltimoResultado = correcto
	return db.Save(&usuario).Error
}
