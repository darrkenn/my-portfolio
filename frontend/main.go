package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
)

func main() {
	//Setup
	r := gin.Default()
	db, dbErr := sql.Open("sqlite", "db/portfolio.db")
	if dbErr != nil {
		log.Fatal("Cant open database: ", dbErr)
	}
	defer db.Close()

	//No route page
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "PAGE_NOT_FOUND", "msg": "PAGE_NOT_FOUND",
		})
	})

	// Html/Static file LoadHTMLGlob
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Html page routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", gin.H{})
	})
	r.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", gin.H{})
	})
	r.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", gin.H{})
	})

	//Database routes
	r.GET("/api/getProjects", func(c *gin.Context) {
	})
	r.GET("/api/getBlogs", func(c *gin.Context) {

	})

	runErr := r.Run(":1375")
	if runErr != nil {
		return
	}
}
