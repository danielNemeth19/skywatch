package main

import (
	"fmt"
	"go.zoe.im/surferua"
	"io"
	"net/http"
	"os"
)

type Client struct {
	url urlParts
}

func (c Client) getData() {
	req, err := http.NewRequest(http.MethodGet, c.url.Compose(), nil)

	userAgent := surferua.New().Desktop().Chrome().String()
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", userAgent)

	if err != nil {
		panic("error")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("error")
	}
	fmt.Printf("Response code is: %d\n", res.StatusCode)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("error")
	}
	os.WriteFile("output/test.json", []byte(body), 0644)
}
