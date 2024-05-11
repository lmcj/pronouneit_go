package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
)

func RoutesEjercicio(e *gin.Engine) {
	ejercicio := e.Group("/ejercicios")
	ejercicio.POST("/crear", controllers.CreateEjercicio)
	ejercicio.GET("/listar", controllers.GetEjercicios)
	ejercicio.GET("/obtener/:id", controllers.GetEjercicioById)
	ejercicio.PUT("/actualizar/:id", controllers.UpdateEjercicio)
	ejercicio.DELETE("/eliminar/:id", controllers.DeleteEjercicio)
	ejercicio.GET("/por-nivel/:id", controllers.GetEjerciciosPorNivel)
}
