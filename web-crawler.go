package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Page struct {
    Title string
    Url string
    Links []string
}

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
    c.OnHTML("title", func(e *colly.HTMLElement) {
        siteTitle = e.Text
    })
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        if !strings.Contains(e.Attr("href"), "https://edgewebware.com") {
			return
		}
		e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Start scraping on https://edgewebware.com
	c.Visit("https://edgewebware.com")
    
	fmt.Println("Site Title: ", siteTitle)
}
