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
	payload      Payload
}

func (s SkyScannerClient) getData() AirData {
	payload, err := json.Marshal(s.payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(payload))

	req, err := http.NewRequest(http.MethodPost, s.url.Compose(), bytes.NewReader(payload))
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
