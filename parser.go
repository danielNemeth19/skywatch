package main

import (
	"fmt"
)

type Parser struct {
	targetCities []string
	client       Client
}

func (p Parser) summarize(data AirData) {
	fmt.Printf("Summary for origin airport: %s\n", data.Origin)
	fmt.Println(len(data.Destinations))
}
