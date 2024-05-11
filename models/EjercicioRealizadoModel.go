package models

import (
	"gorm.io/gorm"
)

type EjercicioRealizado struct {
	gorm.Model
	Resultado   string    `gorm:"not null"`
	UsuarioID   uint      `gorm:"not null"`
	EjercicioID uint      `gorm:"not null"`
	Aprobado    bool      `gorm:"default:false"`
	Usuario     Usuario   `gorm:"foreignKey:UsuarioID"`
	Ejercicio   Ejercicio `gorm:"foreignKey:EjercicioID"`
}
