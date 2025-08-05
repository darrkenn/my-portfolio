package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"my-portfolio/models"
	"net/http"
)

func GetBlogs(c *gin.Context, db *gorm.DB) {
	var blogs []models.Blog
	result := db.Order("Blog_id desc").Find(&blogs)

	if result.Error != nil {
		log.Fatal("Cant get all blogs: ", result.Error)
	}
	c.HTML(http.StatusOK, "blogs.gohtml", gin.H{
		"blogs": blogs,
	})
}
