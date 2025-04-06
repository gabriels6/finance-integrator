package twelvedataapi

import (
	"fmt"

	"github.com/gabriels6/finance-integrator/apis"
	"github.com/gabriels6/finance-integrator/utils"
)

const TWELVE_DATA_BASE_URL = "https://api.twelvedata.com"

func GetSeries(symbols string) []byte {
	return apis.CallApi(fmt.Sprintf("%s/time_series?symbol=%s&apikey=%s&interval=1day", TWELVE_DATA_BASE_URL, symbols, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetEodPrices(symbols string) []byte {
	// return []byte(fmt.Sprintf("{\"req_1\":{ \"/eod?symbol=%s&apikey=%s\"} }", symbols, utils.GetEnv("TWELVE_DATA_API_KEY")))
	return apis.PostApi(fmt.Sprintf("%s/batch", TWELVE_DATA_BASE_URL), fmt.Sprintf("{\"req_1\": { \"url\": \"/eod?symbol=%s&apikey=%s\"} }", symbols, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func SearchSymbol(keyword string) []byte {
	return apis.CallApi(fmt.Sprintf("%s/symbol_search?symbol=%s&apikey=%s", TWELVE_DATA_BASE_URL, keyword, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetStocks() []byte {
	return apis.CallApi(fmt.Sprintf("%s/stocks?apikey=%s", TWELVE_DATA_BASE_URL, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetStock(symbol string) []byte {
	return apis.CallApi(fmt.Sprintf("%s/stocks?symbol=%s&apikey=%s", TWELVE_DATA_BASE_URL, symbol, utils.GetEnv("TWELVE_DATA_API_KEY")))
}

func GetETFs() []byte {
	return apis.CallApi(fmt.Sprintf("%s/etfs?apikey=%s", TWELVE_DATA_BASE_URL, utils.GetEnv("TWELVE_DATA_API_KEY")))
}
