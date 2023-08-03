package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type SkyScannerClient struct {
	rapidApiKey  string
	rapidApiHost string
	urlParts
	PayloadBuilder
}

func (s SkyScannerClient) getData() AirData {
	payload, err := json.Marshal(s.PayloadBuilder.Assemble())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(payload))

	req, err := http.NewRequest(http.MethodPost, s.urlParts.Compose(), bytes.NewReader(payload))
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

	var token SessionToken
	if err := json.Unmarshal(body, &token); err != nil {
		panic(err)
	}

	url := urlParts{
		base:      "https://skyscanner-api.p.rapidapi.com",
		pathParam: "v3/flights/live/search/poll/" + token.Token,
	}
	req, err = http.NewRequest(http.MethodPost, url.Compose(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-RapidAPI-Key", s.rapidApiKey)
	req.Header.Add("X-RapidAPI-Host", s.rapidApiHost)
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != err {
		log.Fatal(err)
	}

	var finalJson bytes.Buffer
	err = json.Indent(&finalJson, body, "", "\t")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output/skyscanner_final.json", finalJson.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	return AirData{}
}
