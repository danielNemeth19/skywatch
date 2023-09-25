package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

type Parser struct {
	data AirData
}

func (p Parser) checkPlaces() {
	for id, data := range p.data.Content.Results.Segments {
		var ps []string
		fmt.Printf("id: %s, data: %v\n", id, data)
		placeStack := p.findPlace(data.OriginId, ps)
		fmt.Printf("Places: %#v\n", placeStack)
	}
}

func (p Parser) findPlace(id string, ps []string) []string {
	place := p.data.Content.Results.Places[id]
	ps = append(ps, place.Name)
	parentID := place.ParentID
	if parentID == "" {
		return ps
	}
	return p.findPlace(parentID, ps)
}

func (p Parser) getOptionData() []OptionData {
	var options []OptionData
	for id, itinerary := range p.data.Content.Results.Itineraries {
		for i, option := range itinerary.PricingOptions {
			score := p.findBestScore(id)
			legData := p.data.Content.Results.Legs[id]
			od := OptionData{
				itineraryId:    id,
				optionIndex:    i,
				score:          score,
				segmentDetails: p.collectSegmentDetails(legData.SegmentIds),
				price:          p.convertPrice(option.Price),
				isDirect:       p.isDirectFlight(legData),
				numAgents:      len(option.AgentIds),
				numItems:       len(option.Items),
				numFares:       len(legData.SegmentIds),
			}
			options = append(options, od)
		}
	}
	sort.Slice(options, func(i, j int) bool { return options[i].score > options[j].score })

	printResult(options)
	return options
}

func (p Parser) collectSegmentDetails(segmentIds []string) []SegmentData {
	var sg []SegmentData
	for _, segmentId := range segmentIds {
		segment := p.data.Content.Results.Segments[segmentId]
		var op, dp []string
		departure := time.Date(
			segment.DepartureDateTime.Year,
			time.Month(segment.DepartureDateTime.Month),
			segment.DepartureDateTime.Day,
			segment.DepartureDateTime.Hour,
			segment.DepartureDateTime.Minute,
			0, 0, time.Local,
		)
		arrival := time.Date(
			segment.ArrivalDateTime.Year,
			time.Month(segment.ArrivalDateTime.Month),
			segment.ArrivalDateTime.Day,
			segment.ArrivalDateTime.Hour,
			segment.ArrivalDateTime.Minute,
			0, 0, time.Local,
		)
		segmentData := SegmentData{
			OriginPlaces:      p.findPlace(segment.OriginId, op),
			DestinationPlaces: p.findPlace(segment.DestinationId, dp),
			Departure:         departure,
			Arrival:           arrival,
			DurationInMinutes: segment.DurationInMinutes,
			MarketingCarrierId: p.data.Content.Results.Carriers[segment.MarketingCarrierId].Name,
		}
		sg = append(sg, segmentData)
	}
	return sg
}

func printResult(options []OptionData) {
	for i, data := range options {
		fmt.Printf("\n%d --%s\n", i, data.itineraryId)
		fmt.Printf("Price: %f score (%f) direct: %v\n", data.price, data.score, data.isDirect)
		for _, s := range data.segmentDetails {
			fmt.Printf("Departure:\n\tFrom:%v\n\tTime: %s\n", s.OriginPlaces, s.Departure)
			fmt.Printf("Arrival:\n\tFrom:%v\n\tTime: %s\n", s.DestinationPlaces, s.Arrival)
			fmt.Printf("Carrier:\n\t%s\n", s.MarketingCarrierId)
		}
	}
}

func (p Parser) isDirectFlight(legData Leg) bool {
	if legData.StopCount == 0 {
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

func (p Parser) findBestScore(itinerary string) float64 {
	var score float64
	for _, v := range p.data.Content.SortingOptions.Best {
		if v.ItineraryId == itinerary {
			score = v.Score
			break
		}
	}
	return score
}
