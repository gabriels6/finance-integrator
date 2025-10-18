package controllers

import (
	"net/http"
	"time"

	"github.com/gabriels6/finance-integrator/apis"
	"github.com/gin-gonic/gin"
)

func GetBcbApiExchangeRate(c *gin.Context) {
	currency, ok := c.GetQuery("currency")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'currency'"}`))
		return
	}
	timeString, ok := c.GetQuery("date")
	date := time.Now()
	if ok {
		parsedTime, err := time.Parse("2006-01-02", timeString)
		if err == nil {
			date = parsedTime
		}
	}
	c.Data(http.StatusOK, "application/json", apis.GetBcbExchangeRate(currency, date.Format("01-02-2006")))
}

func GetBcbApiExchangeRateByPeriod(c *gin.Context) {
	currency, ok := c.GetQuery("currency")
	if !ok {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"Not found parameter: 'currency'"}`))
		return
	}
	startTimeString, ok := c.GetQuery("start_date")
	startDate := time.Now()
	if ok {
		parsedTime, err := time.Parse("2006-01-02", startTimeString)
		if err == nil {
			startDate = parsedTime
		}
	}
	endTimeString, ok := c.GetQuery("end_date")
	endDate := time.Now()
	if ok {
		parsedTime, err := time.Parse("2006-01-02", endTimeString)
		if err == nil {
			endDate = parsedTime
		}
	}
	c.Data(http.StatusOK, "application/json", apis.GetBcbExchangeRateByPeriod(currency, startDate.Format("01-02-2006"), endDate.Format("01-02-2006")))
}
