package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/configs"
	"github.com/lmcj/pronouneit_go.git/models"
	"github.com/lmcj/pronouneit_go.git/services"
)

func CreateEjercicioRealizado(c *gin.Context) {
	database := configs.ConnectToDB()

	var ejercicioRealizado models.EjercicioRealizado

	userID, exists := c.Get("userID")

	if !exists {
		c.String(400, "ID de usuario inválido")
		return
	}

	userIDStr, ok := userID.(string)

	if !ok {
		c.String(400, "ID de usuario no String")
		return
	}

	userIDUint64, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.String(400, "Error al convertir userID a uint64:")
		return
	}

	userIDUint := uint(userIDUint64)

	ejercicioRealizado.UsuarioID = userIDUint

	if err := c.BindJSON(&ejercicioRealizado); err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.CreateEjercicioRealizado(database, ejercicioRealizado)
	if !success {
		c.String(400, message)
		return
	}
	c.String(200, message)
}

func GetEjerciciosRealizadosByUsuarioID(c *gin.Context) {
	database := configs.ConnectToDB()

	usuarioID := c.Param("usuario_id")
	usuarioIDInt, err := strconv.Atoi(usuarioID)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	ejerciciosRealizados, success := services.GetEjerciciosRealizadosByUsuarioID(database, uint(usuarioIDInt))
	if !success {
		c.String(400, "Error al obtener los ejercicios realizados")
		return
	}
	c.JSON(200, ejerciciosRealizados)
}

func GetEjercicioRealizadoById(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	ejercicioRealizado, success := services.GetEjercicioRealizadoByID(database, uint(idInt))
	if !success {
		c.String(404, "Ejercicio realizado no encontrado")
		return
	}

	c.JSON(200, ejercicioRealizado)
}

func UpdateEjercicioRealizado(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Invalid exercise performed ID")
		return
	}

	var ejercicioRealizado models.EjercicioRealizado
	if err := database.First(&ejercicioRealizado, idInt).Error; err != nil {
		c.String(404, "Exercise performed not found")
		return
	}

	if err := c.BindJSON(&ejercicioRealizado); err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateEjercicioRealizado(database, ejercicioRealizado)
	if !success {
		c.String(400, message)
		return
	}

	c.String(200, message)
}

func DeleteEjercicioRealizado(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteEjercicioRealizado(database, uint(idInt))
	if !success {
		c.String(400, message)
		return
	}
	c.String(200, message)
}
