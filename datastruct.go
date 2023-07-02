package main

type AirData struct {
	Destinations []Destinations
	Origin       Origin
}

type Destinations struct {
	Airline    string     `json:"airline,omitempty"`
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
	Price           int `json:"price,omitempty"`
	PriceUSD        int `json:"priceUSD,omitempty"`
	HistoricalPrice int `json:"historicalPrice,omitempty"`
}

type City struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
