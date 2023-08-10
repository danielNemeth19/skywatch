package main

import (
	"fmt"
)

type Parser struct {
	targetCities []string
	client       Client
}

func (p Parser) summarize(data AirData) {
	fmt.Printf("parsed: %v\n", data)
}
