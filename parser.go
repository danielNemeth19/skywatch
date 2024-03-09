package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"time"
)

const (
	Departure = "Departure"
	Arrival   = "Arrival"
)

type Parser struct {
	data AirData
}

func (p Parser) getFlightData(maxStops int, minScore float64) FlightData {
	options := p.getOptionData(maxStops, minScore)
	return FlightData{
		Payload: p.data.Payload,
		Options: options,
	}
}

// Function to check that within a pricing option we have as many items as number of agents
func (p Parser) counter() {
	pass, fail := 0, 0
	for _, itinerary := range p.data.Content.Results.Itineraries {
		for _, option := range itinerary.PricingOptions {
			numberAgents := len(option.AgentIds)
			numberItems := len(option.Items)
			if numberAgents == numberItems {
				pass++
			} else {
				fail++
			}
		}
	}
	fmt.Printf("Passes: %d -- Failes: %d\n", pass, fail)
}

func (p Parser) setSegmentPrices(items []Item, price float64) map[string]float64 {
	priceMap := make(map[string]float64)
	for _, item := range items {
		itemPrice := p.convertPrice(item.Price)
		for _, fare := range item.Fares {
			if itemPrice != price {
				priceMap[fare.SegmentId] = itemPrice
			} else {
				priceMap[fare.SegmentId] = 0
			}
		}
	}
	return priceMap
}

func (p Parser) getOptionData(maxStops int, minScore float64) []OptionData {
	var options []OptionData

	for id, itinerary := range p.data.Content.Results.Itineraries {
		for i, option := range itinerary.PricingOptions {
			score := p.findBestScore(id)

			if score < float64(minScore) {
				continue
			}

			legData := p.data.Content.Results.Legs[id]
			price := p.convertPrice(option.Price)
			segPriceMap := p.setSegmentPrices(option.Items, price)
			segmentDetails := p.collectSegmentDetails(legData.SegmentIds, segPriceMap)
			totalFlightTime := p.calculateFlightTime(segmentDetails)

			// Think about this.. aren't pricing options only different in price..?
			if legData.StopCount <= maxStops {
				od := OptionData{
					ItineraryId:     id,
					OptionIndex:     i,
					Score:           score,
					Price:           price,
					SegmentDetails:  segmentDetails,
					IsDirect:        p.isDirectFlight(legData),
					NumAgents:       len(option.AgentIds),
					NumItems:        len(option.Items),
					NumFares:        len(legData.SegmentIds),
					TotalTravelTime: legData.DurationInMinutes,
					TotalFlightTime: totalFlightTime,
					TotalTransitTime: legData.DurationInMinutes - totalFlightTime,
				}
				options = append(options, od)
			} else {
				log.Printf("Skipping as %d > %d - %s\n", legData.StopCount, maxStops, id)
			}
		}
	}
	sort.Slice(options, func(i, j int) bool {
		switch {
		case options[i].Score != options[j].Score:
			return options[i].Score > options[j].Score
		default:
			return options[i].Price < options[j].Price
		}
	})
	printResult(options)
	return options
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

func (p Parser) collectSegmentDetails(segmentIds []string, segPriceMap map[string]float64) []SegmentData {
	var sg []SegmentData
	for _, segmentId := range segmentIds {
		segment := p.data.Content.Results.Segments[segmentId]
		var op, dp []string
		departure := segment.createDateTime(Departure)
		arrival := segment.createDateTime(Arrival)
		segmentData := SegmentData{
			Price:             segPriceMap[segmentId],
			OriginPlaces:      p.findPlace(segment.OriginId, op),
			DestinationPlaces: p.findPlace(segment.DestinationId, dp),
			DepartAt:          departure.Format("2006 Jan 2"),
			DepartTime:        departure.Format("15:04"),
			ArriveAt:          arrival.Format("2006 Jan 2"),
			ArriveTime:        arrival.Format("15:04"),
			DurationInMinutes: segment.DurationInMinutes,
			MarketingCarrier:  p.data.Content.Results.Carriers[segment.MarketingCarrierId].Name,
		}
		sg = append(sg, segmentData)
	}
	return sg
}

func (p Parser) calculateFlightTime(segmentData []SegmentData) int {
	var totalTime int
	for _, segment := range segmentData {
		totalTime += segment.DurationInMinutes
	}
	return totalTime
}

func printResult(options []OptionData) {
	for i, data := range options {
		fmt.Printf("\n%d --%s\n", i, data.ItineraryId)
		fmt.Printf("Price: %f Score: (%f) Direct: %v\n", data.Price, data.Score, data.IsDirect)
		fmt.Printf("Total flight time: %d Total travel time: %d Total transit time: %d\n", data.TotalFlightTime, data.TotalTravelTime, data.TotalTransitTime)
		for _, s := range data.SegmentDetails {
			fmt.Printf("Departure:\n\tFrom:%v\n\tTime: %s\n", s.OriginPlaces, s.DepartAt)
			fmt.Printf("Arrival:\n\tTo:%v\n\tTime: %s\n", s.DestinationPlaces, s.ArriveAt)
			fmt.Printf("Duration: %d\n", s.DurationInMinutes)
			fmt.Printf("Carrier: %s\n", s.MarketingCarrier)
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
		return 0
	}
	return math.Round(price / convertWith)
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
