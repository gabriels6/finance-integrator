package scrapers

import (
	"github.com/gocolly/colly/v2"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func YearlyQuotes(symbol string) []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("finance.yahoo.com"),
	)

	// Extracts asset value
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		goquerySelection.Find(`table[data-test="historical-prices"] tbody tr`).Each (func(index int,item *goquery.Selection) {
			indexItem := ""
			
			date := item.Find("td:nth-child(1)").Text()
			value := item.Find("td:nth-child(5)").Text()

			indexItem = indexItem + CreateJsonStringField("date",date, true)
			indexItem = indexItem + CreateJsonStringField("value",value, false)

			indexItem = "{"+indexItem+"},"

			body = body + indexItem
		})
	})

	c.Visit("https://finance.yahoo.com/quote/"+symbol+"/history?frequency=1wk")

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