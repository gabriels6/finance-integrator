package scrapers

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"time"
)

func GetStockData(assetName string) []byte {

	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("statusinvest.com.br"),
	)

	// Extracts asset value
	c.OnHTML(`main`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		assetValue := goquerySelection.Find(`div[title="Valor atual do ativo"] strong.value`).Text()
		minValue := goquerySelection.Find(`div[title="Valor mínimo das últimas 52 semanas"] strong.value`).Text()
		maxValue := goquerySelection.Find(`div[title="Valor máximo das últimas 52 semanas"] strong.value`).Text()
		dividendYield := goquerySelection.Find(`div[title="Indicador utilizado para relacionar os proventos pagos por uma companhia e o preço atual de suas ações."] strong.value`).Text()
		dividendMoney12Months := goquerySelection.Find(`div[title="Soma total de proventos distribuídos nos últimos 12 meses"] span.sub-value`).Text()
		precoLucro := goquerySelection.Find(`div[title="Dá uma ideia do quanto o mercado está disposto a pagar pelos lucros da empresa."] strong.value`).Text()
		pVP := goquerySelection.Find(`div[title="Facilita a análise e comparação da relação do preço de negociação de um ativo com seu VPA."] strong.value`).Text()
		pEbitda := goquerySelection.Find(`div[title="O EBITDA permite conhecer quanto a companhia está gerando de caixa com base exclusivamente em suas atividades operacionais, desconsiderando os impactos financeiros e dos impostos."] strong.value`).Text()
		valorPatrimonialAcao := goquerySelection.Find(`div[title="Indica qual o valor patrimonial de uma ação."] strong.value`).Text()
		lucroPorAcao := goquerySelection.Find(`div[title="Indicar se a empresa é ou não lucrativa. Se este número estiver negativo, a empresa está com margens baixas, acumulando prejuízos."] strong.value`).Text()
		variacaoDiaria := strings.Replace(goquerySelection.Find(`span[title="Variação do valor do ativo com base no dia anterior"] b`).Text(), "%", "" , 400)
		roe := goquerySelection.Find(`div[title="Mede a capacidade de agregar valor de uma empresa a partir de seus próprios recursos e do dinheiro de investidores."] strong.value`).Text()
		roa := goquerySelection.Find(`div[title="O retorno sobre os ativos ou Return on Assets, é um indicador de rentabilidade, que calcula a capacidade de uma empresa gerar lucro a partir dos seus ativos, além de indiretamente, indicar a eficiência dos seus gestores."] strong.value`).Text()
		margemLucro := goquerySelection.Find(`div[title="Revela a porcentagem de lucro em relação às receitas de uma empresa."] strong.value`).Text()
		cagr5YearEarnings := goquerySelection.Find(`div[data-group="4"] div.item:nth-child(1) strong.value`).Text()
		cagr5YearProfit := goquerySelection.Find(`div[data-group="4"] div.item:nth-child(2) strong.value`).Text()

		if(assetValue==""){
			return
		}

		body = body + CreateJsonStringField("asset",assetName, true)
		body = body + CreateJsonStringField("date",getDate(), true)
		body = body + CreateJsonStringField("assetValue",assetValue, true)
		body = body + CreateJsonStringField("minValue",minValue, true)
		body = body + CreateJsonStringField("maxValue",maxValue, true)
		body = body + CreateJsonStringField("dividendYield",dividendYield, true)
		body = body + CreateJsonStringField("dividendMoney12Months",dividendMoney12Months, true)
		body = body + CreateJsonStringField("precoLucro",precoLucro, true)
		body = body + CreateJsonStringField("pVP",pVP, true)
		body = body + CreateJsonStringField("pEbitda",pEbitda, true)
		body = body + CreateJsonStringField("vpa",valorPatrimonialAcao, true)
		body = body + CreateJsonStringField("variacaoDiaria",variacaoDiaria, true)
		body = body + CreateJsonStringField("roe",roe, true)
		body = body + CreateJsonStringField("roa",roa, true)
		body = body + CreateJsonStringField("margemLucro",margemLucro, true)
		body = body + CreateJsonStringField("cagr5YearEarnings",cagr5YearEarnings, true)
		body = body + CreateJsonStringField("cagr5YearProfit",cagr5YearProfit, true)
		body = body + CreateJsonStringField("lpa",lucroPorAcao, false)
		
	})

	urlArray := []string { 
		"https://statusinvest.com.br/acoes/",
		"https://statusinvest.com.br/acoes/eua/",
		"https://statusinvest.com.br/reits/",
		"https://statusinvest.com.br/fundos-imobiliarios/",
		"https://statusinvest.com.br/fiagros/",
		"https://statusinvest.com.br/etf/eua/",
	}

	for _, url := range urlArray {
		if body != "" {
			break;
		}
		c.Visit(url+assetName)
	}

	body = "{"+body+"}"

	return []byte(body)
}

