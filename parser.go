package main

import (
	"fmt"
	"sort"
)

type Parser struct {
	client Client
}

func (p Parser) summarize(data AirData) {
	best := data.Content.SortingOptions.Best[0]
	fmt.Printf("Best score is -- score: %f -- itinerary id: %s\n", best.Score, best.ItineraryId)
	bestItinerary := data.Content.Results.Itineraries[best.ItineraryId]
	fmt.Printf("LegId: %s\n", bestItinerary.LegIds)
	fmt.Printf("SustainabilityData: %s\n", bestItinerary.SustainabilityData)
	fmt.Println("Pricing options are below")
	for i, option := range bestItinerary.PricingOptions {
		fmt.Printf("Option %d\n", i)
		fmt.Printf("\tPrice: %s\n", option.Price)
		fmt.Printf("\tAgent ids: %s\n", option.AgentIds)
		fmt.Printf("\tItems: %s\n", option.Items)
		fmt.Printf("\tTransfer type: %s\n", option.TransferType)
		fmt.Printf("\tId: %s\n", option.Id)
		fmt.Printf("\tPricing option fare: %s\n", option.PricingOptionFare)
	}
}


func (p Parser) summarizeAgents(data AirData) {
	var ratings []AgentRating
	fmt.Printf("agent: -- name -- rating\n")
	for _, data := range data.Content.Results.Agents {
		ratings = append(ratings, AgentRating{Name: data.Name, Rating: data.Rating})
	}
	sort.Slice(ratings, func(i, j int) bool { return ratings[i].Rating > ratings[j].Rating })
	for _, agentRating := range(ratings) {
		fmt.Printf("Agent: %s -- rating %f\n", agentRating.Name, agentRating.Rating)
	}
 }

func (p Parser) findBestItinerary(data AirData) {
	best := data.Content.SortingOptions.Best[0]
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
