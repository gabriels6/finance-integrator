package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetSymbol(c *gin.Context) {
	keyword, ok := c.GetQuery("keyword")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'keyword'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", SearchSymbol(keyword))	
}

func GetQuote(c *gin.Context) {
	symbol, ok := c.GetQuery("symbol")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbol'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", GlobalQuotes(symbol))
}

func GetTimeSeriesWeekly(c *gin.Context) {
	symbol, ok := c.GetQuery("symbol")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbol'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", TimeSeriesWeekly(symbol))
}

func GetOverview(c *gin.Context) {
	symbol, ok := c.GetQuery("symbol")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbol'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", Overview(symbol))
}

func GetNews(c *gin.Context) {
	symbols, symError := c.GetQuery("symbols")
	topics, _ := c.GetQuery("topics")
	sort, _ := c.GetQuery("sort")
	limit, _ := c.GetQuery("limit")
	if !symError {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbols'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", News(symbols, topics, sort, limit))
}

func GetExRate(c *gin.Context) {
	fromCurrency, fromCurrError := c.GetQuery("fromCurrency")
	toCurrency, toCurrError := c.GetQuery("toCurrency")
	if !fromCurrError {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'fromCurrency'"}`))
		return
	}
	if !toCurrError {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'toCurrency'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", ExchangeRate(fromCurrency, toCurrency))
}