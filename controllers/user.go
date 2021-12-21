package controllers

import (
	"github.com/gin-gonic/gin"
	"mall/database"
)

type User struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Address  string    `json:"address"`
	UserType   int    `json:"userType"`
}

func CreateUser(c * gin.Context) {
	db := database.DBConn()

	type CreateProduct struct {
		UserName string `json:"userName" form:"userName"  binding:"required"`
		PassWord string `json:"passWord" form:"passWord"  binding:"required"`
		Address  string `json:"address" form:"address"  binding:"required"`
		UserType int `json:"userType" form:"userType"  binding:"required"`
	}

	var json CreateProduct

	if err := c.ShouldBindJSON(&json); err == nil {
		insUser, err := db.Prepare("INSERT INTO user(UserName, PassWord, Address, userType) VALUE(?,?,?,?)")
		if err != nil {
			c.JSON(500, gin.H {
				"message": err,
			})
		}

		insUser.Exec(json.UserName, json.PassWord, json.Address, json.UserType)
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

func ShowUser(c * gin.Context) {
	db := database.DBConn()

	rows, err := db.Query("SELECT userId, userName, address FROM user ")
	if err != nil {
		c.JSON(500, gin.H {
			"message": err.Error(),
		})
	}

	var listUsers [] User

	for rows.Next() {
		var userId int
		var userName, address string
		users := User{}

		err = rows.Scan(&userId, &userName, &address)
		if err != nil {
			panic(err.Error())
		}

		users.UserId = userId
		users.UserName = userName
		users.Address = address

		listUsers = append(listUsers, users)
	}

	c.JSON(200, listUsers)
	defer db.Close()
}