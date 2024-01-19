package scrapers

import (
	"github.com/gocolly/colly/v2"
	"strings"
	"time"
)

func GetInvestingExchangeRate(fromCurrency string, toCurrency string) []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("br.investing.com"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
	})

	// Extracts asset value
	c.OnHTML(`#__next`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		price := ""
		price = goquerySelection.Find(`div[data-test=instrument-price-last]`).Text()
		

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
	for _, item := range assets {
		time.Sleep(time.Second * 1/100)
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

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
	})

	// Extracts asset value
	c.OnHTML(`#__next`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		price := goquerySelection.Find(`.flex div.leading-9`).Text()

		if price == "" {
			price = goquerySelection.Find(`.flex div.font-bold:first-child`).Text()
		}

		
		body = body + CreateJsonStringField("asset",asset, true)
		body = body + CreateJsonStringField("price",price, false)
	})

	c.Visit("https://br.investing.com/equities/"+asset)

	body = "{"+body+"}"

	return body
}