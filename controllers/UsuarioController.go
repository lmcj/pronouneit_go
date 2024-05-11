package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/configs"
	"github.com/lmcj/pronouneit_go.git/models"
	"github.com/lmcj/pronouneit_go.git/services"
)

func CreateUsuario(c *gin.Context) {
	database := configs.ConnectToDB()

	var usuario models.Usuario

	err := c.BindJSON(&usuario)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	usuarioNew, success := services.CreateUsuario(database, usuario)
	if !success {
		c.String(400, "Error")
		return
	}
	c.JSON(200, usuarioNew)
}

func GetUsuarios(c *gin.Context) {
	database := configs.ConnectToDB()

	usuarios, success := services.GetUsuarios(database)
	if !success {
		c.String(400, "Error")
		return
	}
	c.JSON(200, usuarios)

}

func GetUsuarioById(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	var usuario models.Usuario
	usuario, success := services.GetUsuarioByID(database, strconv.Itoa(idInt))
	if !success {
		c.String(400, "Error")
		return
	}
	var aux models.Usuario = models.Usuario{}
	if usuario.ID == (aux.ID) {
		c.String(404, "User not found")
		return
	}

	c.JSON(200, usuario)
}

func UpdateUsuario(c *gin.Context) {
	// Conectar a la base de datos
	database := configs.ConnectToDB()

	// Crear una variable para almacenar los datos del usuario actualizado
	var usuario models.Usuario

	// Bind JSON body al struct Usuario
	if err := c.BindJSON(&usuario); err != nil {
		c.String(400, "Bad request")
		return
	}

	// Extraer el ID del usuario de la URL
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Invalid user ID")
		return
	}

	// Buscar el usuario existente por su ID en la base de datos
	var usuarioExistente models.Usuario
	if err := database.First(&usuarioExistente, idInt).Error; err != nil {
		c.String(404, "User not found")
		return
	}

	// Aplicar las actualizaciones al usuario existente
	usuarioExistente.Nombre = usuario.Nombre
	usuarioExistente.Apellido = usuario.Apellido
	usuarioExistente.Correo = usuario.Correo

	// Llamar al servicio para actualizar el usuario en la base de datos
	message, success := services.UpdateUsuario(database, usuarioExistente)
	if !success {
		c.String(400, message)
		return
	}

	// Responder con un mensaje de éxito
	c.String(200, message)
}

func DeleteUsuario(c *gin.Context) {
	database := configs.ConnectToDB()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteUsuario(database, idInt)
	if !success {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}
