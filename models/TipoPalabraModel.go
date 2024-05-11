package models

import (
	"gorm.io/gorm"
)

type TipoPalabra struct {
	gorm.Model
	Grupo string `gorm:"not null"`
}
