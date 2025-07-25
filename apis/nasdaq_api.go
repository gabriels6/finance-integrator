package apis

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func StocksScreener() []byte {

	body := ""

	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("api.nasdaq.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "application/json, text/plain, */*")
		r.Headers.Set("Origin", "https://www.nasdaq.com")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")
		r.Headers.Set("Referer", "https://www.nasdaq.com/")
	})

	c.OnResponse(func(r *colly.Response) {
		body = string(r.Body)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})

	c.Visit("https://api.nasdaq.com/api/screener/stocks?download=true")

	return []byte(body)
}
