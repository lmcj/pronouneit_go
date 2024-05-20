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

func CreateAdmin(c *gin.Context) {

	database := configs.ConnectToDB()

	var usuario models.Usuario

	err := c.BindJSON(&usuario)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	usuario.Rol = "admin"

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
	database := configs.ConnectToDB()

	var usuario models.Usuario

	if err := c.BindJSON(&usuario); err != nil {
		c.String(400, "Bad request")
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Invalid user ID")
		return
	}

	var usuarioExistente models.Usuario
	if err := database.First(&usuarioExistente, idInt).Error; err != nil {
		c.String(404, "User not found")
		return
	}

	usuarioExistente.Nombre = usuario.Nombre
	usuarioExistente.Apellido = usuario.Apellido
	usuarioExistente.Correo = usuario.Correo

	message, success := services.UpdateUsuario(database, usuarioExistente)
	if !success {
		c.String(400, message)
		return
	}

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

func CambioContrasenia(c *gin.Context) {
	database := configs.ConnectToDB()

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.String(400, "ID de usuario inválido")
		return
	}

	var requestBody struct {
		ContraseniaActual string `json:"contraseniaActual"`
		NuevaContrasenia  string `json:"nuevaContrasenia"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.String(400, "Solicitud JSON inválida")
		return
	}

	var usuario models.Usuario
	if err := database.First(&usuario, id).Error; err != nil {
		c.String(404, "Usuario no encontrado")
		return
	}

	if !services.CompararContrasenia(usuario.Contrasenia, requestBody.ContraseniaActual) {
		c.String(401, "La contraseña actual es incorrecta")
		return
	}

	if err := services.CambiarContrasenia(database, uint(id), requestBody.NuevaContrasenia); err != nil {
		c.String(500, "Error al actualizar la contraseña del usuario")
		return
	}

	c.String(200, "Contraseña actualizada exitosamente")
}

func GetEjerciciosRealizadosPorUsuario(c *gin.Context) {
	database := configs.ConnectToDB()

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "ID de usuario inválido")
		return
	}

	ejerciciosRealizados, err := services.GetEjerciciosRealizadosPorUsuario(database, uint(userID))
	if err != nil {
		c.String(500, "Error al obtener los ejercicios realizados por el usuario")
		return
	}

	c.JSON(200, ejerciciosRealizados)
}

func GetEjercicio(c *gin.Context) {
	database := configs.ConnectToDB()

	userID, exists := c.Get("userID")
	if !exists {
		c.String(400, "ID de usuario inválido")
		return
	}

	exercise, err := services.GetNextExercise(database, userID.(string))
	if err != nil {
		c.String(500, "Error al obtener el ejercicio")
		return
	}

	ejercicio := models.EjercicioDTO{
		ID:        exercise.ID,
		Nombre:    exercise.Nombre,
		Contenido: exercise.Contenido,
		Nivel:     exercise.NivelID,
		Tipo:      exercise.TipoID,
	}

	c.JSON(200, ejercicio)
}
