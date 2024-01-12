package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gabriels6/finance-integrator/scrapers"
)

func GetIndexesData(c *gin.Context) {
	symbol, okSymbol := c.GetQuery("symbol")
	if !okSymbol {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'symbol'"}`))
		return
	}
	if symbol == "CDI" {
		c.Data(http.StatusOK, "application/json", scrapers.CDIData())	
	} else if symbol == "IPCA" {
		c.Data(http.StatusOK, "application/json", scrapers.IPCAData())	
	} else if symbol == "IBOV" {
		c.Data(http.StatusOK, "application/json", scrapers.IBOVData())	
	} else if symbol == "IBXX" {
		c.Data(http.StatusOK, "application/json", scrapers.IBXXData())	
	} else if symbol == "IDIV" {
		c.Data(http.StatusOK, "application/json", scrapers.IDIVData())	
	} else if symbol == "IFIX" {
		c.Data(http.StatusOK, "application/json", scrapers.IFIXData())	
	} else if symbol == "SP500" {
		c.Data(http.StatusOK, "application/json", scrapers.SP500Data())	
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
	c.Data(http.StatusOK, "application/json", scrapers.IndexDataByInvesting(symbols))	
}