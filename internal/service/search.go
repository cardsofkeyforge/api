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

func SearchCards(event events.APIGatewayProxyRequest) (result *[]model.Card, err error) {
	var table string
	cards := make([]model.Card, 0)

	table = fmt.Sprintf("cards_%s", language(event))
	cr := api.NewCardRequest(&event)
	filter := cr.Filter()

	err = database.Scan(table, filter, &cards)

	if err != nil {
		log.Error(err.Error())
		return
	}

	houseFilter(cr, cards, result)
	return
}

func houseFilter(cr *api.CardRequest, cards []model.Card, results *[]model.Card) []model.Card {
	if cr.House != "" && isValidHouse(cr.House) {
		for _, c := range cards {
			for _, h := range c.Houses {
				if h.House == cr.House {
					*results = append(*results, c)
				}
			}
		}
	} else {
		*results = append(*results, cards...)
	}
	return *results
}

func language(event events.APIGatewayProxyRequest) string {
	if val, ok := event.Headers["Accept-Language"]; ok {
		return val
	} else {
		return "pt"
	}
}

func isValidHouse(house string) bool {
	for _, h := range houses {
		if h == house {
			return true
		}
	}
	return false
}
