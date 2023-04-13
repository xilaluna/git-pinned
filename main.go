package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/joho/godotenv"
)

type Repo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Language    string `json:"language"`
	Stars       string `json:"stars"`
	Forks       string `json:"forks"`
}

func scrape(link string) []Repo {
	var repos []Repo
	collector := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)
	extensions.RandomUserAgent(collector)
	collector.OnHTML(".pinned-item-list-item-content", func(content *colly.HTMLElement) {
		title := content.ChildText("span.repo[title]")
		description := content.ChildText("p.pinned-item-desc")
		repoLink := fmt.Sprintf("https://github.com%s", content.ChildAttr("a", "href"))
		language := content.ChildText("span[itemprop=programmingLanguage]")
		stars := content.ChildText(`a[href$="/stargazers"]`)
		forks := content.ChildText(`a[href$="/forks"]`)

		repos = append(repos, Repo{Title: title, Description: description, Link: repoLink, Language: language, Stars: stars, Forks: forks})
	})
	collector.Visit(link)
	return repos
}

func main() {
	godotenv.Load()
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		repos := scrape("https://github.com/xilaluna")
		context.IndentedJSON(http.StatusOK, repos)
	})
	router.GET("/:id", func(context *gin.Context) {
		link := fmt.Sprintf("https://github.com/%s", context.Param("id"))
		repos := scrape(link)
		context.IndentedJSON(http.StatusOK, repos)
	})
	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
