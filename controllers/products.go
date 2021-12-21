package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/database"
)

type Product struct {
	ProductId   int    `json:"productId"`
	ProductName  string `json:"productName"`
	ProductIntro string `json:"productIntro"`
	Price    string `json:"price"`
	StockNum int    `json:"stockNum"`
}

func CreateProduct(c * gin.Context) {
	db := database.DBConn()
	type CreateProduct struct {
		ProductName  string `form:"productName" json:"productName" binding:"required"`
		ProductIntro string `form:"productIntro" json:"productIntro"`
		Price    string `form:"price" json:"price" binding:"required"`
		StockNum int    `form:"stockNum" json:"stockNum" binding:"required"`
	}

	var json CreateProduct

	if err := c.ShouldBindJSON(&json); err == nil {
		insProduct, err := db.Prepare("INSERT INTO product(productName, productIntro, price, stockNum) VALUE(?,?,?,?)")
		if err != nil {
			c.JSON(500, gin.H {
				"message": err,
			})
		}

		insProduct.Exec(json.ProductName, json.ProductIntro, json.Price, json.StockNum)
		c.JSON(200, gin.H {
			"message": "inserted",
		})
	} else {
		c.JSON(500, gin.H {
			"error": err.Error(),
		})
	}

	defer db.Close()
}


func UpdateProduct(c * gin.Context) {
	db := database.DBConn()
	type UpdateStory struct {
		//Title string `form:"title" json:"title" binding:"required"`
		//Body string `form:"body" json:"body" binding:"required"`
		ProductName  string `form:"productName" json:"productName" `
		ProductIntro string `form:"productIntro" json:"productIntro"`
		Price    string `form:"price" json:"price" `
		StockNum int    `form:"stockNum" json:"stockNum" `
	}

	var json UpdateStory
	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println(c.Param("productId"))
		edit, err := db.Prepare("UPDATE product SET productName = ?, productIntro = ?, price = ?, stockNum = ? WHERE productId = " + c.Param("productId"))
		if err != nil {
			panic(err.Error())
		}
		edit.Exec(json.ProductName, json.ProductIntro, json.Price, json.StockNum)

		c.JSON(200, gin.H {
			"message": "edited",
		})
	} else {
		c.JSON(500, gin.H {
			"error": err.Error(),
		})
	}
	defer db.Close()
}

