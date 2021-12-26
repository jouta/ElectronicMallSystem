package controllers

import (

	"fmt"
	"github.com/Chain-Zhang/pinyin"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mall/models"
	"strings"
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
		product.ProductId = "product-" +  uuid.New().String()
		product.ProductName = json.ProductName
		product.ProductIntro = json.ProductIntro
		product.Price = json.Price
		product.StockNum = json.StockNum
		product.ProductImg = json.ProductImg
	}
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
	productData.ProductId = productId
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

func (connRedis *ConnRedis) DeleteProduct(ctx *gin.Context) {
	var productId string
	productId = ctx.Query("productId")
	err := models.DeleteProduct(connRedis.DB, productId)
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
		"result": productId,
	})
}

func (connRedis *ConnRedis) ShowProduct(c *gin.Context) {
	product := models.Product{}
	err, productData := product.GetAllProduct(connRedis.DB)
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


	var listProducts []models.Product

	for _, productdata := range productData {
		products := models.Product{}
		products.ProductId = productdata.ProductId
		products.ProductName = productdata.ProductName
		products.ProductIntro = productdata.ProductIntro
		products.Price = productdata.Price
		products.StockNum = productdata.StockNum
		products.ProductImg = productdata.ProductImg
		listProducts = append(listProducts, products)
	}

	c.JSON(200, gin.H{
		"status": true,
		"result": listProducts,
	})

}


func (connRedis *ConnRedis) UpdateProduct(c *gin.Context) {
	productID := c.Query("productid")
	//先查询，获取原来的值放在productData中
	product := models.Product{}
	err, productData := product.GetProduct(connRedis.DB, productID)

	//绑定从前端查到的值json，空的值就保持原来的值
	json := models.Product{}
	err = c.ShouldBindJSON(&json)
	if err != nil {
		resData := &Response{
			status:  false,
			message: err.Error(),
		}
		c.JSON(500, gin.H{
			"status":  resData.status,
			"message": resData.message,
		})
	}
	json.ProductId = productData.ProductId
	if json.ProductName == "" {
		json.ProductName = productData.ProductName
	}
	if json.ProductIntro == "" {
		json.ProductIntro = productData.ProductIntro
	}
	if json.Price == "" {
		json.Price = productData.Price
	}
	if json.StockNum == 0 {
		json.StockNum = productData.StockNum
	}
	if json.ProductImg == ""{
		json.ProductImg = productData.ProductImg
	}

	err = json.UpdateProduct(connRedis.DB, productID)
	if err == nil {
		c.JSON(200, gin.H{
			"status": true,
			"result": json,
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


func (connRedis *ConnRedis) SearchProduct(c *gin.Context) {
	var keyWord string
	keyWord = c.Query("keyWord")
	product := models.Product{}
	err, productData := product.GetAllProduct(connRedis.DB)

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

	var listProducts []models.Product
	str1, err := pinyin.New(keyWord).Convert()
	if err != nil {
		fmt.Println(err)
	}
	for _, productdata := range productData {
		products := models.Product{}
		str2, err := pinyin.New(productdata.ProductName).Convert()
		if err != nil {
			fmt.Println(err)
		}
		reg := strings.Contains(str2, str1)
		if(reg == true){
			products.ProductId = productdata.ProductId
			products.ProductName = productdata.ProductName
			products.ProductIntro = productdata.ProductIntro
			products.Price = productdata.Price
			products.StockNum = productdata.StockNum
			products.ProductImg = productdata.ProductImg
			listProducts = append(listProducts, products)
		}
	}

	c.JSON(200, gin.H{
		"status": true,
		"result": listProducts,
	})

}

