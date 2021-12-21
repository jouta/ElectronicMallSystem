package controllers

import (
	"github.com/gin-gonic/gin"
	"mall/database"
)

type Product struct {
	ProductId   int    `json:"productId"`
	ProductName  string `json:"productName"`
	ProductIntro string `json:"productIntro"`
	Prize    string `json:"prize"`
	StockNum int    `json:"stockNum"`
}

func CreateProduct(c * gin.Context) {
	db := database.DBConn()

	type CreateProduct struct {
		ProductName  string `form:"productName" json:"productName" binding:"required"`
		ProductIntro string `form:"productIntro" json:"productIntro"`
		Prize    string `form:"prize" json:"prize" binding:"required"`
		StockNum int    `form:"stockNum" json:"stockNum" binding:"required"`
	}

	var json CreateProduct

	if err := c.ShouldBindJSON(&json); err == nil {
		insProduct, err := db.Prepare("INSERT INTO product(productName, productIntro, prize, stockNum) VALUE(?,?,?,?)")
		if err != nil {
			c.JSON(500, gin.H {
				"message": err,
			})
		}

		insProduct.Exec(json.ProductName, json.ProductIntro, json.Prize, json.StockNum)
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