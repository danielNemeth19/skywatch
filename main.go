package main

import (
	"flag"
	"os"
)

func main() {
	useDisk := flag.Bool("test", true, "If true, local json file will be used")
	base := flag.String("url", "https://www.kayak.com", "Specifies base url")
	pathParam := flag.String("path", "s/horizon/exploreapi/elasticbox", "Specifies path parameter")
	airport := flag.String("airport", "BUD", "Specifies source airport")
	zoomLevel := flag.String("zl", "2", "Specifies zoom level")
	departDate := flag.String("ddate", "", "Specifies depart date")
	returnDate := flag.String("rdate", "", "Specifies return date")
	budget := flag.String("budget", "", "Specifies budget")
	stopsFilterActive := flag.String("sfa", "false", "Specifies stops filter active")
	flag.Parse()

	if *useDisk == true {
		var testParser Parser
		testParser = Parser{targetCities: []string{"Madrid", "Barcelona", "Lisbon", "Milano"}}
		testParser.getData()
		testParser.summarize()
		os.Exit(0)
	}

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
	client := Client{url: url}
	client.getData()
}
