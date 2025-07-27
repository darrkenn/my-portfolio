package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"my-portfolio-api/models"
	"net/http"
)

func CreateProject(c *gin.Context, db *gorm.DB) {
	var newProject models.Project
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&newProject)
	if result.Error != nil {
		log.Fatal("Cant create new project: ", result.Error)
	}
	c.JSON(http.StatusOK, newProject)
}
