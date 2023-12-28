package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/yash-raj10/JWT_GIN/controllers"
	"github.com/yash-raj10/JWT_GIN/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}
