package main

import (
	"github.com/curd/project/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test api is worddking!",
			"status":  "success",
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
