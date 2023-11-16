package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

type Page struct {
	Title string
	Url   string
	Links []string
}

func main() {
	visited := make(map[string]bool)

	c := colly.NewCollector()
	var siteTitle string

	// Register Hooks
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if match, err := regexp.MatchString("^/docs.*", e.Attr("href")); err != nil || !match {
			return
		}

		if visited[e.Attr("href")] {
			return
		}

		visited[e.Attr("href")] = true
		fmt.Println(e.Attr("href"))
		e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Start scraping on https://edgewebware.com
	visited["https://keystonejs.com/docs"] = true
	c.Visit("https://keystonejs.com/docs")

	fmt.Println("Site Title: ", siteTitle)
}
