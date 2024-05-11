package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
)

func RoutesUsuario(e *gin.Engine) {
	usuario := e.Group("/usuarios")
	usuario.POST("/crear", controllers.CreateUsuario)
	usuario.GET("/listar", controllers.GetUsuarios)
	usuario.GET("/obtener/:id", controllers.GetUsuarioById)
	usuario.PUT("/actualizar/:id", controllers.UpdateUsuario)
	usuario.DELETE("/eliminar/:id", controllers.DeleteUsuario)
	usuario.POST("/:id/cambiar-contrasenia", controllers.CambioContrasenia)
	usuario.GET("/obtener/:id/ejercicios-realizados", controllers.GetEjerciciosRealizadosPorUsuario)
}
