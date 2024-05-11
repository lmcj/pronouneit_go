package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/configs"
	"github.com/lmcj/pronouneit_go.git/models"
	"github.com/lmcj/pronouneit_go.git/services"
)

func CreateEjercicio(c *gin.Context) {
	database := configs.ConnectToDB()

	var ejercicio models.Ejercicio

	if err := c.BindJSON(&ejercicio); err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.CreateEjercicio(database, ejercicio)
	if !success {
		c.String(400, message)
		return
	}
	c.String(200, message)
}

func GetEjercicios(c *gin.Context) {
	database := configs.ConnectToDB()

	ejercicios, success := services.GetEjercicios(database)
	if !success {
		c.String(400, "Error al obtener los ejercicios")
		return
	}
	c.JSON(200, ejercicios)
}

func GetEjercicioById(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	ejercicio, success := services.GetEjercicioByID(database, uint(idInt))
	if !success {
		c.String(404, "Ejercicio no encontrado")
		return
	}

	c.JSON(200, ejercicio)
}

func UpdateEjercicio(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Invalid exercise ID")
		return
	}

	var ejercicio models.Ejercicio
	if err := database.First(&ejercicio, idInt).Error; err != nil {
		c.String(404, "Exercise not found")
		return
	}

	if err := c.BindJSON(&ejercicio); err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateEjercicio(database, ejercicio)
	if !success {
		c.String(400, message)
		return
	}

	c.String(200, message)
}

func DeleteEjercicio(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteEjercicio(database, uint(idInt))
	if !success {
		c.String(400, message)
		return
	}
	c.String(200, message)
}

func GetEjerciciosPorNivel(c *gin.Context) {
	database := configs.ConnectToDB()

	nivelID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "ID de nivel inválido")
		return
	}

	ejercicios, err := services.GetEjerciciosPorNivel(database, uint(nivelID))
	if err != nil {
		c.String(500, "Error al obtener los ejercicios por nivel")
		return
	}

	c.JSON(200, ejercicios)
}
