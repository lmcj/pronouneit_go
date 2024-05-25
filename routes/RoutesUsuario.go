package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
	"github.com/lmcj/pronouneit_go.git/middleware"
)

func RoutesUsuario(e *gin.Engine) {
	usuario := e.Group("/usuarios")
	usuario.POST("/crear", controllers.CreateUsuario)
	usuario.GET("/listar", middleware.AdminAuthMiddleware(), controllers.GetUsuarios)
	usuario.GET("/obtener/:id", middleware.AuthMiddleware(), controllers.GetUsuarioById)
	usuario.PUT("/actualizar/:id", middleware.AuthMiddleware(), controllers.UpdateUsuario)
	usuario.DELETE("/eliminar/:id", middleware.AdminAuthMiddleware(), controllers.DeleteUsuario)
	usuario.POST("/:id/cambiar-contrasenia", middleware.AuthMiddleware(), controllers.CambioContrasenia)
	usuario.GET("/obtener/:id/ejercicios-realizados", middleware.AuthMiddleware(), controllers.GetEjerciciosRealizadosPorUsuario)
	usuario.POST("/crear-admin", controllers.CreateAdmin)
	usuario.GET("/ejercicio", middleware.AuthMiddleware(), controllers.GetEjercicio)
}
