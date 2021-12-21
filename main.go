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

	client := r.Group("/storie")
	{
		client.GET("/stories", controllers.Show)
		client.GET("/stories/:id", controllers.Read)
		client.POST("/stories", controllers.Create)
		client.PUT("/stories/:id", controllers.Update)
		client.DELETE("/stories/:id", controllers.Delete)
	}

	admin := r.Group("/admin")
	{
		// 商品增删改
		admin.POST("/createProduct", controllers.CreateProduct)
		admin.POST("/UpdateProduct/:productId", controllers.UpdateProduct)
		admin.DELETE("/DeleteProduct/:productId", controllers.DeleteProduct)
		//用户增删改查
		admin.POST("/CreateUser", controllers.CreateUser)
		admin.GET("/ShowUser", controllers.ShowUser)

	}



	r.Run(":3000")
}
