package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gabriels6/finance-integrator/scrapers"
)

func GetYearlyQuotes(c *gin.Context) {
	asset, okAsset := c.GetQuery("asset")
	if !okAsset {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'asset'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.YearlyQuotes(asset))
}

func GetHistoricalQuotes(c *gin.Context) {
	asset, okAsset := c.GetQuery("asset")
	if !okAsset {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'asset'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.HistoricalQuotes(asset))
}