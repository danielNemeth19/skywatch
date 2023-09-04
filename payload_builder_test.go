package main

import (
	"fmt"
	"testing"

	"skywatch.danielnemeth.net/internal/assert"
)

func TestPayloadBuilder_parseDateString(t *testing.T) {
	pb := PayloadBuilder{sDate: "20230326"}
	time := pb.parseDateString()
	assert.Equal(t, time.Year(), 2023)
	assert.Equal(t, time.Month(), 3)
	assert.Equal(t, time.Day(), 26)
}

func TestPayloadBuilder_setIncludeAgentIds(t *testing.T) {
	pb := PayloadBuilder{budgetAgent: false}
	includedAgentIds := pb.setIncludedAgentIds()
	fmt.Println(includedAgentIds)
	assert.Equal(t, len(includedAgentIds), 0)
}

func TestPayloadBuilder_setIncludeAgentIds_budget(t *testing.T) {
	pb := PayloadBuilder{budgetAgent: true}
	includedAgentIds := pb.setIncludedAgentIds()
	assert.Equal(t, len(includedAgentIds), 4)
}

func TestPayloadBuilder_Base_Assemble(t *testing.T) {
	pb := PayloadBuilder{
		market:      "HU",
		locale:      "hu-HU",
		currency:    "HUF",
		origin:      "BUD",
		destination: "MAD",
		sDate:       "20230326",
		budgetAgent: false,
	}
	data := pb.Assemble()
	assert.Equal(t, data.Query.Market, "HU")
	assert.Equal(t, data.Query.Locale, "hu-HU")
	assert.Equal(t, data.Query.Currency, "HUF")
	assert.Equal(t, len(data.Query.IncludedAgentsIds), 0)
}

func TestPayloadBuilder_QueryLegs_Assemble(t *testing.T) {
	pb := PayloadBuilder{
		origin:      "BUD",
		destination: "MAD",
		sDate:       "20230326",
	}
	data := pb.Assemble()
	assert.Equal(t, data.Query.QueryLegs[0].OriginPlaceId.Iata, "BUD")
	assert.Equal(t, data.Query.QueryLegs[0].DestinationPlaceId.Iata, "MAD")
	assert.Equal(t, data.Query.QueryLegs[0].Date.Year, 2023)
	assert.Equal(t, data.Query.QueryLegs[0].Date.Month, 3)
	assert.Equal(t, data.Query.QueryLegs[0].Date.Day, 26)
}
