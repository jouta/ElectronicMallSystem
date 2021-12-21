package main

import (
	"github.com/gin-gonic/gin"
	"mall/controllers"
)


func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api/v1")
	{
		client.GET("/stories", controllers.Show)
		client.GET("/stories/:id", controllers.Read)
		client.POST("/stories", controllers.Create)
		client.PUT("/stories/:id", controllers.Update)
		client.DELETE("/stories/:id", controllers.Delete)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":3000")
}
