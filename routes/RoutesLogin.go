package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcj/pronouneit_go.git/controllers"
)

func RoutesLogin(e *gin.Engine) {
	login := e.Group("/auth")
	login.POST("/login", controllers.Login)

}
