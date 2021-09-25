package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

type MarkdownRequest struct {
	RequestId       int64  `json:"requestId"`
	MarkdownContent string `json:"content"`
}

func main() {

	get_password()
	get_username()

	r := gin.Default()

	// authorization

	v1 := r.Group("/", gin.BasicAuth(gin.Accounts{
		get_username(): get_password(),
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

		htmlBytes := markdownToHtml(markdownBytes)

		fmt.Println("##################################")
		fmt.Print(string(sanatizeHtml(markdownToHtml(htmlBytes))))
		fmt.Println("##################################")

	})

	r.Run()
}

/**
* Converts markdown bytes to html bytes
**/
func markdownToHtml(markdownBytes []byte) []byte {
	parser := parser.New()
	return markdown.ToHTML(markdownBytes, parser, nil)
}

/**
* cleans the html of potential malicous content
* bluemonday.UGCPolicy() which allows a broad selection of HTML elements and attributes that are safe for user generated content.
* Note that this policy does not allow iframes, object, embed, styles, script, etc.
* An example usage scenario would be blog post bodies where a variety of formatting is expected along with the potential for TABLEs and IMGs.
**/
func sanatizeHtml(harmfullHtml []byte) []byte {
	return bluemonday.UGCPolicy().SanitizeBytes(harmfullHtml)
}

// Retrives the password from the env varbiable
func get_password() string {
	password := os.Getenv("markdown_service_password")
	if password == "" {
		log.Fatalf("Could not retrieve password")
	}

	return password
}

// Retrives the username from the env varbiable
func get_username() string {
	username := os.Getenv("markdown_service_username")

	if username == "" {
		log.Fatal("Could not retrieve username")
	}
	return username
}
