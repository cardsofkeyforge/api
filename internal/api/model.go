package api

import (
	"github.com/aws/aws-lambda-go/events"
	"keyforge-cards-backend/internal/database"
	"strconv"
	"strings"
)

type CardRequest struct {
	Name     string
	Set      string
	Amber    int
	Type     string
	Anomaly  *bool
	Maverick *bool
	Rarity   string
	Power    string
	House    string
}

func NewCardRequest(request *events.APIGatewayProxyRequest) *CardRequest {
	cr := CardRequest{}
	parameters := request.QueryStringParameters
	cr.Name = strings.Title(strings.ToLower(parameters["name"]))
	cr.Set = strings.ToLower(parameters["set"])
	cr.Type = strings.Title(strings.ToLower(strings.ReplaceAll(parameters["type"], "_", " ")))
	cr.Rarity = strings.Title(strings.ToLower(parameters["rarity"]))
	cr.Power = parameters["power"]
	cr.House = strings.Title(strings.ToLower(strings.ReplaceAll(parameters["house"], "_", " ")))

	atoi, err := strconv.Atoi(parameters["amber"])
	if err == nil {
		cr.Amber = atoi
	}
	b1, err := strconv.ParseBool(parameters["anomaly"])
	if err == nil {
		cr.Anomaly = &b1
	}
	b2, err := strconv.ParseBool(parameters["maverick"])
	if err == nil {
		cr.Maverick = &b2
	}
	return &cr
}

func (cr *CardRequest) Filter() *database.Filter {
	fb := database.FilterBuilder{}
	if cr.Name != "" {
		fb.Contains("CardTitle", cr.Name).And()
	}
	if cr.Set != "" {
		fb.Eq("Set", cr.Set).And()
	}
	if cr.Amber > 0 {
		fb.Ge("Amber", cr.Amber).And()
	}
	if cr.Power != "" {
		fb.Ge("Power", cr.Power).And()
	}
	if cr.Type != "" {
		fb.Eq("CardType", cr.Type).And()
	}
	if cr.Rarity != "" {
		fb.Eq("Rarity", cr.Rarity).And()
	}
	if cr.Anomaly != nil {
		fb.Eq("IsAnomaly", *cr.Anomaly).And()
	}
	if cr.Anomaly != nil {
		fb.Eq("IsMaverick", *cr.Maverick)
	}

	return fb.Build()
}
