package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Technologies struct {
	Tech []string `json:"tech"`
}

func GetTechnologies(c *gin.Context, tLocation string) {
	file, fileErr := os.Open(tLocation)
	if fileErr != nil {
		fmt.Println("Cant open file: ", fileErr)
	}
	defer file.Close()

	var tech Technologies
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tech); err != nil {
		fmt.Println("Cant decode file: ", err)
	}
	c.HTML(http.StatusOK, "technologies.gohtml", gin.H{
		"tech": tech,
	})
}
