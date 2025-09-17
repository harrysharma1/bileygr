package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func TestScrape() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.webtoons.com"),
	)

	c.OnHTML("input.input_search._txtKeyword", func(h *colly.HTMLElement) {
		placeholder := h.Attr("placeholder")
		ariaLabel := h.Attr("aria-label")

		fmt.Println("Placeholder:", placeholder)
		fmt.Println("Aria-label:", ariaLabel)
	})

}
