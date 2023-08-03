package main

import (
	"time"
)

type PayloadBuilder struct {
	market      string
	locale      string
	currency    string
	origin      string
	destination string
	dateString  string
}

func (builder PayloadBuilder) parseDateString() time.Time {
	date, err := time.Parse("20060102", builder.dateString)
	if err != nil {
		panic(err)
	}
	return date
}

func (builder PayloadBuilder) Assemble() Payload {
	date := builder.parseDateString()
	var payload Payload
	payload.Query = PayloadData{
		Market:   builder.market,
		Locale:   builder.locale,
		Currency: builder.currency,
		QueryLegs: []QueryLegs{
			{
				OriginPlaceId:      IataInfo{Iata: builder.origin},
				DestinationPlaceId: IataInfo{Iata: builder.destination},
				Date: DateInfo{
					Year: date.Year(), Month: date.Month(), Day: date.Day(),
				},
			},
		},
		CabinClass: "CABIN_CLASS_ECONOMY",
		Adults:     1,
	}
	return payload
}
