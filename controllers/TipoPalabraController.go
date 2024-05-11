package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/configs"
	"github.com/lmcj/pronouneit_go.git/models"
	"github.com/lmcj/pronouneit_go.git/services"
)

func CreateTipoPalabra(c *gin.Context) {
	database := configs.ConnectToDB()

	var tipoPalabra models.TipoPalabra

	if err := c.BindJSON(&tipoPalabra); err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.CreateTipoPalabra(database, tipoPalabra)
	if !success {
		c.String(400, message)
		return
	}
	c.String(200, message)
}

func GetTiposPalabra(c *gin.Context) {
	database := configs.ConnectToDB()

	tiposPalabra, success := services.GetTiposPalabra(database)
	if !success {
		c.String(400, "Error al obtener los tipos de palabra")
		return
	}
	c.JSON(200, tiposPalabra)
}

func GetTipoPalabraById(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	tipoPalabra, success := services.GetTipoPalabraByID(database, uint(idInt))
	if !success {
		c.String(404, "Tipo de palabra no encontrado")
		return
	}

	c.JSON(200, tipoPalabra)
}

func UpdateTipoPalabra(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Invalid type ID")
		return
	}

	var tipoPalabra models.TipoPalabra
	if err := database.First(&tipoPalabra, idInt).Error; err != nil {
		c.String(404, "Type of word not found")
		return
	}

	if err := c.BindJSON(&tipoPalabra); err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateTipoPalabra(database, tipoPalabra)
	if !success {
		c.String(400, message)
		return
	}

	c.String(200, message)
}

func DeleteTipoPalabra(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteTipoPalabra(database, uint(idInt))
	if !success {
		c.String(400, message)
		return
	}
	c.String(200, message)
}

func CreateTiposPalabras(c *gin.Context) {
	database := configs.ConnectToDB()

	var tiposPalabras []models.TipoPalabra

	if err := c.BindJSON(&tiposPalabras); err != nil {
		c.String(400, "Bad request")
		return
	}

	messages, success := services.CreateTiposPalabras(database, tiposPalabras)
	if !success {
		c.String(400, "Error al crear los tipos de palabras")
		return
	}

	c.JSON(200, gin.H{"messages": messages})
}
