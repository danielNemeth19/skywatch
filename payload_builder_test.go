package main

import (
	"skywatch.danielnemeth.net/internal/assert"
	"testing"
)

func TestPayloadBuilder_parseDateString(t *testing.T) {
	pb := PayloadBuilder{dateString: "20230326"}
	time := pb.parseDateString()
	assert.Equal(t, time.Year(), 2023)
	assert.Equal(t, time.Month(), 3)
	assert.Equal(t, time.Day(), 26)
}

func TestPayloadBuilder_Base_Assemble(t *testing.T) {
	pb := PayloadBuilder{
		market:      "HU",
		locale:      "hu-HU",
		currency:    "HUF",
		origin:      "BUD",
		destination: "MAD",
		dateString:  "20230326",
	}
	data := pb.Assemble()
	assert.Equal(t, data.Query.Market, "HU")
	assert.Equal(t, data.Query.Locale, "hu-HU")
	assert.Equal(t, data.Query.Currency, "HUF")
}

func TestPayloadBuilder_QueryLegs_Assemble(t *testing.T) {
	pb := PayloadBuilder{
		origin:      "BUD",
		destination: "MAD",
		dateString:  "20230326",
	}
	data := pb.Assemble()
	assert.Equal(t, data.Query.QueryLegs[0].OriginPlaceId.Iata, "BUD")
	assert.Equal(t, data.Query.QueryLegs[0].DestinationPlaceId.Iata, "MAD")
	assert.Equal(t, data.Query.QueryLegs[0].Date.Year, 2023)
	assert.Equal(t, data.Query.QueryLegs[0].Date.Month, 3)
	assert.Equal(t, data.Query.QueryLegs[0].Date.Day, 26)
}
