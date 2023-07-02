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
		fmt.Println(myStruct.Destinations[0].Airport)
		os.Exit(0)
	}

	url := urlParts{
		base:       *base,
		pathParam:  *pathParam,
		airport:    *airport,
		zoomLevel:  *zoomLevel,
		departDate: *departDate,
		returnDate: *returnDate,
		budget:     *budget,
	}
	client := Client{url: url}
	client.getData()
}
