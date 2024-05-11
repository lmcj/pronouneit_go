package models

import (
	"gorm.io/gorm"
)

type Nivel struct {
	gorm.Model
	Nivel       int    `gorm:"not null"`
	Descripcion string `gorm:"not null"`
}
