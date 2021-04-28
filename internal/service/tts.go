package service

import (
	log "keyforge-cards-backend/internal/logging"
	"keyforge-cards-backend/internal/model/tts"
	"keyforge-cards-backend/internal/model/vault"
	"strconv"
)

func ImportDeck(id string, lang string) (*tts.ObjectTTS, error) {
	vaultDeck, err := RetrieveDeck(id, lang)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	mainDeck := tts.DefaultDeckTTS()
	mainDeck.ContainedObjects = make([]tts.CardTTS, 36)
	mainDeck.DeckIDs = make([]int, 36)
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
			}
			currDeck = sideDeck
		}

		if lastCardName != card.Title {
			idx++
			lastCardName = card.Title
			currDeck.CustomDeck[strconv.Itoa(idx)] = tts.DefaultCardDataTTS("", "")
		}

		// TODO ADD CARD TO CURR DECK
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

func filterCard(uuid string, cards *[]vault.CardVault) *vault.CardVault {
	for _, card := range *cards {
		if card.Id == uuid {
			return &card
		}
	}

	return nil
}
