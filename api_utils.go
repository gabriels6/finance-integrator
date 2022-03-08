package main

import (
	"net/http"
	"io/ioutil"
)

func CallApi(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		return []byte("Failed to get response")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("Failed to get body")
	}
	return body
}