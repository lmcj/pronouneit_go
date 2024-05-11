package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/configs"
	"github.com/lmcj/pronouneit_go.git/models"
	"github.com/lmcj/pronouneit_go.git/services"
)

func CreateNivel(c *gin.Context) {
	database := configs.ConnectToDB()

	var nivel models.Nivel

	if err := c.BindJSON(&nivel); err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.CreateNivel(database, nivel)
	if !success {
		c.String(400, message)
		return
	}
	c.String(200, message)
}

func GetNiveles(c *gin.Context) {
	database := configs.ConnectToDB()

	niveles, success := services.GetNiveles(database)
	if !success {
		c.String(400, "Error al obtener los niveles")
		return
	}
	c.JSON(200, niveles)
}

func GetNivelById(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	nivel, success := services.GetNivelByID(database, uint(idInt))
	if !success {
		c.String(404, "Nivel no encontrado")
		return
	}

	c.JSON(200, nivel)
}

func UpdateNivel(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Invalid level ID")
		return
	}

	var nivel models.Nivel
	if err := database.First(&nivel, idInt).Error; err != nil {
		c.String(404, "Level not found")
		return
	}

	if err := c.BindJSON(&nivel); err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateNivel(database, nivel)
	if !success {
		c.String(400, message)
		return
	}

	c.String(200, message)
}

func DeleteNivel(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteNivel(database, uint(idInt))
	if !success {
		c.String(400, message)
		return
	}
	c.String(200, message)
}

func GetNivelMaximo(c *gin.Context) {
	database := configs.ConnectToDB()

	nivelMaximo, err := services.GetNivelMaximo(database)
	if err != nil {
		c.String(500, "Error al obtener el nivel máximo")
		return
	}

	c.JSON(200, gin.H{"nivel_maximo": nivelMaximo})
}
