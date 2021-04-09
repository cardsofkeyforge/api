package models

type Source struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Url     string `json:"url"`
}

type Rules struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Source Source `json:"source"`
}

type House struct {
	Id     string `json:"id"`
	House  string `json:"house"`
	Normal string `json:"normal"`
	Zoom   string `json:"zoom"`
}

type Card struct {
	CardTitle  string  `json:"card_title" dynamo:",hash" index:"set-index,range"` // Hash key, a.k.a. partition key
	Set        string  `json:"set" dynamo:",range" index:"set-index,hash"`        // Range key, a.k.a. sort key
	Amber      int     `json:"amber"`
	CardNumber string  `json:"card_number"`
	CardText   string  `json:"card_text"`
	CardType   string  `json:"card_type"`
	Expansion  int64   `json:"expansion"`
	FlavorText string  `json:"flavor_text"`
	Houses     []House `json:"houses"`
	Id         string  `json:"id"`
	IsAnomaly  bool    `json:"is_anomaly"`
	IsMaverick bool    `json:"is_maverick"`
	Power      string  `json:"power"`
	Rarity     string  `json:"rarity"`
	Traits     string  `json:"traits"`
	Errata     string  `json:"errata"`
	Rules      []Rules `json:"rules"`
}
