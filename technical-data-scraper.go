package main

import (
	"github.com/gocolly/colly/v2"
	"strings"
)

func GetInvestingExchangeRate(fromCurrency string, toCurrency string) []byte {
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
		price = goquerySelection.Find(`.text-base`).Text()
		

		body = body + CreateJsonStringField("from",fromCurrency, true)
		body = body + CreateJsonStringField("to",toCurrency, true)
		body = body + CreateJsonStringField("price",price, false)
		
	})

	c.Visit("https://br.investing.com/currencies/"+strings.ToLower(fromCurrency)+"-"+strings.ToLower(toCurrency))

	body = "{"+body+"}"

	return []byte(body)
}

func GetCurrentAssetData(assets []string) []byte {
	results := ""
	body := ""
	asset := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("br.investing.com"),
		colly.AllowURLRevisit(),
	)


	// Extracts asset value
	c.OnHTML(`#__next`, func(e *colly.HTMLElement) {
		body = ""

		goquerySelection := e.DOM

		price := ""
		price = goquerySelection.Find(`.flex div.leading-9`).Text()

		body = body + CreateJsonStringField("asset",asset, true)
		body = body + CreateJsonStringField("price",price, false)
	})

	for _, item := range assets {
		asset = item
		c.Visit("https://br.investing.com/equities/"+asset)
		body = "{"+body+"}"
		results = results + body + ","
	}

	if len(results) > 0 {
		// Removes last comma
		results = results[:len(results) - 1]
	}
	results = "["+results+"]"
	return []byte(results)
}