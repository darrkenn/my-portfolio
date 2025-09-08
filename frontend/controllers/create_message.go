package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewMessage(c *gin.Context, mLocation string) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")

	data := map[string]interface{}{
		"name":    name,
		"email":   email,
		"message": message,
	}

	jsonData, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		fmt.Println("Cant create json", jsonErr)
		c.HTML(http.StatusOK, "messageNotReceived.html", nil)
		return
	}

	newUUid := uuid.New()

	filename := fmt.Sprintf("%s[]%s.json", email, newUUid)
	file := filepath.Join(mLocation, filename)

	jsonFile, createErr := os.Create(file)
	if createErr != nil {
		fmt.Println("Cant create json", createErr)
		c.HTML(http.StatusOK, "messageNotReceived.html", nil)
		return
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
		}
	}(jsonFile)

	_, writeErr := jsonFile.Write(jsonData)
	if writeErr != nil {
		fmt.Println("Cant write to file", writeErr)
		c.HTML(http.StatusOK, "messageNotReceived.html", nil)
		return
	}
	c.HTML(http.StatusOK, "messageReceived.html", nil)
}
