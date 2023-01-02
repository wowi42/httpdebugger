package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tpkeeper/gin-dump"
	"io/ioutil"
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
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.JSON(200, gin.H{
			"message": body,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
