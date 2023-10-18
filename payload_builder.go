package main

import (
	"time"
)

type PayloadBuilder struct {
	Market      string `json:"market,omitempty"`
	Locale      string `json:"locale,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Destination string `json:"destination,omitempty"`
	SDate       string` json:"sDate,omitempty"`
	BudgetAgent bool
}

func (builder PayloadBuilder) parseDateString() time.Time {
	date, err := time.Parse("20060102", builder.SDate)
	if err != nil {
		panic(err)
	}
	return date
}

func (builder PayloadBuilder) setIncludedAgentIds() []string {
	if builder.BudgetAgent == true {
		return []string{"wizz", "ryan", "easy", "eapr"}
	}
	return []string{}
}

func (builder PayloadBuilder) Assemble() Payload {
	date := builder.parseDateString()
	var payload Payload
	payload.Query = PayloadData{
		Market:   builder.Market,
		Locale:   builder.Locale,
		Currency: builder.Currency,
		QueryLegs: []QueryLegs{
			{
				OriginPlaceId:      IataInfo{Iata: builder.Origin},
				DestinationPlaceId: IataInfo{Iata: builder.Destination},
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
