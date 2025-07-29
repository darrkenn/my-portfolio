package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Song struct {
	Title   string   `json:"title"`
	Link    string   `json:"link"`
	Artists []string `json:"artists"`
}

func FavouriteSong(c *gin.Context, sLocation string) {
	file, fileErr := os.Open(sLocation)
	if fileErr != nil {
		fmt.Println("Cant open file: ", fileErr)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var song Song
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&song); err != nil {
		fmt.Println("Cant open file: ", err)
	}
	c.HTML(http.StatusOK, "song.gohtml", gin.H{
		"data": song,
	})
}
