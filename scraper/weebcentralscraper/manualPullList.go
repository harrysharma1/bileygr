package weebcentralscraper

type WeebCentralScrapeInfo struct {
	name        string
	seriesUrl   string
	chapterUrls []string
	updateTime  string
}

var ManualWeebCentralList = []WeebCentralScrapeInfo{}
