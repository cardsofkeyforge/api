package vault

type Deck struct {
	Id        string
	Name      string
	Expansion int
}

type CardVault struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"card_title,omitempty"`
	House     string `json:"house,omitempty"`
	Type      string `json:"card_type,omitempty"`
	Image     string `json:"front_image,omitempty"`
	Text      string `json:"card_text,omitempty"`
	Trail     string `json:"traits,omitempty"`
	Number    string `json:"card_number,omitempty"`
	Rarity    string `json:"rarity,omitempty"`
	Expansion int    `json:"expansion,omitempty"`
	Maverick  bool   `json:"is_maverick,omitempty"`
	Anomaly   bool   `json:"is_anomaly,omitempty"`
	Enhanced  bool   `json:"is_enhanced,omitempty"`
	NonDeck   bool   `json:"is_non_deck,omitempty"`
}

type DeckLinkIdVault struct {
	CardIds []string `json:"cards,omitempty"`
}

type DeckLinkVault struct {
	Cards []CardVault `json:"cards,omitempty"`
}

type DeckDataVault struct {
	Name      string          `json:"name,omitempty"`
	Expansion int             `json:"expansion,omitempty"`
	Id        string          `json:"id,omitempty"`
	InfoId    DeckLinkIdVault `json:"_links,omitempty"`
}

type DeckVault struct {
	Data DeckDataVault `json:"data,omitempty"`
	Info DeckLinkVault `json:"_linked,omitempty"`
}

type DecksVault struct {
	Data    []DeckDataVault `json:"data,omitempty"`
	Counter int             `json:"count,omitempty"`
}
