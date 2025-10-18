package controllertwelvedataapi

import (
	"fmt"
	"net/http"

	"github.com/gabriels6/finance-integrator/apis"
	twelvedataapi "github.com/gabriels6/finance-integrator/apis/twelve_data_api"
	"github.com/gin-gonic/gin"
)

func GetSymbol(c *gin.Context) {
	keyword, ok := c.GetQuery("keyword")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'keyword'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", apis.SearchSymbol(keyword))
}

func GetTimeSeriesDaily(c *gin.Context) {
	symbol, ok := c.GetQuery("symbol")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbol'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", twelvedataapi.GetSeries(symbol))
}

func GetEodPrices(c *gin.Context) {
	symbols, ok := c.GetQuery("symbols")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbols'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", twelvedataapi.GetEodPrices(symbols))
}

func GetStocks(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", twelvedataapi.GetStocks())
}

func GetETFs(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", twelvedataapi.GetETFs())
}

func GatherWebsocketRealTimeQuotes(c *gin.Context) {
	symbols, ok := c.GetQuery("symbols")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbols'"}`))
		return
	}
	err := twelvedataapi.GatherWebsocketRealTimeQuotes(symbols)
	message := ""
	if err != nil {
		message = fmt.Sprintf(`{"message": "%v"}`, err)
	}
	c.Data(http.StatusOK, "application/json", []byte(message))
}

func GetWsRealtimePrices(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", twelvedataapi.GetWsRealtimePrices())
}
