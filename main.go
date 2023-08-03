package main

import (
	"flag"
	"os"
)

func main() {
	useDisk := flag.Bool("test", false, "If true, local json file will be used")
	base := flag.String("url", "https://skyscanner-api.p.rapidapi.com", "Specifies base url")
	pathParam := flag.String("path", "v3/flights/live/search/create", "Specifies path parameter")
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
			rapidApiKey:  os.Getenv("rapidApiKey"),
			rapidApiHost: "skyscanner-api.p.rapidapi.com",
			urlParts: urlParts{
				base:      *base,
				pathParam: *pathParam,
			},
			PayloadBuilder: PayloadBuilder{
				origin:      *origin,
				destination: *destination,
				dateString:  "20231201",
			},
		}
		parser.client = skyClient
	}
	data := parser.client.getData()
	parser.summarize(data)
}
