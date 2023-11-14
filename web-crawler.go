package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
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
		if !strings.Contains(e.Attr("href"), "edgewebware.com") {
			return
		}
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Title found: ", e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	fmt.Println("Site Title: ", siteTitle)

	// Start scraping on https://edgewebware.com
	c.Visit("https://edgewebware.com")
}
