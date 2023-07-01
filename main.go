package main

import (
	"flag"
	"fmt"
	"go.zoe.im/surferua"
	"io"
	"net/http"
	"os"
)

func main() {
	base := flag.String("url", "https://www.kayak.com", "Specifies base url")
	pathParam := flag.String("path", "s/horizon/exploreapi/elasticbox", "Specifies path parameter")
	airport := flag.String("airport", "BUD", "Specifies source airport")
	zoomLevel := flag.String("zl", "2", "Specifies zoom level")
	departDate := flag.String("ddate", "", "Specifies depart date")
	returnDate := flag.String("rdate", "", "Specifies return date")
	budget := flag.String("budget", "1", "Specifies budget")
	flag.Parse()

	url := urlParts{
		base:       *base,
		pathParam:  *pathParam,
		airport:    *airport,
		zoomLevel:  *zoomLevel,
		departDate: *departDate,
		returnDate: *returnDate,
		budget:     *budget,
	}.Compose()

	req, err := http.NewRequest(http.MethodGet, url, nil)

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
	os.WriteFile("output/test.html", []byte(body), 0644)
}
