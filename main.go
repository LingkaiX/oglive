package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", pingHundle)
	r.Run()
}

func pingHundle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong!",
	})
}
