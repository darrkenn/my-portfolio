package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Project struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Tech    string `json:"Tech"`
	GitLink string `json:"GitLink"`
	WebLink string `json:"WebLink"`
	BlogId  int    `json:"BlogId"`
}

func CreateProject(c *gin.Context, db *sql.DB) {
	var newProject Project
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println("NEW PROJECT", newProject)
	c.JSON(http.StatusOK, newProject)
}
