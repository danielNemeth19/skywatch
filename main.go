package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func getApiKey() string {
	apiKey := os.Getenv("apikey")
	if apiKey == "" {
		log.Fatalf("Api key not found!")
	}
	return apiKey
}

var agents string

func main() {
	useDisk := flag.Bool("test", false, "If true, local json file will be used")
	fileName := flag.String("fn", "skyscanner_final.json", "Name of the file to be loaded or written")
	base := flag.String("url", "https://partners.api.skyscanner.net", "Specifies base url")
	pathParam := flag.String("path", "apiservices/v3/flights/live/search/create", "Specifies path parameter")
	market := flag.String("market", "UK", "Defines the market that the search is for")
	locale := flag.String("locale", "en-GB", "Locale that the results are returned in")
	currency := flag.String("currency", "GBP", "Currency that the search results are returned in")
	startDate := flag.String("sdate", "20231201", "Start date of travel in YYYYMMDD format")
	origin := flag.String("origin", "BUD", "Specifies source airport")
	destination := flag.String("destination", "", "Specifies destination airport")
	budgetAgent := flag.Bool("budget", true, "If true, only budget airlines will be queried")
	flag.Parse()

	var parser Parser
	parser = Parser{}

	if *useDisk == true {
		parser.client = LocalClient{fileName: *fileName}
	} else {
		skyClient := SkyScannerClient{
			apiKey:   getApiKey(),
			fileName: *fileName,
			urlParts: urlParts{
				base:      *base,
				pathParam: *pathParam,
			},
			PayloadBuilder: PayloadBuilder{
				origin:      *origin,
				destination: *destination,
				market:      *market,
				locale:      *locale,
				currency:    *currency,
				sDate:       *startDate,
				budgetAgent: *budgetAgent,
			},
		}
		parser.client = skyClient
	}
	data := parser.client.getData()
	options := parser.getOptionData(data)
	for _, option := range options {
		fmt.Printf("Score: %f -- price: %.2f -- direct: %v\n", option.bestScore, option.price, option.isDirect)
	}
	parser.checkLegIds(data)
}
