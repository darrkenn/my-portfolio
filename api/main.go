package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"my-portfolio-api/controllers"
	"net/http"
	"strconv"
)

func main() {
	//Setup
	r := gin.Default()
	db, dbErr := gorm.Open(sqlite.Open("db/portfolio.db"), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("Cant open database: ", dbErr)
	}

	r.POST("/addProject", func(c *gin.Context) {
		controllers.CreateProject(c, db)
	})
	r.POST("/addBlog", func(c *gin.Context) {
		controllers.CreateBlog(c, db)
	})
	r.GET("/deleteBlog/:id", func(c *gin.Context) {
		stringId := c.Param("id")
		id, idErr := strconv.Atoi(stringId)
		if idErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": idErr,
			})
		}
		controllers.DeleteBlog(c, id, db)
	})
	r.GET("/deleteProject/:id", func(c *gin.Context) {
		stringId := c.Param("id")
		id, idErr := strconv.Atoi(stringId)
		if idErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": idErr,
			})
		}
		controllers.DeleteProject(c, id, db)
	})

	err := r.Run(":5731")
	if err != nil {
		return
	}
}
