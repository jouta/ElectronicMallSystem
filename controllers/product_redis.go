package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mall/models"
	_ "strconv"
)


func (connRedis *ConnRedis) CreateProduct(ctx *gin.Context) {
	type CreateProduct struct {
		ProductName  string `form:"productName" json:"productName" binding:"required"`
		ProductIntro string `form:"productIntro" json:"productIntro"`
		Price    string `form:"price" json:"price" binding:"required"`
		StockNum int    `form:"stockNum" json:"stockNum" binding:"required"`
		ProductImg string `form:"productImg" json:"productImg" binding:"required"`
	}
	var json CreateProduct
	product := models.Product{}
	err := ctx.ShouldBindJSON(&json)
	if err == nil{
		product.ProductId = "product-" + uuid.New().String()
		product.ProductName = json.ProductName
		product.ProductIntro = json.ProductIntro
		product.Price = json.Price
		product.StockNum = json.StockNum
		product.ProductImg = json.ProductImg
	}
	fmt.Println(product)
	err = product.CreateProduct(connRedis.DB)
	if err == nil {
		if err == nil {
			ctx.JSON(200, gin.H{
				"status": true,
				"result": product,
			})
		}
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

func (connRedis *ConnRedis) GetOneProduct(ctx *gin.Context) {
	var productId string
	productId = ctx.Query("productId")
	product := models.Product{}
	err, productData := product.GetProduct(connRedis.DB, productId)
	fmt.Println(productData)
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
		"result": productData,
	})
}