package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/configs"
	"github.com/lmcj/pronouneit_go.git/models"
	"github.com/lmcj/pronouneit_go.git/services"
)

func Login(c *gin.Context) {
	database := configs.ConnectToDB()

	var login models.Login

	err := c.BindJSON(&login)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	loginNew, success := services.Login(database, login)
	if !success {
		c.String(401, "Error en las credenciales")
		return
	}
	c.JSON(200, loginNew)
}
