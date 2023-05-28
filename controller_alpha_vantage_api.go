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