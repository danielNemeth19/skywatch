package main

import (
	"flag"
	"os"
)

func main() {
	useDisk := flag.Bool("test", true, "If true, local json file will be used")
	skyApi := flag.Bool("skyapi", true, "If true, skyapi will be used")
	base := flag.String("url", "https://www.kayak.com", "Specifies base url")
	pathParam := flag.String("path", "s/horizon/exploreapi/elasticbox", "Specifies path parameter")
	airport := flag.String("airport", "BUD", "Specifies source airport")
	zoomLevel := flag.String("zl", "2", "Specifies zoom level")
	departDate := flag.String("ddate", "", "Specifies depart date")
	returnDate := flag.String("rdate", "", "Specifies return date")
	budget := flag.String("budget", "", "Specifies budget")
	stopsFilterActive := flag.String("sfa", "false", "Specifies stops filter active")
	flag.Parse()

	var parser Parser
	targetCities := []string{"Madrid", "Barcelona", "Lisbon", "Milano"}
	parser = Parser{targetCities: targetCities}

	if *useDisk == true {
		parser.client = LocalClient{}
	} else if *skyApi == true {
		skyClient := SkyScannerClient{
			url: urlParts{
				base:      "https://skyscanner-api.p.rapidapi.com",
				pathParam: "v3/flights/live/search/create",
			},
			rapidApiKey:  os.Getenv("rapidApiKey"),
			rapidApiHost: "skyscanner-api.p.rapidapi.com",
		}
		parser.client = skyClient
	} else {
		url := urlParts{
			base:              *base,
			pathParam:         *pathParam,
			airport:           *airport,
			zoomLevel:         *zoomLevel,
			departDate:        *departDate,
			returnDate:        *returnDate,
			budget:            *budget,
			stopsFilterActive: *stopsFilterActive,
		}
		parser.client = WebClient{url: url}
	}
	data := parser.client.getData()
	parser.summarize(data)
}
