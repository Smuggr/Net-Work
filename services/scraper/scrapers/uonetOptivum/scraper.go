package uonetoptivum

import (
	"fmt"

	"github.com/chromedp/chromedp"
)

type UonetOptivum struct {
	Timestamps []Timestamp
}

func ScrapeUONetOptivum(baseUrl string) {
	fmt.Printf("Scraping %s", baseUrl)
	
}