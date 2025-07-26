package apis

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func StocksScreener() []byte {

	body := ""

	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("api.nasdaq.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br, zstd")
		r.Headers.Set("Accept-Language", "pt-BR,pt;q=0.5")
		r.Headers.Set("Cookie", "akaalb_ALB_Default=~op=ao_api__east1:ao_api_central1|~rv=87~m=ao_api_central1:0|~os=ff51b6e767de05e2054c5c99e232919a~id=9f35d6d40093920a3c5d926866688e1b; ak_bmsc=5E68A9724568A77716D612294F4CA30B~000000000000000000000000000000~YAAQjgMVAh0Z9hOYAQAAWAFTRBzOBtWZ8DaEKXGFykax4icTZzpapAERJF6eM9Qka/fHGyl1T7ZoIF/sOFBhyZ0RUr/0WE4cRIAmcTkCCCZ6U6u5E+xJIuF1X+zTKBpAwC6Bcjo6yPXEFtfg3GPRnk2caILC6PmLOpjDBF567R7FGXGFndIsyA1Lc8GQmoucir1pIwgW1zYTi4SH69J+7iRyOVwqjlJGNnF/1K4aYF+oLEQCk44oX4owzKPlD8sD8QbLNANIzBoH6nVatKDdF8wL66dEJeR6QlmdAt82q0lnTrKrONlFx8GLx2PLfbpRPkzdiFG0c9n1dsTxTOx8+g1yYT35CylsIbsG5UHHyM2RhRf1jesicx6I4vDjvpq4OGqZCX9/cbx+dY7MY4GF4zyp+NUi0NLNdL1BIOLDOZd2; bm_sv=B08812E763F847690E3569974DA1FFD8~YAAQjgMVAsUZ9hOYAQAAwUdTRBzwbiP1zhka1JNE7+ZKKmbsGP6FuYyKME7KF1j5JG3HAJiLzN7sXUWtO1LQR77g7PeDhOtcicdsAzyDMkEf3gUKEc3gvTijqqu6Uw6EkRWVDasCYwVO1pxLgOFAR4n+3GWtsUdc8jFtdpaOfxEkS6rXgLTwozhXOxUe2PGuoXsA9tU7DInB6g0TkSGOwRwDuVaVAkWxew1novc4hVE6qyVe3MUsJj/cxspEF95e~1")
		r.Headers.Set("Priority", "u=0, i")
		r.Headers.Set("Sec-Ch-Ua", "\"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"138\", \"Brave\";v=\"138\"")
		r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
		r.Headers.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-Site", "none")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Sec-Gpc", "1")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Accept", "application/json, text/plain, */*")
		r.Headers.Set("Origin", "https://www.nasdaq.com")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")
		r.Headers.Set("Referer", "https://www.nasdaq.com/")
	})

	c.OnResponse(func(r *colly.Response) {
		body = string(r.Body)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})

	c.Visit("https://api.nasdaq.com/api/screener/stocks?download=true")

	return []byte(body)
}
