package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Parser struct {
	data         AirData
	targetCities []string
}

func (p *Parser) getData() {
	data, err := os.ReadFile("output/test.json")
	if err != nil {
		panic(err)
	}
	var airData AirData
	if err := json.Unmarshal(data, &airData); err != nil {
		panic(err)
	}
	p.data = airData
}

func (p *Parser) summarize() {
	fmt.Println(len(p.data.Destinations))
	for k, v := range p.data.Destinations {
		for _, c := range p.targetCities {
			if v.City.Name == c {
				fmt.Printf("Item: %d\n", k)
				fmt.Printf("\tDestination: %s\n", v.City.Name)
				fmt.Printf("\tAirline: %s\n", v.Airline)
				fmt.Printf("\tDepart: %s\n", v.DepartDate)
				fmt.Printf("\tReturn: %s\n", v.ReturnDate)
				fmt.Printf("\tPriceUSD: %2.f\n", v.FlightInfo.PriceUSD)
				fmt.Printf("\tPrice: %2.f\n", v.FlightInfo.Price)
			}
		}
	}
}
