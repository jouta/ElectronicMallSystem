package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mall/models"
)

type Response struct {
	status  bool
	message string
	result  string
}

func (connRedis *ConnRedis) CreateUser(c *gin.Context) {
	json := models.User{}
	json.UserId = "user-" + uuid.New().String()
	if err := c.ShouldBindJSON(&json); err == nil {
		err = json.Create(connRedis.DB)
		if err == nil {
			if err == nil {
				c.JSON(200, gin.H{
					"status": true,
					"result": json,
				})
			}
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

	//defer connRedis.DB.Close()

}

//查询所有用户
func (connRedis *ConnRedis) ShowUser(c *gin.Context) {
	user := models.User{}
	err, userData := user.GetAllUser(connRedis.DB)
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

	type ShowUser struct {
		UserId string  `json:"userId" form:"userId"  binding:"required"`
		UserName string `json:"userName" form:"userName"  binding:"required"`
		Address  string `json:"address" form:"address"  binding:"required"`
		UserType int `json:"userType" form:"userType"  binding:"required"`
	}
	var listUsers []ShowUser

	for _, userdata := range userData {
		users := ShowUser{}
		users.UserId = userdata.UserId
		users.UserName = userdata.UserName
		users.Address = userdata.Address
		users.UserType = userdata.UserType
		listUsers = append(listUsers, users)
	}

	c.JSON(200, gin.H{
		"status": true,
		"result": listUsers,
	})

}

//查询单个用户
func (connRedis *ConnRedis) GetUser(c *gin.Context) {
	var userID string
	userID = c.Query("userid")
	user := models.User{}
	err, userData := user.GetUser(connRedis.DB, userID)
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
	type ShowUser struct {
		UserId string  `json:"userId" form:"userId"  binding:"required"`
		UserName string `json:"userName" form:"userName"  binding:"required"`
		Address  string `json:"address" form:"address"  binding:"required"`
		UserType int `json:"userType" form:"userType"  binding:"required"`
	}

	users := ShowUser{}
	users.UserId = userData.UserId
	users.UserName = userData.UserName
	users.Address = userData.Address
	users.UserType = userData.UserType

	c.JSON(200, gin.H{
		"status": true,
		"result": users,
	})

}




/*
func (connRedis *ConnRedis) GetTop(ctx *gin.Context) {
	var lengthValue int
	lengthValue, err := strconv.Atoi(ctx.Query("length"))

	user := models.User{}
	topData, err := user.GetTop(connRedis.DB, lengthValue)
	if err != nil {
		// handle error
		ctx.JSON(500, gin.H{"error": "topData invalid"})
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"result": topData,
	})
}
*/