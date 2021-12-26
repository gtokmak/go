package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/gtokmak/golang-jwt-project/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/singup", controller.Singup())
	incomingRoutes.POST("users/login", controller.Login())
}
