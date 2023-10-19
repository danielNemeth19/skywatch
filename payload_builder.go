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
	sDate       string
	budgetAgent bool
}

func (builder PayloadBuilder) parseDateString() time.Time {
	date, err := time.Parse("20060102", builder.sDate)
	if err != nil {
		panic(err)
	}
	return date
}

func (builder PayloadBuilder) setIncludedAgentIds() []string {
	if builder.budgetAgent == true {
		return []string{"wizz", "ryan", "easy", "eapr"}
	}
	return []string{}
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
		CabinClass:        "CABIN_CLASS_ECONOMY",
		Adults:            1,
		IncludedAgentsIds: builder.setIncludedAgentIds(),
	}
	return payload
}
