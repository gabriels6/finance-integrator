package main

func SearchSymbol(keyword string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=" + keyword + "&apikey=" + GetEnv("ALPHA_VANTAGE_API_KEY"))
}

func GlobalQuotes(symbol string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + GetEnv("ALPHA_VANTAGE_API_KEY"))
}

func TimeSeriesWeekly(symbol string) []byte {
	return CallApi("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=" + symbol + "&apikey=" + GetEnv("ALPHA_VANTAGE_API_KEY"))
}
