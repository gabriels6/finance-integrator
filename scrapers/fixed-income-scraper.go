package scrapers

import (
	"github.com/gocolly/colly/v2"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

func GetBrazilianGovernmentBonds() []byte {
	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("statusinvest.com.br"),
	)

	// Extracts asset value
	c.OnHTML(`#tesouro-section`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		
		goquerySelection.Find(`div[title="Ir para o detalhe do título"]`).Each (func(index int,item *goquery.Selection) {
			fixedIncomeItem := ""
			name := item.Find(`.tesouro h4`).Text()
			indexer := strings.Replace(item.Find(`.tesouro span[title="Tipo de rendimento"]`).Text(), "\n", "", 3)
			tax := item.Find(`.tesouro b[title="Juros aplicado"]`).Text()
			buyPrice := strings.Replace(item.Find(`.tesouro span[title="Valor de compra do título"]`).Text(), "\n", "", 3)
			sellPrice := strings.Replace(item.Find(`.tesouro span[title="Valor de venda do título"]`).Text(), "\n", "", 3)
			fixedIncomeItem = fixedIncomeItem + CreateJsonStringField("asset",name, true)
			fixedIncomeItem = fixedIncomeItem + CreateJsonStringField("indexer",indexer, true)
			fixedIncomeItem = fixedIncomeItem + CreateJsonStringField("tax",tax, true)
			fixedIncomeItem = fixedIncomeItem + CreateJsonStringField("buyPrice",buyPrice, true)
			fixedIncomeItem = fixedIncomeItem + CreateJsonStringField("sellPrice",sellPrice, false)
			fixedIncomeItem = "{"+fixedIncomeItem+"}"
			body = body + fixedIncomeItem + ","
		})
	})

	c.Visit("https://statusinvest.com.br/#tesouro-section")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
}

func GetDebentures(date time.Time) []byte {
	body := ""

	currentTime := date

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.anbima.com.br"),
	)

	// Extracts asset value
	c.OnResponse(func(r *colly.Response) {

		for index, element := range strings.Split(string(r.Body[:]),"\n") {
			if index > 2 {
				item := strings.Split(element, "@")
				if len(item) > 1 {
					asset := item[0]
					issuer := item[1]
					expireDate := item[2]
					tax := item[3]
					indexer := strings.Split(tax, " ")[0]
					buyTax := item[4]
					sellTax :=item[5]
					price := item[10]
					duration := item[12]

					json := ""

					json = json + CreateJsonStringField("asset",asset, true)
					json = json + CreateJsonStringField("issuer",issuer, true)
					json = json + CreateJsonStringField("expireDate",expireDate, true)
					json = json + CreateJsonStringField("indexer",indexer, true)
					json = json + CreateJsonStringField("tax",tax, true)
					json = json + CreateJsonStringField("buyTax",buyTax, true)
					json = json + CreateJsonStringField("sellTax",sellTax, true)
					json = json + CreateJsonStringField("price",price, true)
					json = json + CreateJsonStringField("duration",duration, false)
					body = body + "{" + json + "},"
				}
			}
		}

	})

	c.Visit("https://www.anbima.com.br/informacoes/merc-sec-debentures/arqs/db"+currentTime.Format("060102")+".txt")

	if len(body) > 0 {
		// Removes last comma
		body = body[:len(body) - 1]
	}

	body = "["+body+"]"

	return []byte(body)
}