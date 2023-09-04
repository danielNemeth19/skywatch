package main

import "time"

type AirData struct {
	Content Content `json:"content,omitempty"`
}

type Content struct {
	Results        Results        `json:"results,omitempty"`
	SortingOptions SortingOptions `json:"sortingOptions,omitempty"`
}

type Results struct {
	Itineraries map[string]Itinerary `json:"itineraries,omitempty"`
	Agents      map[string]Agent     `json:"agents"`
}

type Agent struct {
	Name                 string          `json:"name,omitempty"`
	Type                 string          `json:"type,omitempty"`
	ImageUrl             string          `json:"ImageUrl,omitempty"`
	FeedbackCount        int             `json:"feedbackCount,omitempty"`
	Rating               float32         `json:"rating,omitempty"`
	RatingBreakDown      RatingBreakDown `json:"ratingBreakDown,omitempty"`
	IsOptimizedForMobile bool            `json:"isOptimizedForMobile,omitempty"`
}

type AgentRating struct {
	Id     string
	Name   string
	Rating float32
}

type RatingBreakDown struct {
	CustomerService int `json:"customerService,omitempty"`
	ReliablePrices  int `json:"reliablePrices,omitempty"`
	ClearExtraFees  int `json:"clearExtraFees,omitempty"`
	EaseOfBooking   int `json:"easeOfBooking,omitempty"`
	Other           int `json:"other,omitempty"`
}

type SortingOptions struct {
	Best []Best `json:"best,omitempty"`
}

type Best struct {
	Score       float32 `json:"score,omitempty"`
	ItineraryId string  `json:"itineraryId,omitempty"`
}

type Itinerary struct {
	PricingOptions     []PricingOption `json:"pricingOptions,omitempty"`
	LegIds             []string        `json:"legIds,omitempty"`
	SustainabilityData string          `json:"sustainabilitydata,omitempty"`
}

type PricingOption struct {
	Price             Price    `json:"price,omitempty"`
	AgentIds          []string `json:"agentIds,omitempty"`
	Items             []Item   `json:"items,omitempty"`
	TransferType      string   `json:"transferType,omitempty"`
	Id                string   `json:"id,omitempty"`
	PricingOptionFare string   `json:"pricingOptionFare,omitempty"`
}

type Price struct {
	Amount       string `json:"amount,omitempty"`
	Unit         string `json:"unit,omitempty"`
	UpdateStatus string `json:"updateStatus,omitempty"`
}

type Item struct {
	Price    Price  `json:"price,omitempty"`
	AgentId  string `json:"agentId,omitempty"`
	DeepLink string `json:"deepLink,omitempty"`
	Fares    []Fare `json:"fares,omitempty"`
}

type Fare struct {
	SegmentId     string `json:"segmentId,omitempty"`
	BookingCode   string `json:"bookingCode,omitempty"`
	FareBasisCode string `json:"fareBasisCode,omitempty"`
}

type Payload struct {
	Query PayloadData `json:"query,omitempty"`
}

type PayloadData struct {
	Market            string      `json:"market,omitempty"`
	Locale            string      `json:"locale,omitempty"`
	Currency          string      `json:"currency,omitempty"`
	QueryLegs         []QueryLegs `json:"queryLegs,omitempty"`
	CabinClass        string      `json:"cabinClass,omitempty"`
	Adults            int         `json:"adults,omitempty"`
	IncludedAgentsIds []string    `json:"includedAgentsIds,omitempty"`
}

type QueryLegs struct {
	OriginPlaceId      IataInfo `json:"originPlaceId,omitempty"`
	DestinationPlaceId IataInfo `json:"destinationPlaceId,omitempty"`
	Date               DateInfo `json:"date,omitempty"`
}

type IataInfo struct {
	Iata string `json:"iata,omitempty"`
}

type DateInfo struct {
	Year  int        `json:"year,omitempty"`
	Month time.Month `json:"month,omitempty"`
	Day   int        `json:"day,omitempty"`
}

type SessionInfo struct {
	Status string `json:"status,omitempty"`
	Token  string `json:"sessionToken,omitempty"`
}
