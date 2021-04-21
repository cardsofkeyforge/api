package model

import (
	"encoding/json"
)

type SetHouse struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Set struct {
	Sequence   int                 `json:"sequence"`
	Longname   string              `json:"longname"`
	Code       int                 `json:"code"`
	Developers []string            `json:"developers"`
	Release    string              `json:"release"`
	NewCount   int                 `json:"new_count"`
	Name       string              `json:"name"`
	Icon       string              `json:"icon"`
	Houses     []SetHouse          `json:"houses"`
	Langs      []map[string]string `json:"langs"`
	Designer   string              `json:"designer"`
	CardCount  int                 `json:"card_count"`
}

func (s *Set) UnmarshalJSON(data []byte) error {
	type Alias Set
	aux := &struct {
		*Alias
		Houses []map[string]map[string]string `json:"houses"`
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for _, hu := range aux.Houses {
		for _, v := range hu {
			sh := SetHouse{
				Name: v["name"],
				Icon: v["icon"],
			}
			s.Houses = append(s.Houses, sh)
		}
	}

	return nil
}
