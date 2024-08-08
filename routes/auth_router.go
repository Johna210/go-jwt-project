package routes

import (
	controller "github.com/Johna210/go-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.Singup())
	incomingRoutes.POST("/users/login", controller.Login())
}
