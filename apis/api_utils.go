package apis

import (
	"bytes"
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

func PostApi(url string, body string) []byte {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte("Failed to get response")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte("Failed to get body")
	}
	return respBody
}
