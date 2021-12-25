package controllers

import (
	"mall/models"

	"github.com/gin-gonic/gin"
)

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
		OrderId     string `json:"orderId" form:"orderId" binding:"required"`
		UserId      string `json:"userId" form:"userId"  binding:"required"`
		ProductId   string `json:"productId" form:"productId" binding:"required"`
		Price       string `json:"price" form:"price" binding:"required"`
		OrderStatus int    `json:"orderStatus" form:"orderStatus" binding:"required"`
		PayTime     string `json:"payTime" form:"payTime" binding:"required"`
		OrderTime   string `json:"orderTime" form:"orderTime" binding:"required"`
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
