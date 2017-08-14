package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/instructions", GetInstruction)
	}
	return router
}

func GetInstruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "Welcome to Hell!"})
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}