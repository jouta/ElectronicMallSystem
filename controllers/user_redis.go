package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mall/models"
)


func (connRedis *ConnRedis) CreateUser(c *gin.Context) {
	json := models.User{}
	json.UserId = "user-" + uuid.New().String()
	if err := c.ShouldBindJSON(&json); err == nil {
		err = json.Create(connRedis.DB)
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

//删除用户
func (connRedis *ConnRedis) DeleteUser(ctx *gin.Context) {
	var userId string
	userId = ctx.Query("userid")
	err := models.DeleteUser(connRedis.DB, userId)
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
		"result": userId,
	})
}

func (connRedis *ConnRedis) UpdateUser(c *gin.Context) {
	userID := c.Query("userid")
	//先查询，获取原来的值放在userData中
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

	//绑定从前端查到的值json，空的值就保持原来的值
	json := models.User{}
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
	json.UserId = userData.UserId
	if json.UserName == "" {
		json.UserName = userData.UserName
	}
	if json.PassWord == "" {
		json.PassWord = userData.PassWord
	}
	if json.Address == "" {
		json.Address = userData.Address
	}
	if json.UserType == 0 {
		json.UserType = userData.UserType
	}

	err = json.UpdateUser(connRedis.DB, userID)
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

//用户登录
func (connRedis *ConnRedis) Login(c *gin.Context) {
	json := models.User{}
	err := c.ShouldBindJSON(&json)
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
	flag := false
	for _, userdata := range userData {
		if userdata.UserName == json.UserName && userdata.PassWord == json.PassWord && userdata.UserType == json.UserType{
			c.JSON(200, gin.H{
				"status": true,
				"result": userdata,
			})
			flag = true
			break
		}
	}
	if !flag {
		c.JSON(500, gin.H{
			"status": false,
			"message": "用户名或密码不正确，或者用户类型选择错误",
		})
	}

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