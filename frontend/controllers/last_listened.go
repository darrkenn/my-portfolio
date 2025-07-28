package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Song struct {
	Title  string `json:"song"`
	Link   string `json:"link"`
	Artist string `json:"artist"`
}

func LastListened(c *gin.Context, sLocation string) {
	file, fileErr := os.Open(sLocation)
	if fileErr != nil {
		fmt.Println("Cant open file: ", fileErr)
	}
	defer file.Close()

	var song Song
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&song); err != nil {
		fmt.Println("Cant decode file: ", err)
	}

	fmt.Println("THIS IS WORKING: ", song)

	c.HTML(http.StatusOK, "song.gohtml", gin.H{
		"song": song,
	})

}
