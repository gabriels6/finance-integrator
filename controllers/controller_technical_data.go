package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/gabriels6/finance-integrator/scrapers"
)

func GetTechnicalStocksData(c *gin.Context) {
	assets, okAssets := c.GetQuery("assets")
	if !okAssets {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'assets'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.GetCurrentAssetData(strings.Split(assets,",")))	
}

func GetInvestingExchangeRateRoute(c *gin.Context) {
	fromCurrency, okFromCurrency := c.GetQuery("fromCurrency")
	toCurrency, okToCurrency := c.GetQuery("toCurrency")
	if !okFromCurrency {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'fromCurrency'"}`))
		return
	}
	if !okToCurrency {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'toCurrency'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.GetInvestingExchangeRate(fromCurrency, toCurrency))	
}