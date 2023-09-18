package main

import "time"

var priceUnitEnum = map[string]float64{
	"PRICE_UNIT_WHOLE": 1,
	"PRICE_UNIT_CENTI": 100,
	"PRICE_UNIT_MILLI": 1000,
	"PRICE_UNIT_MICRO": 1000000,
}

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
	Legs        map[string]Leg       `json:"legs,omitempty"`
	Segments    map[string]Segment   `json:"segments,omitempty"`
	Places      map[string]Place     `json:"places,omitempty"`
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

type Leg struct {
	OriginId          string   `json:"originPlaceId,omitempty"`
	DestinationId     string   `json:"destinationPlaceId,omitempty"`
	DepartureDateTime DateTime `json:"departureDateTime,omitempty"`
	ArrivalDateTime   DateTime `json:"arrivalDateTime,omitempty"`
	DurationInMinutes int      `json:"durationInMinutes,omitempty"`
	StopCount         int      `json:"stopCount,omitempty"`
	SegmentIds        []string `json:"segmentIds,omitempty"`
}

type Segment struct {
	OriginId          string   `json:"originPlaceId,omitempty"`
	DestinationId     string   `json:"destinationPlaceId,omitempty"`
	DepartureDateTime DateTime `json:"departureDateTime,omitempty"`
	ArrivalDateTime   DateTime `json:"arrivalDateTime,omitempty"`
	DurationInMinutes int      `json:"durationInMinutes,omitempty"`
}

type Place struct {
	EntityID    string `json:"entityId,omitempty"`
	ParentID    string `json:"parentId,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Iata        string `json:"iata,omitempty"`
	Coordinates string `json:"coordinates,omitempty"`
}

type DateTime struct {
	Year   int `json:"year,omitempty"`
	Month  int `json:"month,omitempty"`
	Day    int `json:"day,omitempty"`
	Hour   int `json:"hour,omitempty"`
	Minute int `json:"minute,omitempty"`
	Second int `json:"second,omitempty"`
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
	Score       float64 `json:"score,omitempty"`
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

type OptionData struct {
	itineraryId    string
	optionIndex    int
	price          float64
	segmentDetails []SegmentData
	numAgents      int
	numItems       int
	numFares       int
	isDirect       bool
	score          float64
}

type SegmentData struct {
	OriginPlaces      []string
	DestinationPlaces []string
	Departure         time.Time
	Arrival           time.Time
	DurationInMinutes int
}
