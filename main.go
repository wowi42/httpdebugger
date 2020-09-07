package main

import "github.com/gin-gonic/gin"
import "io/ioutil"

func main() {
	r := gin.Default()
	r.GET("/*log", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		println(string(body))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/*log", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		println(string(body))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
