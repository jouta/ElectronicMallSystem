package main

import (
	"mall/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

	//redis连接
	redis_controllers := controllers.ConnRedis{}
	redis_controllers.Connect()

	r := gin.Default()
	r.Use(Cors())
	gin.SetMode(viper.GetString("mode"))

	//参考用
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
		// 商品增删改mysql
		admin.POST("/CreateProduct_mysql", controllers.CreateProduct_mysql)
		admin.POST("/UpdateProduct/:productId", controllers.UpdateProduct)
		admin.DELETE("/DeleteProduct/:productId", controllers.DeleteProduct)
		// 商品增删改redis
		admin.POST("/CreateProduct", redis_controllers.CreateProduct)
		admin.GET("/DeleteProduct", redis_controllers.DeleteProduct)
		admin.POST("/UpdateProduct", redis_controllers.UpdateProduct)
		//用户增删改查mysql
		admin.POST("/CreateUser_mysql", controllers.CreateUser)
		admin.GET("/ShowUser_mysql", controllers.ShowUser)
		// 用户增删改redis
		admin.POST("/CreateUser", redis_controllers.CreateUser)
		admin.GET("/ShowUser", redis_controllers.ShowUser)
		admin.GET("/GetOneUser", redis_controllers.GetUser)
		admin.DELETE("/DeleteUser", redis_controllers.DeleteUser)
		admin.POST("/UpdateUser", redis_controllers.UpdateUser)
		//查看所有订单
		admin.GET("/ShowOrder", redis_controllers.ShowOrder)
	}

	common := r.Group("/common")
	{
		//商品列表
		common.GET("/ShowProduct_mysql", controllers.ShowProduct)
		common.GET("/ShowProduct", redis_controllers.ShowProduct)
		common.GET("/GetOneProduct", redis_controllers.GetOneProduct)
		common.GET("/SearchProduct", redis_controllers.SearchProduct)
		common.POST("/Login", redis_controllers.Login)
	}

	user := r.Group("/user")
	{
		//创建订单
		user.POST("/CreateOrder", redis_controllers.CreateOrder)
		//支付订单
		user.POST("/PayOrder", redis_controllers.PayOrder)

	}

	r.Run(":3000")
}
