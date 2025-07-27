package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	_ "modernc.org/sqlite"
)

func main() {
	//Setup
	r := gin.Default()
	db, dbErr := sql.Open("sqlite", "db/portfolio.db")
	if dbErr != nil {
		log.Fatal("Cant open database: ", dbErr)
	}
	defer db.Close()

	r.POST("/addProject", func(c *gin.Context) {
		CreateProject(c, db)
	})
	r.POST("/addBlog", func(c *gin.Context) {

	})

	err := r.Run(":5731")
	if err != nil {
		return
	}
}
