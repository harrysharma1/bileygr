package weebcentralscraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/darenliang/jikan-go"
	"github.com/gocolly/colly"
)

func MAL() {
	anime, err := jikan.GetAnimeById(1)
	if err != nil {
		log.Fatalf("error getting anime, %e", err)
	}

	fmt.Println(anime)
}

func WeebCentralScrape() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if strings.Contains(link, "/chapters/") {
			fmt.Printf("Chapter Link found: %s\n", link)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with response: %v\nError: %v", r.Request.URL, r, err)
	})

	err := c.Visit("https://weebcentral.com/series/01J76XYBD282Y3XKX0KRGD64Q5/full-chapter-list")
	if err != nil {
		log.Fatalf("error visiting Lookism: %e", err)
	}

}
