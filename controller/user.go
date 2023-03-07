package controller

import "github.com/gin-gonic/gin"

func RegisterUser(c *gin.Context) {
	c.JSON(200, gin.H{"ping": "pong"})
}
