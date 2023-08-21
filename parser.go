package main

import (
	"fmt"
)

type Parser struct {
	client Client
}

func (p Parser) summarize(data AirData) {
	best := data.Content.SortingOptions.Best[0]
	fmt.Printf("Best score is -- score: %f -- itinerary id: %s\n", best.Score, best.ItineraryId)
	fmt.Printf("Details of best itinerary: %s\n", data.Content.Results.Itineraries[best.ItineraryId])

	for key, data := range data.Content.Results.Itineraries {
		fmt.Printf("Key %s\n", key)
		for j, pricingOption := range data.PricingOptions {
			fmt.Printf("\tpricing option #%d\n", j)
			for i, item := range pricingOption.Items {
				fmt.Printf("\t\titem #%d, agentID: %s\n", i, item.AgentId)
			}
		}
	}
	counter := 0
	bestId := best.ItineraryId
	for key := range data.Content.Results.Itineraries {
		if key == bestId {
			fmt.Printf("Found at: %d\n", counter)
		} else {
			counter += 1
		}
	}
}
