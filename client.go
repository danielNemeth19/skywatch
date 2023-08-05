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

const (
	ResultUnspecified = "RESULT_STATUS_UNSPECIFIED"
	ResultIncomplete  = "RESULT_STATUS_INCOMPLETE"
	ResultComplete    = "RESULT_STATUS_COMPLETE"
	ResultFailed      = "RESULT_STATUS_FAILED"
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
	apiKey string
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
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != err {
		log.Fatal(err)
	}
	return body
}

// Think about handling error results
// Think about the urlParts struct - should it really be embedded in client?
func (s SkyScannerClient) callUntilDone(body []byte) []byte {
	fmt.Println("CALLING INTO RECURSIVE FUNCTION")
	var sessionInfo SessionInfo
	if err := json.Unmarshal(body, &sessionInfo); err != nil {
		log.Fatal(err)
	}
	if sessionInfo.Status == ResultIncomplete {
		fmt.Printf("Status is incomplete: %s\n", sessionInfo.Status)
		url := urlParts{
			base:      s.urlParts.base,
			pathParam: "apiservices/v3/flights/live/search/poll/" + sessionInfo.Token,
		}
		newBody := s.sendRequest(http.MethodPost, url.Compose(), nil)
		return s.callUntilDone(newBody)
	}
	fmt.Printf("Status is NOT incomplete: %s\n", sessionInfo.Status)
	return body
}

func (s SkyScannerClient) getData() AirData {
	payload, err := json.Marshal(s.PayloadBuilder.Assemble())
	if err != nil {
		panic(err)
	}
	fmt.Printf("payload is: %s\n", payload)

	body := s.sendRequest(http.MethodPost, s.urlParts.Compose(), payload)
	fb := s.callUntilDone(body)

	var finalJson bytes.Buffer
	err = json.Indent(&finalJson, fb, "", "\t")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output/skyscanner_final.json", finalJson.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	return AirData{}
}
