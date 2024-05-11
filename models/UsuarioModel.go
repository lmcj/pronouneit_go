package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre      string `gorm:"size:50;not null" json:"nombre"`
	Apellido    string `gorm:"size:50;not null" json:"apellido"`
	Correo      string `gorm:"size:50;not null;unique" json:"correo" validate:"required,email"`
	Contrasenia string `gorm:"size:255;not null" json:"contrasenia"`
	FotoURL     string `gorm:"size:255" json:"fotoURL"`
}

type UsuarioDTO struct {
	ID       uint   `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
	FotoURL  string `json:"fotoURL"`
}

func MapUsuarioToDTO(usuario Usuario) UsuarioDTO {
	return UsuarioDTO{
		ID:       usuario.ID,
		Nombre:   usuario.Nombre,
		Apellido: usuario.Apellido,
		Correo:   usuario.Correo,
		FotoURL:  usuario.FotoURL,
	}
}