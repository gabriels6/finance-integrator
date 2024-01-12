package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/gabriels6/finance-integrator/scrapers"
)

func GetFundamentalistStockData(c *gin.Context) {
	asset, okAsset := c.GetQuery("asset")
	if !okAsset {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'asset'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.GetStockData(asset))	
}

func GetFundamentalistImobiliaryFundData(c *gin.Context) {
	asset, okAsset := c.GetQuery("asset")
	if !okAsset {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'asset'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.GetImobiliaryFundData(asset))	
}

func GetFundamentalistAllImobiliaryFundData(c *gin.Context) {
	amountOfElements, okAmountOfElements := c.GetQuery("amountOfElements")
	offset, okOffset := c.GetQuery("offset")

	if !okOffset {
		offset = "0"
	}
	if !okAmountOfElements {
		amountOfElements = "10"
	}

	resultAmountOfElements, errorAmountOfElements := strconv.Atoi(amountOfElements)
	resultOffset, errorOffset := strconv.Atoi(offset)
	


	if errorAmountOfElements != nil{
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Error converting param: 'amountOfElements'"}`))
		return
	}
	if errorOffset != nil{
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Error converting param: 'offset'"}`))
		return
	}

	c.Data(http.StatusOK, "application/json", scrapers.GetAllImoboliaryFundsData(resultOffset, resultAmountOfElements))
}

func GetFundamentalistAllStocksData(c *gin.Context) {
	pages, okPages := c.GetQuery("pages")
	offset, okOffset := c.GetQuery("offset")
	if !okPages {
		pages = "1"
	}
	if !okOffset {
		offset = "1"
	}
	pagesAmount, errorPagesAmount := strconv.Atoi(pages)
	offsetAmount, errorOffsetAmount := strconv.Atoi(offset)
	if errorPagesAmount != nil {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Error converting param: 'pages'"}`))
		return
	}
	if errorOffsetAmount != nil {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Error converting param: 'offset'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.GetAllFundamentslistStocksData(pagesAmount, offsetAmount))
}

func GetDividendsData(c *gin.Context) {
	asset, okAsset := c.GetQuery("asset")
	if !okAsset {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'asset'"}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.GetDividends(asset))	
}

func GetHistoricalExchangeRatesData(c *gin.Context) {
	fromCurrency, okFromCurrency := c.GetQuery("fromCurrency")
	toCurrency, okToCurrency := c.GetQuery("toCurrency")
	if !okFromCurrency || !okToCurrency {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Params 'fromCurrency' or 'toCurrency' not found."}`))
		return
	}
	c.Data(http.StatusOK, "application/json", scrapers.GetHistoricalExchangeRates(fromCurrency, toCurrency))
}
