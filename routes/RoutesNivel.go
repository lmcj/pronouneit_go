package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
)

func RoutesNivel(e *gin.Engine) {
	nivel := e.Group("/niveles")
	nivel.POST("/crear", controllers.CreateNivel)
	nivel.GET("/listar", controllers.GetNiveles)
	nivel.GET("/obtener/:id", controllers.GetNivelById)
	nivel.PUT("/actualizar/:id", controllers.UpdateNivel)
	nivel.DELETE("/eliminar/:id", controllers.DeleteNivel)
}
