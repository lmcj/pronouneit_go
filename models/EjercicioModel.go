package models

import (
	"gorm.io/gorm"
)

type Ejercicio struct {
	gorm.Model
	Nombre    string      `gorm:"not null"`
	Contenido string      `gorm:"not null"`
	NivelID   uint        `gorm:"not null"`
	TipoID    uint        `gorm:"not null"`
	Nivel     Nivel       `gorm:"foreignKey:NivelID"`
	Tipo      TipoPalabra `gorm:"foreignKey:TipoID"`
}

type EjercicioDTO struct {
	ID        uint   `json:"id"`
	Nombre    string `json:"nombre"`
	Contenido string `json:"contenido"`
	Nivel     uint   `json:"nivel"`
	Tipo      uint   `json:"tipo"`
}
