package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
	budget := flag.String("budget", "1", "Specifies budget")
	stopsFilterActive := flag.String("sfa", "false", "Specifies stops filter active")
	flag.Parse()

	if *useDisk == true {
		data, err := os.ReadFile("output/test.json")
		if err != nil {
			panic(err)
		}
		var myStruct AirData
		if err := json.Unmarshal(data, &myStruct); err != nil {
			panic(err)
		}
		fmt.Println(len(myStruct.Destinations))
		for k, v := range myStruct.Destinations {
			if v.FlightInfo.PriceUSD != 0 {
				fmt.Printf("Item: %d\n", k)
				fmt.Printf("\tDestination: %s\n", v.City.Name)
				fmt.Printf("\tAirline: %s\n", v.Airline)
				fmt.Printf("\tPriceUSD: %d\n", v.FlightInfo.PriceUSD)
				fmt.Printf("\tPrice: %d\n", v.FlightInfo.Price)
			}
		}
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
