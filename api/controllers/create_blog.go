package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"my-portfolio-api/models"
	"net/http"
)

func CreateBlog(c *gin.Context, db *gorm.DB) {
	var newBlog = models.Blog{}
	if err := c.ShouldBindJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&newBlog)
	if result.Error != nil {
		log.Fatal("Cant create new blog: ", result.Error)
	}
	c.JSON(http.StatusOK, newBlog)
}
