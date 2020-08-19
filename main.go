package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"dev": "go",
			"message": "Hello World!",
		})
	})

	r.Run()
}
