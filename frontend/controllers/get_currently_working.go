package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type CurrentlyWorking struct {
	Project string `json:"project"`
	Link    string `json:"link"`
}

func GetCW(c *gin.Context, cwLocation string) {
	file, fileErr := os.Open(cwLocation)
	if fileErr != nil {

		fmt.Println("Cant open file: ", fileErr)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var cw CurrentlyWorking
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cw); err != nil {
		fmt.Println("Cant open file: ", err)
	}
	c.HTML(http.StatusOK, "currentlyWorking.gohtml", gin.H{
		"data": cw,
	})
}
