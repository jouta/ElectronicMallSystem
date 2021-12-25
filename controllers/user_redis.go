package controllers

import (
	"mall/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Response struct {
	status  bool
	message string
	result  string
}

func (connRedis *ConnRedis) GetUser(ctx *gin.Context) {
	var userID string
	userID = ctx.Query("userid")
	user := models.User{}
	err, userData := user.GetUser(connRedis.DB, userID)
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
		"result": userData,
	})
}

func (connRedis *ConnRedis) AddUser(ctx *gin.Context) {
	var scoreValue int
	scoreValue, err := strconv.Atoi(ctx.Query("score"))
	userName := ctx.Query("name")
	if err != nil {
		// handle error
		ctx.JSON(500, gin.H{"error": "user invalid"})
	}

	user := models.User{}
	user.Id = "user-" + uuid.New().String()
	user.Score = scoreValue
	user.Name = userName
	err = user.Create(connRedis.DB)
	if err == nil {
		if err == nil {
			ctx.JSON(200, gin.H{
				"status": true,
				"result": user,
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
