package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
	"github.com/lmcj/pronouneit_go.git/middleware"
)

func RoutesEjercicioRealizado(e *gin.Engine) {
	ejercicioRealizado := e.Group("/ejercicios-realizados")
	ejercicioRealizado.POST("/crear", middleware.AdminAuthMiddleware(), controllers.CreateEjercicioRealizado)
	ejercicioRealizado.GET("/usuario/:usuario_id", middleware.AuthMiddleware(), controllers.GetEjerciciosRealizadosByUsuarioID)
	ejercicioRealizado.GET("/obtener/:id", middleware.AuthMiddleware(), controllers.GetEjercicioRealizadoById)
	ejercicioRealizado.PUT("/actualizar/:id", middleware.AdminAuthMiddleware(), controllers.UpdateEjercicioRealizado)
	ejercicioRealizado.DELETE("/eliminar/:id", middleware.AdminAuthMiddleware(), controllers.DeleteEjercicioRealizado)
}
