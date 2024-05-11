package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
)

func RoutesTipoPalabra(e *gin.Engine) {
	tipoPalabra := e.Group("/tipos-palabra")
	tipoPalabra.POST("/crear", controllers.CreateTipoPalabra)
	tipoPalabra.GET("/listar", controllers.GetTiposPalabra)
	tipoPalabra.GET("/obtener/:id", controllers.GetTipoPalabraById)
	tipoPalabra.PUT("/actualizar/:id", controllers.UpdateTipoPalabra)
	tipoPalabra.DELETE("/eliminar/:id", controllers.DeleteTipoPalabra)
}
