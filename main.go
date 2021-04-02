package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tpkeeper/gin-dump"
)

func main() {
	r := gin.Default()
	r.Use(gindump.Dump())
	r.GET("/*log", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/*log", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
