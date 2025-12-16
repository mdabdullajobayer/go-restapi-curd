package main

import (
	"github.com/curd/project/controller"
	"github.com/curd/project/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	// create data
	router.POST("/create", controller.CreatePost)
	// get all data
	router.GET("/", controller.GetPosts)
	// Find by Id
	router.GET("/:id", controller.GetPostByID)
	// update data
	router.PUT("/:id", controller.UpdatePost)
	// delete data
	router.DELETE("/:id", controller.DeletePost)
	router.Run() // listens on 0.0.0.0:8080 by default
}
