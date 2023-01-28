package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/PuerkitoBio/goquery"
)

func GetCurrentAssetData(assets []string) []byte {
	results := ""
	for _, item := range assets {
		results = results + GetInvestingData(item) + ","
	}
	if len(results) > 0 {
		// Removes last comma
		results = results[:len(results) - 1]
	}
	results = "["+results+"]"
	return []byte(results)
}

func GetInvestingData(asset string) string {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("br.investing.com"),
	)

	// Extracts asset value
	c.OnHTML(`#__next`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		price := ""
		goquerySelection.Find(`div[data-test=instrument-header-details] span[data-test=instrument-price-last]`).Each (func(index int,item *goquery.Selection) {
			price = item.Text()
		})
		

		body = body + CreateJsonStringField("asset",asset, true)
		body = body + CreateJsonStringField("price",price, false)
		
	})

	c.Visit("https://br.investing.com/equities/"+asset)

	body = "{"+body+"}"

	return body
}