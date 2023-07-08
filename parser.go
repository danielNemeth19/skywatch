package main

import (
	"fmt"
)

type Parser struct {
	targetCities []string
	client       Client
}

func (p Parser) summarize(data AirData) {
	fmt.Println(len(data.Destinations))
	for k, v := range data.Destinations {
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
