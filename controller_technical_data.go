package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetTechnicalStocksData(c *gin.Context) {
	assets, okAssets := c.GetQuery("assets")
	if !okAssets {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'assets'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", GetCurrentAssetData(strings.Split(assets,",")))	
}