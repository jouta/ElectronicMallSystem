package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mall/controllers"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {

	r := gin.Default()
	r.Use(Cors())
	gin.SetMode(viper.GetString("mode"))

	client := r.Group("/user")
	{
		client.GET("/stories", controllers.Show)
		client.GET("/stories/:id", controllers.Read)
		client.POST("/stories", controllers.Create)
		client.PUT("/stories/:id", controllers.Update)
		client.DELETE("/stories/:id", controllers.Delete)
	}

	product := r.Group("/product")
	{
		product.POST("/create", controllers.CreateProduct)
		product.GET("/stories", controllers.Show)
		product.GET("/stories/:id", controllers.Read)
		product.POST("/stories", controllers.Create)
		product.PUT("/stories/:id", controllers.Update)
		product.DELETE("/stories/:id", controllers.Delete)
	}
	//user := r.Group("/api/user")
	//{
	//
	//	user.GET("/info/:id", UserHandler.UserInfoHandler)
	//	user.POST("/add", UserHandler.AddUserHandler)
	//	user.POST("/edit", UserHandler.EditUserHandler)
	//	user.POST("/delete/:id", UserHandler.DeleteUserHandler)
	//}

	//port := viper.GetString("port")

	r.Run(":3000")
}
