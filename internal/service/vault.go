package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	log "keyforge-cards-backend/internal/logging"
	"keyforge-cards-backend/internal/model/vault"
	"net/http"
)

func RetrieveDeck(id string, lang string) (*vault.DeckVault, error) {
	url := fmt.Sprintf("https://www.keyforgegame.com/api/decks/%s?links=cards", id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	req.Header.Add("Accept-Language", lang)

	res, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var vaultDeck vault.DeckVault
	err = json.Unmarshal(body, &vaultDeck)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &vaultDeck, nil
}
