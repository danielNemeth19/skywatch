package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.zoe.im/surferua"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Client interface {
	getData() AirData
}

type LocalClient struct{}

func (l LocalClient) getData() AirData {
	data, err := os.ReadFile("output/test.json")
	if err != nil {
		panic(err)
	}
	var airData AirData
	if err := json.Unmarshal(data, &airData); err != nil {
		panic(err)
	}
	return airData
}

type WebClient struct {
	url urlParts
}

func (c WebClient) getData() AirData {
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
	os.WriteFile("output/test.json", body, 0644)
	var airData AirData
	if err := json.Unmarshal(body, &airData); err != nil {
		panic(err)
	}
	return airData
}

type SkyScannerClient struct {
	url          urlParts
	rapidApiKey  string
	rapidApiHost string
}

func (s SkyScannerClient) getData() AirData {
	payload := strings.NewReader(
		`{
		"query":
			{
				"market": "UK",
				"locale": "en-GB",
				"currency": "EUR",
				"queryLegs":
					[
						{
							"originPlaceId": {"iata": "BUD"},
							"destinationPlaceId": {"iata": "LIS"},
							"date": {"year": 2023, "month": 11, "day": 1}
						}
					],
				"cabinClass": "CABIN_CLASS_ECONOMY",
				"adults": 1
			}
	}`)
	req, err := http.NewRequest(http.MethodPost, s.url.Compose(), payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", s.rapidApiKey)
	req.Header.Add("X-RapidAPI-Host", s.rapidApiHost)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var myJson bytes.Buffer
	err = json.Indent(&myJson, body, "", "\t")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output/skyscanner.json", myJson.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	return AirData{}
}
