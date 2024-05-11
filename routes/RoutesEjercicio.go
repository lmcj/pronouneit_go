package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
	"github.com/lmcj/pronouneit_go.git/middleware"
)

func RoutesEjercicio(e *gin.Engine) {
	ejercicio := e.Group("/ejercicios")
	ejercicio.POST("/crear", middleware.AdminAuthMiddleware(), controllers.CreateEjercicio)
	ejercicio.GET("/listar", middleware.AdminAuthMiddleware(), controllers.GetEjercicios)
	ejercicio.GET("/obtener/:id", middleware.AuthMiddleware(), controllers.GetEjercicioById)
	ejercicio.PUT("/actualizar/:id", middleware.AdminAuthMiddleware(), controllers.UpdateEjercicio)
	ejercicio.DELETE("/eliminar/:id", middleware.AdminAuthMiddleware(), controllers.DeleteEjercicio)
	ejercicio.GET("/por-nivel/:id", middleware.AdminAuthMiddleware(), controllers.GetEjerciciosPorNivel)
}
