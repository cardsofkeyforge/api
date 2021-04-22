package service

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"keyforge-cards-backend/internal/api"
	log "keyforge-cards-backend/internal/logging"

	"keyforge-cards-backend/internal/database"
	"keyforge-cards-backend/internal/model"
)

var houses = []string{
	"Brobnar",
	"Dis",
	"Logos",
	"Mars",
	"Sanctum",
	"Saurian",
	"Shadows",
	"Star Alliance",
	"Untamed",
	"Unfathomable",
}

func SearchCards(event events.APIGatewayProxyRequest) (*[]model.Card, error) {
	var table string
	results := make([]model.Card, 0)
	cards := make([]model.Card, 0)
	cr := api.NewCardRequest(&event)

	if val, ok := event.Headers["Lang"]; ok {
		table = fmt.Sprintf("cards_%s", val)
	} else {
		table = "cards_pt"
	}

	filter, values, err := cr.Filter()
	if err != nil {
		log.Error(err.Error())
		return &results, err
	}

	err = database.Scan(table, filter, values, &cards)

	//TODO improve this
	if cr.House != "" && isValidHouse(cr.House) {
		for _, c := range cards {
			for _, h := range c.Houses {
				if h.House == cr.House {
					results = append(results, c)
				}
			}
		}
	} else {
		results = append(results, cards...)
	}

	if err != nil {
		log.Error(err.Error())
		return &results, err
	}

	return &results, nil
}

func GetCard() {

}

func GetSet() {

}

func SearchSets() {

}

func isValidHouse(house string) bool {
	for _, h := range houses {
		if h == house {
			return true
		}
	}
	return false
}
