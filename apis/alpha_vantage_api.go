package apis

import (
	"github.com/gabriels6/finance-integrator/utils"
)

func SearchSymbol(keyword string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=" + keyword + "&apikey=" + utils.GetEnv("ALPHA_VANTAGE_API_KEY"))
}

func GlobalQuotes(symbol string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + utils.GetEnv("ALPHA_VANTAGE_API_KEY"))
}

func TimeSeriesWeekly(symbol string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=" + symbol + "&apikey=" + utils.GetEnv("ALPHA_VANTAGE_API_KEY"))
}

func Overview(symbol string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=OVERVIEW&symbol=" + symbol + "&apikey=" + utils.GetEnv("ALPHA_VANTAGE_API_KEY"))
}

func News(symbols string, topics string, sort string, limit string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers=" + symbols + "&topics="+ topics +"&sort="+ sort +"&limit="+ limit +"&apikey=" + utils.GetEnv("ALPHA_VANTAGE_API_KEY"))
}

func ExchangeRate(from_currency string, to_currency string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency="+from_currency+"&to_currency="+to_currency+"&apikey=" + utils.GetEnv("ALPHA_VANTAGE_API_KEY"))
}