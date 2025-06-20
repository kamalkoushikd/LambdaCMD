package main

import (
	"fmt"
	// "log"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("Starting Gin server...")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		name := c.DefaultQuery("name", "World")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Welcome, %s!", name),
		})
	})

	r.Run(":8080")
}
