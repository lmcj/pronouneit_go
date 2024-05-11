package services

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lmcj/pronouneit_go.git/models"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, credentialsUser models.Login) (models.UsuarioDTO, bool) {
	var user models.Usuario
	response := db.Where("correo = ?", credentialsUser.Correo).First(&user)

	if response.Error != nil {
		return models.UsuarioDTO{}, false
	}

	if user.Contrasenia != Hash256(credentialsUser.Contrasenia) {
		return models.UsuarioDTO{}, false
	}

	userDTO := models.MapUsuarioToDTO(user)
	token, error := GenerateToken(userDTO, "secret")
	if error != nil {
		return models.UsuarioDTO{}, false
	}
	userDTO.Token = token

	return userDTO, true
}

func Hash256(password string) string {
	passwordBytes := []byte(password)

	hasher := sha256.New()
	hasher.Write(passwordBytes)

	hashSum := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashSum)

	return hashString
}

func GenerateToken(dto models.UsuarioDTO, secretKey string) (string, error) {

	claims := jwt.MapClaims{
		"id":     dto.ID,
		"nombre": dto.Nombre,
		"email":  dto.Correo,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours from now)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
