package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetIndexesData(c *gin.Context) {
	symbol, okSymbol := c.GetQuery("symbol")
	if !okSymbol {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbol'"}`))
		return
	}
	if symbol == "CDI" {
		c.Data(http.StatusOK, "application/json", CDIData())	
	} else if symbol == "IBOV" {
		c.Data(http.StatusOK, "application/json", IBOVData())	
	} else {
		c.Data(http.StatusOK, "application/json", []byte(""))	
	}
	
}

func GetIndexByInvesting(c *gin.Context) {
	symbols, okSymbols := c.GetQuery("symbols")
	if !okSymbols {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbols'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", IndexDataByInvesting(symbols))	
}