package main

import (
	//"fmt"
	//"strconv"
	"./app"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/instructions", GetInstructions)
		v1.GET("/instructions/:id", GetInstruction)
		v1.POST("/instructions", PostInstruction)
		v1.PUT("/instructions/:id", UpdateInstruction)
		v1.DELETE("/instructions/:id", DeleteInstruction)
	}
	return router
}

func GetInstructions(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "GET api/v1/instructions"})
}

func GetInstruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "GET api/v1/instructions/1"})
}

func PostInstruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})
}

func UpdateInstruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})
}

func DeleteInstruction(c *gin.Context)  {
	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}