package controllers

import (
	"mall/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (connRedis *ConnRedis) CreateOrder(ctx *gin.Context) {
	type CreateOrder struct {
		UserId     string `form:"userId" json:"userId" binding:"required"`
		ProductId  string `form:"productId" json:"productId" binding:"required"`
		ProductNum int    `form:"productNum" json:"productNum" binding:"required"`
		Remark     string `form:"remark" json:"remark" binding:"required"`
	}

	var json CreateOrder
	order := models.Order{}
	err := ctx.ShouldBindJSON(&json)
	if err == nil {
		order.OrderId = "order-" + uuid.New().String()
		order.UserId = json.UserId
		order.ProductId = json.ProductId
		order.OrderStatus = 1
		order.Remark = json.Remark
		order.ProductNum = json.ProductNum

		product := models.Product{}
		err1, productData := product.GetProduct(connRedis.DB, json.ProductId)
		if err1 == nil {
			order.Price = productData.Price * float64(order.ProductNum)
		}

		timeUnix := time.Now().Unix() //已知的时间戳
		timeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
		order.OrderTime = timeStr
		order.PayTime = "" //空表示未支付

	}

	err = order.CreateOrder(connRedis.DB)
	if err == nil {
		ctx.JSON(200, gin.H{
			"status": true,
			"result": order,
		})
	} else {
		resData := &Response{
			status:  false,
			message: err.Error(),
		}
		ctx.JSON(500, gin.H{
			"status":  resData.status,
			"message": resData.message,
		})
	}
}

//查询所有订单
func (connRedis *ConnRedis) ShowOrder(c *gin.Context) {
	order := models.Order{}
	err, orderData := order.GetAllOrder(connRedis.DB)
	if err != nil {
		resData := &Response{
			status:  false,
			message: err.Error(),
		}
		c.JSON(500, gin.H{
			"status":  resData.status,
			"message": resData.message,
		})
		return
	}

	type ShowOrder struct {
		OrderId     string  `json:"orderId" form:"orderId" binding:"required"`
		UserId      string  `json:"userId" form:"userId"  binding:"required"`
		ProductId   string  `json:"productId" form:"productId" binding:"required"`
		Price       float64 `json:"price" form:"price" binding:"required"`
		OrderStatus int     `json:"orderStatus" form:"orderStatus" binding:"required"`
		PayTime     string  `json:"payTime" form:"payTime" binding:"required"`
		OrderTime   string  `json:"orderTime" form:"orderTime" binding:"required"`
		ProductNum  int     `json:"productnum" form:"productNum" binding:"required"`
	}
	var listOrders []ShowOrder

	for _, orderdata := range orderData {
		orders := ShowOrder{}
		orders.OrderId = orderdata.OrderId
		orders.UserId = orderdata.UserId
		orders.ProductId = orderdata.ProductId
		orders.Price = orderdata.Price
		orders.OrderStatus = orderdata.OrderStatus
		orders.PayTime = orderdata.PayTime
		orders.OrderTime = orderdata.OrderTime
		order.ProductNum = orderdata.ProductNum
		listOrders = append(listOrders, orders)
	}

	c.JSON(200, gin.H{
		"status": true,
		"result": listOrders,
	})

}

func (connRedis *ConnRedis) GetOneOrder(ctx *gin.Context) {
	var orderId string
	orderId = ctx.Query("orderId")
	order := models.Order{}
	err, orderData := order.GetOrder(connRedis.DB, orderId)
	orderData.OrderId = orderId
	if err != nil {
		resData := &Response{
			status:  false,
			message: err.Error(),
		}
		ctx.JSON(500, gin.H{
			"status":  resData.status,
			"message": resData.message,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status": true,
		"result": orderData,
	})
}

func (connRedis *ConnRedis) DeleteOrder(ctx *gin.Context) {
	var orderId string
	orderId = ctx.Query("orderId")
	err := models.DeleteOrder(connRedis.DB, orderId)
	if err != nil {
		resData := &Response{
			status:  false,
			message: err.Error(),
		}
		ctx.JSON(500, gin.H{
			"status":  resData.status,
			"message": resData.message,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status": true,
		"result": orderId,
	})
}

func (connRedis *ConnRedis) PayOrder(c *gin.Context) {
	orderID := c.Query("orderid")
	//先查询，获取原来的值放在orderData中
	order := models.Order{}
	err, orderData := order.GetOrder(connRedis.DB, orderID)
	if err != nil {
		resData := &Response{
			status:  false,
			message: err.Error(),
		}
		c.JSON(500, gin.H{
			"status":  resData.status,
			"message": resData.message,
		})
		return
	}

	//将当前时间赋给orderData.PayTime
	timeUnix := time.Now().Unix() //已知的时间戳
	timeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	orderData.PayTime = timeStr
	orderData.OrderStatus = 0

	err = orderData.PayOrder(connRedis.DB, orderID)
	if err == nil {
		c.JSON(200, gin.H{
			"status": true,
			"result": orderData,
		})
	} else {
		resData := &Response{
			status:  false,
			message: err.Error(),
		}
		c.JSON(500, gin.H{
			"status":  resData.status,
			"message": resData.message,
		})
	}

}
