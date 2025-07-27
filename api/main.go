package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"my-portfolio-api/controllers"
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

	err := r.Run(":5731")
	if err != nil {
		return
	}
}
