package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"net/http"
	"os"
	"path"
)

func RenderBlog(c *gin.Context, id string, mdLocation string) {

	mdFile := path.Join(mdLocation + id + ".md")
	content, readErr := os.ReadFile(mdFile)
	if readErr != nil {
		fmt.Println("Cant read file: ", readErr)
		return
	}

	p := parser.NewWithExtensions(parser.CommonExtensions)
	renderer := html.NewRenderer(html.RendererOptions{
		Flags: html.CommonFlags,
	})
	htmlContent := markdown.ToHTML(content, p, renderer)

	c.Header("Content-type", "text/html")
	c.String(http.StatusOK, string(htmlContent))
}
