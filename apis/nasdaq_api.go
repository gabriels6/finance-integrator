package apis

func StocksScreener() []byte {
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36Preconnected origins in 'Network dependency tree' insightThe 'Network dependency tree' insight now shows you a list of used or unused preconnected origins and preconnect candidates, if any.Server response and redirection times in 'Document request latency' insightThe 'Document request latency' insight now shows you server response time and, if any, redirection time.Geolocation accuracy parameter in SensorsThe Sensors panel now lets you set accuracy in geolocation emulation, so you can test the handling of different levels of GPS accuracy."

	return CallApiWithHeaders("https://api.nasdaq.com/api/screener/etf", headers)
}
