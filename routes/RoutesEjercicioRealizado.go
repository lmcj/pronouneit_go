package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
)

func RoutesEjercicioRealizado(e *gin.Engine) {
	ejercicioRealizado := e.Group("/ejercicios-realizados")
	ejercicioRealizado.POST("/crear", controllers.CreateEjercicioRealizado)
	ejercicioRealizado.GET("/usuario/:usuario_id", controllers.GetEjerciciosRealizadosByUsuarioID)
	ejercicioRealizado.GET("/obtener/:id", controllers.GetEjercicioRealizadoById)
	ejercicioRealizado.PUT("/actualizar/:id", controllers.UpdateEjercicioRealizado)
	ejercicioRealizado.DELETE("/eliminar/:id", controllers.DeleteEjercicioRealizado)
}
