package main

import (
	"fmt"
	"log"
	"my-portfolio/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func main() {
	// Setup
	r := gin.Default()
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Cant open env: ", envErr)
	}

	dbLocation := os.Getenv("DATABASE_LOCATION")
	cwLocation := os.Getenv("CW_LOCATION")
	tLocation := os.Getenv("T_LOCATION")
	mdLocation := os.Getenv("MD_LOCATION")
	mLocation := os.Getenv("M_LOCATION")

	db, dbErr := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("Cant open database: ", dbErr)
	}
	fmt.Println("This is the db", db)

	// No route page
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "PAGE_NOT_FOUND", "msg": "PAGE_NOT_FOUND",
		})
	})

	// Html/Static files
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Html page routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
	r.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", nil)
	})
	r.GET("/blogs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blogs.html", nil)
	})
	r.GET("/blogs/:blogId", func(c *gin.Context) {
		blogId := c.Param("blogId")
		c.HTML(http.StatusOK, "blog.gohtml", gin.H{
			"ID": blogId,
		})
	})
	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	// Reusable HTML
	r.GET("/api/navbar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "navbar.html", nil)
	})
	// Database routes
	r.GET("/api/getProjects", func(c *gin.Context) {
		controllers.GetProjects(c, db)
	})
	r.GET("/api/getBlogs", func(c *gin.Context) {
		controllers.GetBlogs(c, db)
	})
	r.GET("/api/blog/:blogId", func(c *gin.Context) {
		blogId := c.Param("blogId")
		controllers.RenderBlog(c, blogId, mdLocation)
	})
	// About page
	r.GET("/api/getCurrentlyWorking", func(c *gin.Context) {
		controllers.GetCW(c, cwLocation)
	})
	r.GET("/api/technologies", func(c *gin.Context) {
		controllers.GetTechnologies(c, tLocation)
	})
	r.POST("/api/message", func(c *gin.Context) {
		controllers.NewMessage(c, mLocation)
	})

	runErr := r.Run(":1375")
	if runErr != nil {
		return
	}
}
