package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

const (
	Departure = "Departure"
	Arrival  = "Arrival"
)

type Parser struct {
	data AirData
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
			od := OptionData {
				ItineraryId:    id,
				OptionIndex:    i,
				Score:          score,
				SegmentDetails: p.collectSegmentDetails(legData.SegmentIds),
				Price:          p.convertPrice(option.Price),
				IsDirect:       p.isDirectFlight(legData),
				NumAgents:      len(option.AgentIds),
				NumItems:       len(option.Items),
				NumFares:       len(legData.SegmentIds),
			}
			options = append(options, od)
		}
	}
	sort.Slice(options, func(i, j int) bool { return options[i].Score > options[j].Score })

	printResult(options)
	return options
}


func (s Segment) createDateTime(direction string) time.Time {
	switch direction {
	case Departure:
		return time.Date(
			s.DepartureDateTime.Year,
			time.Month(s.DepartureDateTime.Month),
			s.DepartureDateTime.Day,
			s.DepartureDateTime.Hour,
			s.DepartureDateTime.Minute,
			0, 0, time.Local,
		)
	default:
		return time.Date(
			s.ArrivalDateTime.Year,
			time.Month(s.ArrivalDateTime.Month),
			s.ArrivalDateTime.Day,
			s.ArrivalDateTime.Hour,
			s.ArrivalDateTime.Minute,
			0, 0, time.Local,
		)
	}
}

func (p Parser) collectSegmentDetails(segmentIds []string) []SegmentData {
	var sg []SegmentData
	for _, segmentId := range segmentIds {
		segment := p.data.Content.Results.Segments[segmentId]
		var op, dp []string
		departure := segment.createDateTime(Departure)
		arrival := segment.createDateTime(Arrival)
		segmentData := SegmentData{
			OriginPlaces:       p.findPlace(segment.OriginId, op),
			DestinationPlaces:  p.findPlace(segment.DestinationId, dp),
			DepartAt:           departure,
			ArriveAt:           arrival,
			DurationInMinutes:  segment.DurationInMinutes,
			MarketingCarrierId: p.data.Content.Results.Carriers[segment.MarketingCarrierId].Name,
		}
		sg = append(sg, segmentData)
	}
	return sg
}

func printResult(options []OptionData) {
	for i, data := range options {
		fmt.Printf("\n%d --%s\n", i, data.ItineraryId)
		fmt.Printf("Price: %f score (%f) direct: %v\n", data.Price, data.Score, data.IsDirect)
		for _, s := range data.SegmentDetails {
			fmt.Printf("Departure:\n\tFrom:%v\n\tTime: %s\n", s.OriginPlaces, s.DepartAt.Format(time.DateTime))
			fmt.Printf("Arrival:\n\tFrom:%v\n\tTime: %s\n", s.DestinationPlaces, s.ArriveAt.Format(time.DateTime))
			fmt.Printf("Duration: %d\n", s.DurationInMinutes)
			fmt.Printf("Carrier: %s\n", s.MarketingCarrierId)
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
