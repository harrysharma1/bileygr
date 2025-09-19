package weebcentralscraper

import (
	"fmt"
	"log"

	"github.com/darenliang/jikan-go"
)

func MAL() {
	manhwa, err := jikan.GetTopManga(jikan.TopMangaTypeManhwa, jikan.TopMangaFilterByPopularity, 1)
	if err != nil {
		log.Fatalf("error getting anime, %e", err)
	}
	fmt.Println(manhwa.Data)
}
