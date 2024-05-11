package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
	"github.com/lmcj/pronouneit_go.git/middleware"
)

func RoutesTipoPalabra(e *gin.Engine) {
	tipoPalabra := e.Group("/tipos-palabra")
	tipoPalabra.POST("/crear", middleware.AdminAuthMiddleware(), controllers.CreateTipoPalabra)
	tipoPalabra.GET("/listar", middleware.AdminAuthMiddleware(), controllers.GetTiposPalabra)
	tipoPalabra.GET("/obtener/:id", middleware.AdminAuthMiddleware(), controllers.GetTipoPalabraById)
	tipoPalabra.PUT("/actualizar/:id", middleware.AdminAuthMiddleware(), controllers.UpdateTipoPalabra)
	tipoPalabra.DELETE("/eliminar/:id", middleware.AdminAuthMiddleware(), controllers.DeleteTipoPalabra)
	tipoPalabra.POST("/crear-tipos-palabras", controllers.CreateTiposPalabras)

}
