package model

type Source struct {
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
	Url     string `json:"url,omitempty"`
}

type Rules struct {
	Title  string `json:"title,omitempty"`
	Text   string `json:"text,omitempty"`
	Source Source `json:"source,omitempty"`
}

type House struct {
	Id     string `json:"id,omitempty"`
	House  string `json:"house,omitempty"`
	Normal string `json:"normal,omitempty"`
	Zoom   string `json:"zoom,omitempty"`
}

type Card struct {
	CardTitle  string  `json:"card_title,omitempty"`
	Set        string  `json:"set,omitempty" dynamo:",hash"`
	Amber      int     `json:"amber,omitempty"`
	CardNumber string  `json:"card_number,omitempty"  dynamo:",range"`
	CardText   string  `json:"card_text,omitempty"`
	CardType   string  `json:"card_type,omitempty"`
	Expansion  int64   `json:"expansion,omitempty"`
	FlavorText string  `json:"flavor_text,omitempty"`
	Houses     []House `json:"houses,omitempty"`
	IsAnomaly  bool    `json:"is_anomaly,omitempty"`
	IsMaverick bool    `json:"is_maverick,omitempty"`
	Power      string  `json:"power,omitempty"`
	Armor      string  `json:"armor,omitempty"`
	Rarity     string  `json:"rarity,omitempty"`
	Traits     string  `json:"traits,omitempty"`
	Errata     string  `json:"errata,omitempty"`
	Rules      []Rules `json:"rules,omitempty"`
}
