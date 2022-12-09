package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyHeaderMiddleware(c *gin.Context) {
	tokens := c.Request.Header.Get("x_api_key")
	if tokens != GetEnv("API_TOKEN") {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message":"Invalid Token"})
		return
	}
	c.Next()
}

func Routes() *gin.Engine {
	router := gin.Default()
	router.Use(VerifyHeaderMiddleware)
	router.GET("/alpha-vantage/search", GetSymbol)
	router.GET("/alpha-vantage/global-quote", GetQuote)
	router.GET("/alpha-vantage/time-series-weekly", GetTimeSeriesWeekly)
	router.GET("/fundamentalist-data/stock", GetFundamentalistStockData)
	router.GET("/fundamentalist-data/imobiliary-fund", GetFundamentalistImobiliaryFundData)
	router.GET("/fundamentalist-data/imobiliary-funds",GetFundamentalistAllImobiliaryFundData)
	router.GET("/fundamentalist-data/stocks",GetFundamentalistAllStocksData)
	router.GET("/fundamentalist-data/dividends",GetDividendsData)
	router.GET("/technical-data/stock",GetTechnicalStocksData)
	return router
}