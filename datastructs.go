package main

import "time"

type AirData struct {
	Destinations []Destinations
	Origin       Origin
}

type Destinations struct{}

type Origin struct{}

type Payload struct {
	Query PayloadData `json:"query,omitempty"`
}

type PayloadData struct {
	Market     string      `json:"market,omitempty"`
	Locale     string      `json:"locale,omitempty"`
	Currency   string      `json:"currency,omitempty"`
	QueryLegs  []QueryLegs `json:"queryLegs,omitempty"`
	CabinClass string      `json:"cabinClass,omitempty"`
	Adults     int         `json:"adults,omitempty"`
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
