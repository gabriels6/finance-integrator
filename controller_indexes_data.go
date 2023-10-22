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
	} else if symbol == "IPCA" {
		c.Data(http.StatusOK, "application/json", IPCAData())	
	} else if symbol == "IBOV" {
		c.Data(http.StatusOK, "application/json", IBOVData())	
	} else if symbol == "IBXX" {
		c.Data(http.StatusOK, "application/json", IBXXData())	
	} else if symbol == "IDIV" {
		c.Data(http.StatusOK, "application/json", IDIVData())	
	} else if symbol == "IFIX" {
		c.Data(http.StatusOK, "application/json", IFIXData())	
	} else if symbol == "SP500" {
		c.Data(http.StatusOK, "application/json", SP500Data())	
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