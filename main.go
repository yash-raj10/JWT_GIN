package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/yash-raj10/JWT_GIN/routes"
	"os"
)

func main() {
	port = os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.New()
	r.Use(gin.Logger())

	routes.AuthRoutes(r)
	routes.UserRouter(r)

	r.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-1"})
	})

	r.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-2"})
	})

	r.Run(":" + port)
}
