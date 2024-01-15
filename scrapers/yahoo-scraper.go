package scrapers

import (
	"github.com/gocolly/colly/v2"
	"strings"
)

func YearlyQuotes(symbol string) []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("query1.finance.yahoo.com"),
	)

	// Extracts asset value
	c.OnResponse(func(r *colly.Response) {
		lines := strings.Split(string(r.Body),"\n")
		for idx, line := range lines {
			if idx > 0 {
				lineValues := strings.Split(line,",")
				item := ""
				item = item + CreateJsonStringField("date",string(lineValues[0]), true)
				item = item + CreateJsonStringField("open",string(lineValues[1]), true)
				item = item + CreateJsonStringField("high",string(lineValues[2]), true)
				item = item + CreateJsonStringField("low",string(lineValues[3]), true)
				item = item + CreateJsonStringField("close",string(lineValues[4]), true)
				item = item + CreateJsonStringField("adjusted_close",string(lineValues[5]), true)
				item = item + CreateJsonStringField("volume",string(lineValues[6]), false)
				body += "{"+item+"},"
			}
		}
	})

	c.Visit("https://query1.finance.yahoo.com/v7/finance/download/"+symbol+"?period1=1673788703&period2=1705324703&interval=1wk&events=history&includeAdjustedClose=true")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
} 

func HistoricalQuotes(symbol string) []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("query1.finance.yahoo.com"),
	)

	// Extracts asset value
	c.OnResponse(func(r *colly.Response) {
		lines := strings.Split(string(r.Body),"\n")
		for idx, line := range lines {
			if idx > 0 {
				lineValues := strings.Split(line,",")
				item := ""
				item = item + CreateJsonStringField("date",string(lineValues[0]), true)
				item = item + CreateJsonStringField("open",string(lineValues[1]), true)
				item = item + CreateJsonStringField("high",string(lineValues[2]), true)
				item = item + CreateJsonStringField("low",string(lineValues[3]), true)
				item = item + CreateJsonStringField("close",string(lineValues[4]), true)
				item = item + CreateJsonStringField("adjusted_close",string(lineValues[5]), true)
				item = item + CreateJsonStringField("volume",string(lineValues[6]), false)
				body += "{"+item+"},"
			}
		}
	})

	c.Visit("https://query1.finance.yahoo.com/v7/finance/download/"+symbol+"?period1=946857600&period2=1705104000&interval=1mo&filter=history&frequency=1mo&includeAdjustedClose=true")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
} 