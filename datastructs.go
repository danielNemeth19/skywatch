package main

import "time"

type AirData struct {
	Destinations []Destinations
	Origin       Origin
}

type Destinations struct {
	Airline    string     `json:"airline,omitempty"`
	DepartDate string     `json:"departd,omitempty"`
	ReturnDate string     `json:"returnd,omitempty"`
	FlightInfo FlightInfo `json:"flightInfo,omitempty"`
	Airport    AirPort    `json:"airport,omitempty"`
	City       City       `json:"city,omitempty"`
}

type Origin struct {
	Name      string  `json:"name,omitempty"`
	ShortName string  `json:"shortName,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	CityName  string  `json:"cityName,omitempty"`
}

type AirPort struct {
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
	Name       string  `json:"name,omitempty"`
	Popularity int     `json:"popularity,omitempty"`
	ShortName  string  `json:"shortName,omitempty"`
}

type FlightInfo struct {
	Price           float32 `json:"price,omitempty"`
	PriceUSD        float32 `json:"priceUSD,omitempty"`
	HistoricalPrice int     `json:"historicalPrice,omitempty"`
}

type City struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

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

type SessionToken struct {
	Token string `json:"sessionToken,omitempty"`
}
