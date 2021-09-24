package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

type MarkdownRequest struct {
	RequestId       int64  `json:"requestId"`
	MarkdownContent string `json:"content"`
}

type Test struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

func main() {
	r := gin.Default()

	// authorization

	v1 := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	v1.POST("/markdown", func(c *gin.Context) {

		var markdownReq MarkdownRequest
		err := c.BindJSON(&markdownReq)

		if err != nil {
			log.Println(err)
		}

		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Println(user)

		markdownBytes := []byte(markdownReq.MarkdownContent)

		if !checkMarkdown(markdownBytes) {
			log.Println("Bad markdown")
		}

		htmlBytes := markdownToHtml(markdownBytes)

		fmt.Println("##################################")
		fmt.Print(string(sanatizeHtml(markdownToHtml(htmlBytes))))
		fmt.Println("##################################")

	})

	r.Run()
}

//TODO
func checkMarkdown(markdownBytes []byte) bool {
	return true
}

/**
* Converts markdown bytes to html bytes
**/
func markdownToHtml(markdownBytes []byte) []byte {
	parser := parser.New()
	return markdown.ToHTML(markdownBytes, parser, nil)
}

/**
* Cleans the html
* TODO: look at settings
**/
func sanatizeHtml(harmfullHtml []byte) []byte {
	return bluemonday.UGCPolicy().SanitizeBytes(harmfullHtml)
}
