package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

type Parser struct {
	client Client
}

func (p Parser) summarize(data AirData) {
	best := data.Content.SortingOptions.Best[0]
	fmt.Printf("Best score is -- score: %f -- itinerary id: %s\n", best.Score, best.ItineraryId)
	bestItinerary := data.Content.Results.Itineraries[best.ItineraryId]
	for i, option := range bestItinerary.PricingOptions {
		fmt.Printf("Option %d\n", i)
		fmt.Printf("\tPrice: %f\n", p.convertPrice(option.Price))
		fmt.Printf("\tAgent ids: %s\n", option.AgentIds)
		fmt.Printf("\tItems: %s\n", option.Items)
	}
}

func (p Parser) checkTransfer(data AirData) {
	for key, value := range data.Content.Results.Itineraries {
		pOptions := len(value.PricingOptions)
		fmt.Printf("%s -- Number of pricing options: %d\n", key, pOptions)
		for i, option := range value.PricingOptions {
			fmt.Printf("%d -- Number of items: %d -- Number of fares in items: %d\n", i, len(option.Items), len(option.Items[0].Fares))
			fmt.Printf("Price: %f\n", p.convertPrice(option.Price))
		}
	}
}

func (p Parser) isDirect(data AirData) {
	numItineraries := len(data.Content.Results.Itineraries)
	counter := 0
	fmt.Printf("Number of Itineraries: %d\n", numItineraries)
	for key, value := range data.Content.Results.Itineraries {
		counter += 1
		for i, option := range value.PricingOptions {
			if len(option.AgentIds) == 1 && len(option.Items) == 1 && len(option.Items[0].Fares) == 1 {
				fmt.Printf("%d itinerary %s -- option %d - must be direct\n", counter, key, i)
				p.analyzePricingOption(option)
			}
			if len(option.AgentIds) != 1 || len(option.Items) != 1 {
				fmt.Printf("%d itinerary %s -- option %d - NOT direct -- agents: %d -- items: %d\n",counter, key, i, len(option.AgentIds), len(option.Items))
				p.analyzePricingOption(option)
			}
			if len(option.AgentIds) == 1 && len(option.Items) == 1 && len(option.Items[0].Fares) != 1 {
				fmt.Printf("%d itinerary %s -- option %d - NOT direct -- agents: %d -- items: %d\n", counter, key, i, len(option.AgentIds), len(option.Items))
				p.analyzePricingOption(option)
			}
		}
	}
}

func (p Parser) analyzePricingOption(po PricingOption) {
	var agents, items, fares int
	agents = len(po.AgentIds)
	items = len(po.Items)
	for _, data := range po.Items {
		fares += len(data.Fares)
	}
	fmt.Printf("Agents: %d, Items: %d, Total fares: %d\n", agents, items, fares)
}

func (p Parser) convertPrice(priceData Price) float64 {
	convertWith := priceUnitEnum[priceData.Unit]
	price, err := strconv.ParseFloat(priceData.Amount, 64)
	if err != nil {
		log.Println("No price data available")
		return 0
	}
	return price / convertWith
}

func (p Parser) summarizeAgents(data AirData) {
	var ratings []AgentRating
	fmt.Printf("agent: -- name -- rating\n")
	for key, data := range data.Content.Results.Agents {
		ratings = append(ratings, AgentRating{Id: key, Name: data.Name, Rating: data.Rating})
	}
	sort.Slice(ratings, func(i, j int) bool { return ratings[i].Rating > ratings[j].Rating })
	for _, agentRating := range ratings {
		fmt.Printf("Agent: %s -- rating %f -- Id: %s\n", agentRating.Name, agentRating.Rating, agentRating.Id)
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
