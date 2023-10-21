package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func CDIData() []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("investidor10.com.br"),
	)

	// Extracts asset value
	c.OnHTML(`main`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		goquerySelection.Find(`table.table-bordered tbody tr`).Each (func(index int,item *goquery.Selection) {
			cdiItem := ""
			year := item.Find("td:nth-child(1)").Text()
			jan := item.Find("td:nth-child(2)").Text()
			fev := item.Find("td:nth-child(3)").Text()
			mar := item.Find("td:nth-child(4)").Text()
			apr := item.Find("td:nth-child(5)").Text()
			mai := item.Find("td:nth-child(6)").Text()
			jun := item.Find("td:nth-child(7)").Text()
			jul := item.Find("td:nth-child(8)").Text()
			aug := item.Find("td:nth-child(9)").Text()
			sep := item.Find("td:nth-child(10)").Text()
			oct := item.Find("td:nth-child(11)").Text()
			nov := item.Find("td:nth-child(12)").Text()
			dec := item.Find("td:nth-child(13)").Text()

			cdiItem = cdiItem + CreateJsonStringField("year",year, true)
			cdiItem = cdiItem + CreateJsonStringField("jan",jan, true)
			cdiItem = cdiItem + CreateJsonStringField("fev",fev, true)
			cdiItem = cdiItem + CreateJsonStringField("mar",mar, true)
			cdiItem = cdiItem + CreateJsonStringField("apr",apr, true)
			cdiItem = cdiItem + CreateJsonStringField("mai",mai, true)
			cdiItem = cdiItem + CreateJsonStringField("jun",jun, true)
			cdiItem = cdiItem + CreateJsonStringField("jul",jul, true)
			cdiItem = cdiItem + CreateJsonStringField("aug",aug, true)
			cdiItem = cdiItem + CreateJsonStringField("sep",sep, true)
			cdiItem = cdiItem + CreateJsonStringField("oct",oct, true)
			cdiItem = cdiItem + CreateJsonStringField("nov",nov, true)
			cdiItem = cdiItem + CreateJsonStringField("dec",dec, false)

			cdiItem = "{"+cdiItem+"},"

			body = body + cdiItem
		})
	})

	c.Visit("https://investidor10.com.br/indices/cdi/")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
}

func IBOVData() []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("br.financas.yahoo.com"),
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

	c.Visit("https://br.financas.yahoo.com/quote/%5EBVSP/history?period1=946857600&period2=1685232000&interval=1mo&filter=history&frequency=1mo&includeAdjustedClose=true")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
}

func IDIVData() []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("br.advfn.com"),
	)

	// Extracts asset value
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		goquerySelection.Find(`.TableElement:nth-child(8) table tbody tr`).Each (func(index int,item *goquery.Selection) {
			indexItem := ""
			
			date := strings.Trim(item.Find("td:nth-child(1)").Text()," ")
			value := strings.Trim(item.Find("td:nth-child(2)").Text()," ")

			if date != "" {
				indexItem = indexItem + CreateJsonStringField("date",date, true)
				indexItem = indexItem + CreateJsonStringField("value",value, false)

				indexItem = "{"+indexItem+"},"

				body = body + indexItem
			}
		})
	})

	c.Visit("https://br.advfn.com/bolsa-de-valores/bovespa/dividend-IDIV/historico")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
}

func IBXXData() []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("br.advfn.com"),
	)

	// Extracts asset value
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		goquerySelection.Find(`.TableElement:nth-child(8) table tbody tr`).Each (func(index int,item *goquery.Selection) {
			indexItem := ""
			
			date := strings.Trim(item.Find("td:nth-child(1)").Text()," ")
			value := strings.Trim(item.Find("td:nth-child(2)").Text()," ")

			if date != "" {
				indexItem = indexItem + CreateJsonStringField("date",date, true)
				indexItem = indexItem + CreateJsonStringField("value",value, false)

				indexItem = "{"+indexItem+"},"

				body = body + indexItem
			}
		})
	})

	c.Visit("https://br.advfn.com/bolsa-de-valores/bovespa/indice-brasil-100-IBXX/historico")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
}

func SP500Data() []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("br.advfn.com"),
	)

	// Extracts asset value
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		goquerySelection.Find(`.TableElement:nth-child(8) table tbody tr`).Each (func(index int,item *goquery.Selection) {
			indexItem := ""
			
			date := strings.Trim(item.Find("td:nth-child(1)").Text()," ")
			value := strings.Trim(item.Find("td:nth-child(2)").Text()," ")

			if date != "" {
				indexItem = indexItem + CreateJsonStringField("date",date, true)
				indexItem = indexItem + CreateJsonStringField("value",value, false)

				indexItem = "{"+indexItem+"},"

				body = body + indexItem
			}
		})
	})

	c.Visit("https://br.advfn.com/bolsa-de-valores/spi/SP500/historico")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
}

func IndexDataByInvesting(symbols string) []byte {

	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("br.investing.com"),
	)

	indexEndpoint := ""
	queryString := ""

	symbolName := ""

	// Extracts asset value
	c.OnHTML(`body`, func(e *colly.HTMLElement) {

		goquerySelection := e.DOM

		indexItem := ""
			
		value := goquerySelection.Find(queryString).Text()

		indexItem = indexItem + CreateJsonStringField("symbol",symbolName, true)
		indexItem = indexItem + CreateJsonStringField("value",value, false)

		body = body + "{"+indexItem+"},"
	})

	for _, symbol := range strings.Split(symbols, ",") {
		symbol = strings.ToLower(symbol)
		symbolName = symbol
	
		if symbol == "ipca" {
			indexEndpoint = "economic-calendar/brazilian-cpi-410"
			queryString = "#releaseInfo span:nth-child(2) div"
		}
	
		if symbol == "ibov" {
			indexEndpoint = "indices/bovespa"
			queryString = ".text-5xl.font-bold.leading-9"
		}
	
		if symbol == "ibrx" {
			indexEndpoint = "indices/brazil-index"
			queryString = ".text-5xl.font-bold.leading-9"
		}
	
		if symbol == "selic" {
			indexEndpoint = "economic-calendar/brazilian-interest-rate-decision-415"
			queryString = "#releaseInfo span:nth-child(2) div"
		}
	
		if symbol == "sp500" {
			indexEndpoint = "indices/us-spx-500"
			queryString = ".text-5xl.font-bold.leading-9"
		}
	
		if symbol == "idiv" {
			indexEndpoint = "indices/bovespa-dividend"
			queryString = ".text-5xl.font-bold.leading-9"
		}
	
		if symbol == "ifix" {
			indexEndpoint = "indices/bm-fbovespa-real-estate-ifix"
			queryString = ".text-5xl.font-bold.leading-9"
		}
		

		c.Visit("https://br.investing.com/"+indexEndpoint)
	}

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"
	

	return []byte(body)
}