func GetImobiliaryFundData(assetName string) []byte {

	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("statusinvest.com.br"),
	)

	// Extracts asset value
	c.OnHTML(`main#main-2`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		assetValue := goquerySelection.Find(`div[title="Valor atual do ativo"] strong.value`).Text()
		minValue := goquerySelection.Find(`div[title="Valor mínimo das últimas 52 semanas"] strong.value`).Text()
		maxValue := goquerySelection.Find(`div[title="Valor máximo das últimas 52 semanas"] strong.value`).Text()
		dividendYield := goquerySelection.Find(`div[title="Dividend Yield com base nos últimos 12 meses"] strong.value`).Text()
		dividendMoney12Months := goquerySelection.Find(`div[title="Soma total de proventos distribuídos nos últimos 12 meses"] span.sub-value`).Text()
		pVP := goquerySelection.Find(`div.info:has(h3:contains("P/VP")) strong.value`).Text()
		patrimonioPorCota := goquerySelection.Find(`div.info:has(span:contains("Valor patrim. p/cota")) strong.value`).Text()

		if assetValue == "" {
			c.Visit("https://statusinvest.com.br/fiagros/"+assetName)
			return
		}

		body = body + CreateJsonStringField("asset",assetName, true)
		body = body + CreateJsonStringField("date",getDate(), true)
		body = body + CreateJsonStringField("assetValue",assetValue, true)
		body = body + CreateJsonStringField("minValue",minValue, true)
		body = body + CreateJsonStringField("maxValue",maxValue, true)
		body = body + CreateJsonStringField("dividendYield",dividendYield, true)
		body = body + CreateJsonStringField("dividendMoney12Months",dividendMoney12Months, true)
		body = body + CreateJsonStringField("pVP",pVP, true)
		body = body + CreateJsonStringField("patrimonioPorCota",patrimonioPorCota, false)
		
	})

	c.Visit("https://statusinvest.com.br/fundos-imobiliarios/"+assetName)

	body = "{"+body+"}"

	return []byte(body)
}

func GetAllImoboliaryFundsData(offset int, amountOfElements int) []byte {
	fundsData := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("fiis.com.br"),
	)


	// Extracts asset value
	c.OnHTML(`div#funds-index`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		elementCounter := 0

		goquerySelection.Find(`#funds-list #items-wrapper .item`).Each (func(index int,item *goquery.Selection) {
			symbol := item.Find(".ticker a").Text()
			if symbol != "" && index >= offset && elementCounter < amountOfElements {
				fundsData = fundsData + string(GetImobiliaryFundData(strings.TrimSpace(symbol))) + ","
				elementCounter = elementCounter + 1
			}
		})
		
	})

	c.Visit("https://fiis.com.br/lista-de-fundos-imobiliarios/")

	if len(fundsData) > 0 {
		// Removes last comma
		fundsData = fundsData[:len(fundsData) - 1]
	}

	fundsData = "["+fundsData+"]"

	return []byte(fundsData)
}

