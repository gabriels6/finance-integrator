package main

import (
	"strings"
	"github.com/gocolly/colly/v2"
)

func GetStockData(assetName string) []byte {

	body := ""

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("statusinvest.com.br"),
	)

	// Extracts asset value
	c.OnHTML(`div.container div.paper`, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		assetValue := goquerySelection.Find(`div[title="Valor atual do ativo"] strong.value`).Text()
		minValue := goquerySelection.Find(`div[title="Valor mínimo das últimas 52 semanas"] strong.value`).Text()
		maxValue := goquerySelection.Find(`div[title="Valor máximo das últimas 52 semanas"] strong.value`).Text()
		dividendYield := goquerySelection.Find(`div[title="Indicador utilizado para relacionar os proventos pagos por uma companhia e o preço atual de suas ações."] strong.value`).Text()
		precoLucro := goquerySelection.Find(`div[title="Dá uma ideia do quanto o mercado está disposto a pagar pelos lucros da empresa."] strong.value`).Text()
		pVP := goquerySelection.Find(`div[title="Facilita a análise e comparação da relação do preço de negociação de um ativo com seu VPA."] strong.value`).Text()
		pEbitda := goquerySelection.Find(`div[title="O EBITDA permite conhecer quanto a companhia está gerando de caixa com base exclusivamente em suas atividades operacionais, desconsiderando os impactos financeiros e dos impostos."] strong.value`).Text()
		valorPatrimonialAcao := goquerySelection.Find(`div[title="Indica qual o valor patrimonial de uma ação."] strong.value`).Text()

		body = body + CreateJsonStringField("asset",assetName, true)
		body = body + CreateJsonStringField("assetValue",assetValue, true)
		body = body + CreateJsonStringField("minValue",minValue, true)
		body = body + CreateJsonStringField("maxValue",maxValue, true)
		body = body + CreateJsonStringField("dividendYield",dividendYield, true)
		body = body + CreateJsonStringField("precoLucro",precoLucro, true)
		body = body + CreateJsonStringField("pVP",pVP, true)
		body = body + CreateJsonStringField("pEbitda",pEbitda, true)
		body = body + CreateJsonStringField("vpa",valorPatrimonialAcao, false)
		
	})

	c.Visit("https://statusinvest.com.br/acoes/"+assetName)

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
		pVP := goquerySelection.Find(`div.info:has(h3:contains("P/VP")) strong.value`).Text()
		patrimonioPorCota := goquerySelection.Find(`div.info:has(span:contains("Valor patrim. p/cota")) strong.value`).Text()

		body = body + CreateJsonStringField("asset",assetName, true)
		body = body + CreateJsonStringField("assetValue",assetValue, true)
		body = body + CreateJsonStringField("minValue",minValue, true)
		body = body + CreateJsonStringField("maxValue",maxValue, true)
		body = body + CreateJsonStringField("dividendYield",dividendYield, true)
		body = body + CreateJsonStringField("pVP",pVP, true)
		body = body + CreateJsonStringField("patrimonioPorCota",patrimonioPorCota, false)
		
	})

	c.Visit("https://statusinvest.com.br/fundos-imobiliarios/"+assetName)

	body = "{"+body+"}"

	return []byte(body)
}

// Removes unwanted characters from fetched string
func ClearString(value string) string {
	return strings.TrimSuffix(value,"-")
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