package main

type AirData struct {
	Destinations []Destinations
	Origin       Origin
}

type Destinations struct {
	Airline string  `json:"airline,omitempty"`
	Airport AirPort `json:"airport"`
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
