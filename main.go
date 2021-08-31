package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println()
	r := gin.Default()
	r.GET("/hello-there", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": sayHello(),
		})
	})
	r.Run()
}
