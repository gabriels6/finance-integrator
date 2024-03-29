package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/gabriels6/finance-integrator/scrapers"
)

func GetBrazilianGovernmentBondsRoute(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", scrapers.GetBrazilianGovernmentBonds())	
}

func GetDebenturesRoute(c *gin.Context) {
	timeString, ok := c.GetQuery("date")
	date := time.Now()
	if ok {
		parsedTime, err := time.Parse("2006-01-02", timeString)
		if err == nil { date = parsedTime }
	}
	c.Data(http.StatusOK, "application/json", scrapers.GetDebentures(date))	
}