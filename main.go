package main

import (
	"flag"
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

func main() {
	useDisk := flag.Bool("test", false, "If true, local json file will be used")
	base := flag.String("url", "https://partners.api.skyscanner.net", "Specifies base url")
	pathParam := flag.String("path", "apiservices/v3/flights/live/search/create", "Specifies path parameter")
	market := flag.String("market", "UK", "Defines the market that the search is for")
	locale := flag.String("locale", "en-GB", "Locale that the results are returned in")
	currency := flag.String("currency", "GBP", "Currency that the search results are returned in")

	origin := flag.String("origin", "BUD", "Specifies source airport")
	destination := flag.String("destination", "", "Specifies destination airport")
	flag.Parse()

	var parser Parser
	targetCities := []string{"Madrid", "Barcelona", "Lisbon", "Milano"}
	parser = Parser{targetCities: targetCities}

	if *useDisk == true {
		parser.client = LocalClient{}
	} else {
		skyClient := SkyScannerClient{
			apiKey: getApiKey(),
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
				dateString:  "20231201",
			},
		}
		parser.client = skyClient
	}
	data := parser.client.getData()
	parser.summarize(data)
}
