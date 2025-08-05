package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"my-portfolio/models"
	"net/http"
	"strings"
)

type HTMLProject struct {
	ProjectId uint
	Title     string
	Desc      string
	Techs     []string
	GitLink   *string
	WebLink   *string
	BlogId    *uint
}

func GetProjects(c *gin.Context, db *gorm.DB) {
	var projects []models.Project
	var htmlProjects []HTMLProject

	result := db.Find(&projects)
	if result.Error != nil {
		log.Fatal("Cant get all projects: ", result.Error)
	}

	for _, project := range projects {
		var htmlProject HTMLProject

		htmlProject.ProjectId = project.ProjectId
		htmlProject.Title = project.Title
		htmlProject.Desc = project.Desc
		htmlProject.Techs = strings.Split(project.Tech, "#")
		htmlProject.GitLink = project.GitLink
		htmlProject.WebLink = project.WebLink
		htmlProject.BlogId = project.BlogId
		htmlProjects = append(htmlProjects, htmlProject)
	}

	c.HTML(http.StatusOK, "projects.gohtml", gin.H{
		"Projects": htmlProjects,
	})
}
