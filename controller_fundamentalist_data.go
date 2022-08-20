package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetFundamentalistStockData(c *gin.Context) {
	asset, okAsset := c.GetQuery("asset")
	if !okAsset {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'asset'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", GetStockData(asset))	
}

func GetFundamentalistImobiliaryFundData(c *gin.Context) {
	asset, okAsset := c.GetQuery("asset")
	if !okAsset {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'asset'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", GetImobiliaryFundData(asset))	
}