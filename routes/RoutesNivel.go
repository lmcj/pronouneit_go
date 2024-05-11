package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
	"github.com/lmcj/pronouneit_go.git/middleware"
)

func RoutesNivel(e *gin.Engine) {
	nivel := e.Group("/niveles")
	nivel.POST("/crear", middleware.AdminAuthMiddleware(), controllers.CreateNivel)
	nivel.GET("/listar", middleware.AdminAuthMiddleware(), controllers.GetNiveles)
	nivel.GET("/obtener/:id", middleware.AdminAuthMiddleware(), controllers.GetNivelById)
	nivel.PUT("/actualizar/:id", middleware.AdminAuthMiddleware(), controllers.UpdateNivel)
	nivel.DELETE("/eliminar/:id", middleware.AdminAuthMiddleware(), controllers.DeleteNivel)
	nivel.GET("/nivel/maximo", middleware.AuthMiddleware(), controllers.GetNivelMaximo)

}
