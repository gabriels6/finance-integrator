package main

import (
	"net/http"

	"github.com/gabriels6/finance-integrator/controllers"
	controllertwelvedataapi "github.com/gabriels6/finance-integrator/controllers/controller_twelve_data_api"
	"github.com/gabriels6/finance-integrator/utils"
	"github.com/gin-gonic/gin"
)

func VerifyHeaderMiddleware(c *gin.Context) {
	tokens := c.Request.Header.Get("x_api_key")
	if tokens != utils.GetEnv("API_TOKEN") {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Invalid Token"})
		return
	}
	c.Next()
}

func Routes() *gin.Engine {
	router := gin.Default()
	router.Use(VerifyHeaderMiddleware)
	router.GET("/alpha-vantage/search", controllers.GetSymbol)
	router.GET("/alpha-vantage/global-quote", controllers.GetQuote)
	router.GET("/alpha-vantage/time-series-weekly", controllers.GetTimeSeriesWeekly)
	router.GET("/alpha-vantage/overview", controllers.GetOverview)
	router.GET("/alpha-vantage/news", controllers.GetNews)
	router.GET("/alpha-vantage/exchange-rate", controllers.GetExRate)
	router.GET("/fundamentalist-data/stock", controllers.GetFundamentalistStockData)
	router.GET("/fundamentalist-data/imobiliary-fund", controllers.GetFundamentalistImobiliaryFundData)
	router.GET("/fundamentalist-data/imobiliary-funds", controllers.GetFundamentalistAllImobiliaryFundData)
	router.GET("/fundamentalist-data/stocks", controllers.GetFundamentalistAllStocksData)
	router.GET("/fundamentalist-data/dividends", controllers.GetDividendsData)
	router.GET("/fundamentalist-data/rates", controllers.GetHistoricalExchangeRatesData)
	router.GET("/technical-data/stock", controllers.GetTechnicalStocksData)
	router.GET("/technical-data/exchange-rate", controllers.GetInvestingExchangeRateRoute)
	router.GET("/indexes", controllers.GetIndexesData)
	router.GET("/investing/indexes", controllers.GetIndexByInvesting)
	router.GET("/fixed-income/government", controllers.GetBrazilianGovernmentBondsRoute)
	router.GET("/fixed-income/debentures", controllers.GetDebenturesRoute)
	router.GET("/yahoo/yearly-quotes", controllers.GetYearlyQuotes)
	router.GET("/yahoo/historical-quotes", controllers.GetHistoricalQuotes)
	router.GET("/yahoo/financial-data", controllers.GetFinancialData)
	router.GET("/bcb/exchange-rate", controllers.GetBcbApiExchangeRate)
	router.GET("/bcb/exchange-rate-period", controllers.GetBcbApiExchangeRateByPeriod)
	router.GET("/twelve-data/eod-prices", controllertwelvedataapi.GetEodPrices)
	router.GET("/twelve-data/stocks", controllertwelvedataapi.GetStocks)
	router.GET("/twelve-data/etfs", controllertwelvedataapi.GetETFs)
	router.GET("/twelve-data/time-series", controllertwelvedataapi.GetTimeSeriesDaily)
	router.GET("/twelve-data/gather-prices", controllertwelvedataapi.GatherWebsocketRealTimeQuotes)
	router.GET("/twelve-data/realtime-prices", controllertwelvedataapi.GetWsRealtimePrices)
	router.GET("/nasdaq/stock-screener", controllers.GetNasdaqStockScreener)
	router.GET("/nasdaq/etf-screener", controllers.GetNasdaqEtfScreener)
	return router
}
