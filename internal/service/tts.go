package service

import (
	"fmt"
	"keyforge-cards-backend/internal/api"
	"keyforge-cards-backend/internal/database"
	log "keyforge-cards-backend/internal/logging"
	"keyforge-cards-backend/internal/model"
	"keyforge-cards-backend/internal/model/tts"
	"keyforge-cards-backend/internal/model/vault"
	"strconv"
)

var houseNames = map[string]string{
	"Brobnar":       "Brobnar",
	"Dis":           "Dis",
	"Logos":         "Logos",
	"Mars":          "Marte",
	"Sanctum":       "Santuário",
	"Saurian":       "Sauro",
	"Shadows":       "Sombras",
	"Star Alliance": "Aliança Estelar",
	"Untamed":       "Abissais",
	"Unfathomable":  "Indomados",
}

func ImportDeck(id string, lang string, sleeve string) (*tts.ObjectTTS, error) {
	vaultDeck, err := RetrieveDeck(id, lang)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	ruleImage := "https://raw.githubusercontent.com/cardsofkeyforge/json/master/decks/assets/kfQuickstartRules.png"
	backImage := fmt.Sprintf("https://raw.githubusercontent.com/cardsofkeyforge/json/master/decks/assets/%sBack.png", sleeve)

	mainDeck := tts.DefaultDeckTTS()
	mainDeck.ContainedObjects = make([]tts.CardTTS, 0)
	mainDeck.DeckIDs = make([]int, 0)
	mainDeck.CustomDeck = make(map[string]tts.CardDataTTS)

	sideDeck := (*tts.DeckTTS)(nil)

	idx := 0
	lastCardName := ""
	for _, cardId := range vaultDeck.Data.InfoId.CardIds {
		card := filterCard(cardId, &vaultDeck.Info.Cards)
		currDeck := &mainDeck
		if card.NonDeck {
			if sideDeck == nil {
				sideDeckData := tts.DefaultDeckTTS()
				sideDeckData.ContainedObjects = make([]tts.CardTTS, 0)
				sideDeckData.DeckIDs = make([]int, 0)
				sideDeckData.CustomDeck = make(map[string]tts.CardDataTTS)
				sideDeck = &sideDeckData
				sideDeck.Transform.PosX = 5 // Shifts to the side

				idx++
				sideDeck.CustomDeck[strconv.Itoa(idx)] = tts.DefaultCardDataTTS(ruleImage, backImage)
				sideDeck.ContainedObjects = append(sideDeck.ContainedObjects, tts.DefaultCardTTS(idx*100, "Regras", "Guia de Referência Rápida"))
				sideDeck.DeckIDs = append(sideDeck.DeckIDs, idx*100)
			}
			currDeck = sideDeck
		}

		if lastCardName != card.Title {
			idx++
			lastCardName = card.Title
			currDeck.CustomDeck[strconv.Itoa(idx)] =
				tts.DefaultCardDataTTS(zoomImage(card, lang), backImage)
		}

		description := ""
		if card.NonDeck {
			description = "Fora do Baralho"
		} else {
			description = houseNames[card.House]
		}
		if card.Maverick {
			description += "\n" + "Maverick"
		}
		if card.Anomaly {
			description += "\n" + "Anomalia"
		}
		if card.Enhanced {
			description += "\n" + "Propagada"
		}

		currDeck.ContainedObjects = append(currDeck.ContainedObjects, tts.DefaultCardTTS(idx*100, card.Title, description))
		currDeck.DeckIDs = append(currDeck.DeckIDs, idx*100)
	}

	ttsDeck := tts.ObjectTTS{
		ObjectStates: nil,
	}
	if sideDeck == nil {
		ttsDeck.ObjectStates = make([]tts.DeckTTS, 1)
	} else {
		ttsDeck.ObjectStates = make([]tts.DeckTTS, 2)
		ttsDeck.ObjectStates[1] = *sideDeck
	}

	ttsDeck.ObjectStates[0] = mainDeck
	return &ttsDeck, nil
}

func zoomImage(card *vault.CardVault, lang string) string {
	return fmt.Sprintf("https://media/%s/%d/%s", lang, card.Expansion, card.Number)
}

func _zoomImage(card *vault.CardVault, lang string) string {
	cr := api.OneCardRequest(int64(card.Expansion), card.Number)
	filter := cr.Filter()

	cards := make([]model.Card, 0)
	err := database.Scan(fmt.Sprintf("cards-%s", lang), filter, &cards)

	if err != nil {
		log.Error(err.Error())
		return ""
	}

	house := filterHouse(card.House, &cards)
	if house != nil {
		return house.Zoom
	}

	return ""
}

func filterHouse(house string, cards *[]model.Card) *model.House {
	for _, card := range *cards {
		for _, cardHouse := range card.Houses {
			if cardHouse.House == house {
				return &cardHouse
			}
		}

		if len(card.Houses) > 0 {
			return &card.Houses[0]
		}
	}

	return nil
}

func filterCard(uuid string, cards *[]vault.CardVault) *vault.CardVault {
	for _, card := range *cards {
		if card.Id == uuid {
			return &card
		}
	}

	return nil
}
