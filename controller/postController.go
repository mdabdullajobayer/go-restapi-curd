package controller

import (
	"github.com/curd/project/initializers"
	"github.com/curd/project/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}
	if err := initializers.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": post})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := initializers.DB.Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": posts})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := initializers.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: input.Title, Content: input.Content})

	c.JSON(200, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := initializers.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{"data": true})
}

func GetPostByID(c *gin.Context) {
	var post models.Post
	if err := initializers.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, gin.H{"data": post})
}
