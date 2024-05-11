package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Goooo.",
		})
	})

	RoutesUsuario(router)
	RoutesTipoPalabra(router)
	RoutesNivel(router)
	RoutesEjercicio(router)
	RoutesEjercicioRealizado(router)
	RoutesLogin(router)

	router.Run(":8080")
}
