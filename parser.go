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

func (p Parser) getOptionData(data AirData) []OptionData {
	var options []OptionData
	for id, itinerary := range data.Content.Results.Itineraries {
		for i, option := range itinerary.PricingOptions {
			var fares int
			for _, item := range option.Items {
				fares += len(item.Fares)
			}
			score := p.findBestScore(id, data.Content.SortingOptions.Best)
			od := OptionData{
				itineraryId: id,
				optionIndex: i,
				bestScore: score,
				price:       p.convertPrice(option.Price),
				numAgents:   len(option.AgentIds),
				numItems:    len(option.Items),
				numFares:    fares,
			}
			od.isDirect = od.isDirectFlight()
			options = append(options, od)
		}
	}
	sort.Slice(options, func(i, j int)bool {return options[i].bestScore > options[j].bestScore})
	return options
}

func (od OptionData) isDirectFlight() bool {
	if od.numAgents == 1 && od.numItems == 1 && od.numFares == 1 {
		return true
	}
	return false
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

func (p Parser) findBestScore(itinerary string, data []Best) float64{
	var score float64
	for _, v := range data {
		if v.ItineraryId == itinerary {
			score = v.Score
		}
	}
	return score
}
