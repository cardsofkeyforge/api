package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	log "keyforge-cards-backend/internal/logging"
	"keyforge-cards-backend/internal/model/vault"
	"math/rand"
	"net"
	"net/http"
	"time"
)

func RetrieveDeck(id string, lang string) (*vault.DeckVault, error) {
	url := fmt.Sprintf("https://www.keyforgegame.com/api/decks/%s?links=cards", id)
	method := "GET"

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

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
		if err := Body.Close(); err != nil {
			log.Error(err.Error())
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var vaultDeck vault.DeckVault
	if err = json.Unmarshal(body, &vaultDeck); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &vaultDeck, nil
}

func RetrieveRandomDeckId(set int) (string, error) {
	lastDeck, err := retrieveLastDecks(set, 1, 0)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	time.Sleep(200 * time.Millisecond)

	deckOffset := rand.Intn(lastDeck.Counter-1) + 1
	deck, err := retrieveLastDecks(set, 1, deckOffset)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	if len(deck.Data) > 0 {
		return deck.Data[0].Id, nil
	}

	return "", nil
}

func retrieveLastDecks(set int, count int, offset int) (*vault.DecksVault, error) {
	filter := ""
	if set > 0 {
		filter = fmt.Sprintf("&expansion=%d", set)
	}
	page := 1 + offset
	url := fmt.Sprintf("https://www.keyforgegame.com/api/decks?page=%d&page_size=%d&ordering=-date%s", page, count, filter)
	method := "GET"

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Error(err.Error())
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var lastDecks vault.DecksVault
	if err = json.Unmarshal(body, &lastDecks); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &lastDecks, nil
}
