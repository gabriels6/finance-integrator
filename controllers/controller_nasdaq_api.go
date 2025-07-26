package controllers

import (
	"net/http"

	"github.com/gabriels6/finance-integrator/apis"
	"github.com/gin-gonic/gin"
)

func GetNasdaqStockScreener(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", apis.StocksScreener())
}

func GetNasdaqEtfScreener(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", apis.EtfScreener())
}
