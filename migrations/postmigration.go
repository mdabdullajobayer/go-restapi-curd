package main

import (
	"github.com/curd/project/initializers"
	"github.com/curd/project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
