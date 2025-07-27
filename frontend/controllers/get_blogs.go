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
	result := db.Find(&blogs)
	if result.Error != nil {
		log.Fatal("Cant get all blogs: ", result.Error)
	}
	c.JSON(http.StatusOK, blogs)
}
