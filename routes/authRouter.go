package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/yash-raj10/JWT_GIN/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup", controller.SignUp())
	incomingRoutes.POST("user/login", controller.Login())
}
