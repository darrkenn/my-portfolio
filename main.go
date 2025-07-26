package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.GET("/working", func(c *gin.Context) {
		isWorking := c.Query("status")

		c.JSON(200, gin.H{
			"Am i working?": isWorking,
		})
	})

	g.Run(":1375")
}
