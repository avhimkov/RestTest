package main

import (
	//"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/instructions", GetInstructions)
	}
	return router
}

func GetInstructions(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "GET api/v1/instructions"})
}

func GetInsruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "GET api/v1/instructions/1"})
}

func PostInsruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})
}

func UpdateInsruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})
}

func DeleteInsruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}