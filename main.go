package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		collector := colly.NewCollector()
		collector.OnHTML(".pinned-item-list-item-content", func(content *colly.HTMLElement) {
			title := content.ChildText("span.repo[title]")
			description := content.ChildText("p.pinned-item-desc")
			fmt.Printf("Pinned repo: %q\n", title)
			fmt.Printf("Description: %q\n", description)
		})
		collector.Visit("https://github.com/xilaluna")
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
