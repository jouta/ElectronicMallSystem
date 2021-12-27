package clusterControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	models "mall/clusterModels"
	//"mall/models"
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