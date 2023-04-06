package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

type Repo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Link        string `json:"link"`
}

func main() {
	godotenv.Load()
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		var repos []Repo

		collector := colly.NewCollector()
		collector.OnHTML(".pinned-item-list-item-content", func(content *colly.HTMLElement) {

			title := content.ChildText("span.repo[title]")

			description := content.ChildText("p.pinned-item-desc")
			language := content.ChildText("span[itemprop=programmingLanguage]")
			repoLink := fmt.Sprintf("https://github.com%s", content.ChildAttr("a", "href"))

			repos = append(repos, Repo{Title: title, Description: description, Language: language, Link: repoLink})
		})
		collector.Visit("https://github.com/xilaluna")
		context.IndentedJSON(http.StatusOK, repos)
	})

	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
