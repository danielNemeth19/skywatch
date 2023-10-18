package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	ResultUnspecified = "RESULT_STATUS_UNSPECIFIED"
	ResultIncomplete  = "RESULT_STATUS_INCOMPLETE"
	ResultFailed      = "RESULT_STATUS_FAILED"
)

type Client interface {
	getData() AirData
}

type LocalClient struct {
	fileName string
}

func (l LocalClient) getData() AirData {
	fileToRead := "output/" + l.fileName
	data, err := os.ReadFile(fileToRead)
	if err != nil {
		log.Fatalf("Check input filename:%v\n", err)
	}
	var airData AirData
	if err := json.Unmarshal(data, &airData); err != nil {
		panic(err)
	}
	return airData
}

type SkyScannerClient struct {
	apiKey   string
	fileName string
	retries  int
	urlParts
	PayloadBuilder
}

func (s SkyScannerClient) sendRequest(method string, url string, payload []byte) []byte {
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-api-key", s.apiKey)
	time.Sleep(2 * time.Second)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Printf("Response status code is %d\n", res.StatusCode)
	if res.StatusCode == http.StatusTooManyRequests {
		log.Fatal("Too many requests")
	}
	body, err := io.ReadAll(res.Body)
	if err != err {
		log.Fatal(err)
	}
	return body
}

func (s SkyScannerClient) PollUntilCompletes(body []byte) []byte {
	var sessionInfo SessionInfo
	if err := json.Unmarshal(body, &sessionInfo); err != nil {
		log.Fatal(err)
	}
	if sessionInfo.Status == "" {
		fmt.Printf("Body is: %s\n", body)
	}
	if sessionInfo.Status == ResultUnspecified || sessionInfo.Status == ResultFailed {
		log.Fatalf("Aborting as status is %s", sessionInfo.Status)
	}
	log.Printf("Status is %s -- number of retries: %d\n", sessionInfo.Status, s.retries)
	if sessionInfo.Status == ResultIncomplete {
		s.retries += 1
		url := urlParts{
			base:      s.urlParts.base,
			pathParam: "apiservices/v3/flights/live/search/poll/" + sessionInfo.Token,
		}
		newBody := s.sendRequest(http.MethodPost, url.Compose(), nil)
		return s.PollUntilCompletes(newBody)
	}
	return body
}


func (s SkyScannerClient) StoreResult(data AirData) {
	marshaled, err := json.Marshal(data)

	var jsonResult bytes.Buffer
	err = json.Indent(&jsonResult, marshaled, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fileToWrite := "output/" + s.fileName
	err = os.WriteFile(fileToWrite, jsonResult.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (s SkyScannerClient) getData() AirData {
	payload, err := json.Marshal(s.PayloadBuilder.Assemble())
	if err != nil {
		panic(err)
	}
	fmt.Printf("payload is: %s\n", payload)

	body := s.sendRequest(http.MethodPost, s.urlParts.Compose(), payload)
	fb := s.PollUntilCompletes(body)

	var airData AirData
	if err := json.Unmarshal(fb, &airData); err != nil {
		panic(err)
	}
	airData.PayloadBuilder = s.PayloadBuilder
	s.StoreResult(airData)
	return airData
}
