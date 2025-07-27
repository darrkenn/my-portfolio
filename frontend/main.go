package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

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

	err := r.Run(":1375")
	if err != nil {
		return
	}
}
