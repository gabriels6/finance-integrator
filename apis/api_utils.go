package apis

import (
	"io"
	"net/http"
)

func CallApi(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		return []byte("Failed to get response")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte("Failed to get body")
	}
	return body
}