func GetAllFundamentslistStocksData(pages int, offset int) []byte {
	stocksData := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.infomoney.com.br"),
	)

	pagesCounter := 0

	// Extracts asset value
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		goquerySelection.Find(`tbody tr`).Each (func(index int,item *goquery.Selection) {
			if(pagesCounter < pages+offset && pagesCounter > offset) {
				symbol := item.Find("td a").Text()
				if symbol != "" {
					stockItem := GetStockData(symbol)
					if string(stockItem) != "{}" {
						stocksData = stocksData + string(stockItem) + ","
					}
					
				}
			}
			pagesCounter += 1
		})
		
	})

	c.Visit("https://www.infomoney.com.br/cotacoes/empresas-b3/")
	
	if len(stocksData) > 0 {
		// Removes last comma
		stocksData = stocksData[:len(stocksData) - 1]
	}

	stocksData = "["+stocksData+"]"

	return []byte(stocksData)
}


// DIVIDENDOS
func GetDividends(assetName string) []byte {
	dividends := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("statusinvest.com.br"),
	)

	// Extracts asset value
	c.OnHTML(`main#main-2`, func(e *colly.HTMLElement) {

		dividends = e.ChildAttr("input[name='results']", "value")

		dividends = strings.Replace(dividends, "\"y\":0", "\"assetName\":\"" + assetName + "\"" , 400)
		dividends = strings.Replace(dividends, "\"m\":0,\"d\":0,", "" , 400)
		dividends = strings.Replace(dividends, "\"ed\"", "\"comDate\"", 400)
		dividends = strings.Replace(dividends, "\"pd\"", "\"paymentDate\"", 400)
		dividends = strings.Replace(dividends, "\"et\"", "\"type\"", 400)
		dividends = strings.Replace(dividends, "\"etd\"", "\"description\"", 400)
		dividends = strings.Replace(dividends, "\"v\"", "\"value\"", 400)
		
		
		
	})

	urlArray := []string { 
		"https://statusinvest.com.br/acoes/",
		"https://statusinvest.com.br/acoes/eua/",
		"https://statusinvest.com.br/reits/",
		"https://statusinvest.com.br/fundos-imobiliarios/",
		"https://statusinvest.com.br/fiagros/",
		"https://statusinvest.com.br/etf/eua/",
	}

	for _, url := range urlArray {
		if dividends != "" {
			break;
		}
		c.Visit(url+assetName)
	}


	return []byte(dividends)
}

//EXCHANGE RATES
func GetHistoricalExchangeRates(fromCurrency string, toCurrency string) []byte {

	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.exchangerates.org.uk"),
	)

	rates := ""

	// Extracts asset value
	c.OnHTML(`div#content-wrap`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		goquerySelection.Find("table#hist tr").Each (func(index int,item *goquery.Selection) {

			if(index > 1) {
				rateItem := ""

				dateString := strings.Split(item.Find(":nth-child(1)").Text(),"for")
				rateString := strings.Split(item.Find(":nth-child(2)").Text(),"=")

				if len(dateString) > 1 && len(rateString) > 1{
					rateItem += rateItem + CreateJsonStringField("date",strings.Trim(dateString[1]," "), true)
					rateItem += rateItem + CreateJsonStringField(fromCurrency,strings.Split(strings.Trim(rateString[0]," ")," ")[0], true)
					rateItem += rateItem + CreateJsonStringField(toCurrency,strings.Split(strings.Trim(rateString[1]," ")," ")[0], false)

					rates += "{"+rateItem+"},"
				}
			}
		})
	
	})

	c.Visit("https://www.exchangerates.org.uk/"+strings.ToUpper(fromCurrency)+"-"+strings.ToUpper(toCurrency)+"-exchange-rate-history.html")

	if len(rates) > 0 {
		// Removes last comma
		rates = rates[:len(rates) - 1]
	}

	rates = "["+rates+"]"

	return []byte(rates)
}

// Removes unwanted characters from fetched string
func ClearString(value string) string {
	return strings.Trim(strings.TrimSuffix(value,"-"),"\n")
}

func CreateJsonStringField(key string, value string, comma bool) string {
	value = ClearString(value)
	value = strings.Replace(value, ",",".",1)
	commaString := ""
	if comma {
		commaString = ","
	}
	return `"`+key+`":"`+value+`"`+commaString
}

func getDate() string {
	return time.Now().Format("2006-01-02")
}