//#!/usr/bin/env gorun
/*
	install go get github.com/erning/gorun
	run: gorun load_data.go -d <path> -l <cards_language>
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Source struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Url     string `json:"url"`
}

type Ruling struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Source Source `json:"source"`
}

type RulingJson struct {
	*Ruling
	Cards []string `json:"cards"`
}

func (c *RulingJson) UnmarshalJSON(data []byte) error {
	type Alias RulingJson
	aux := &struct {
		*Alias
		Cards []string
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.Cards = append(c.Cards, aux.Cards...)
	return nil
}

type Card struct {
	CardTitle   string   `json:"card_title" dynamo:",hash" index:"set-index,range"` // Hash key, a.k.a. partition key
	Set         string   `json:"set" dynamo:",range" index:"set-index,hash"`        // Range key, a.k.a. sort key
	Amber       int      `json:"amber"`
	CardNumber  string   `json:"card_number"`
	CardText    string   `json:"card_text"`
	CardType    string   `json:"card_type"`
	Expansion   int64    `json:"expansion"`
	FlavorText  string   `json:"flavor_text"`
	FrontImages []string `json:"front_images"`
	Houses      []string `json:"houses"`
	Id          string   `json:"id"`
	IsAnomaly   bool     `json:"is_anomaly"`
	IsMaverick  bool     `json:"is_maverick"`
	Power       string   `json:"power"`
	Rarity      string   `json:"rarity"`
	Traits      string   `json:"traits"`
	Errata      string   `json:"errata"`
	Rulings     []Ruling `json:"rulings"`
}

func (c *Card) UnmarshalJSON(data []byte) error {
	type Alias Card
	aux := &struct {
		FrontImage string `json:"front_image"`
		House      string `json:"house"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.House != "" {
		c.Houses = make([]string, 0)
		c.Houses = append(c.Houses, aux.House)
	}
	if aux.FrontImage != "" {
		c.FrontImages = make([]string, 0)
		c.FrontImages = append(c.FrontImages, aux.FrontImage)
	}

	return nil
}

func main() {

	dir := flag.String("d", ".", "Directory to load data from")
	lang := flag.String("l", "", "Required. The Language of the cards (pt, en, es...)")
	flag.Parse()

	if *lang == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	files, err := getFileList(dir)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rulings, err := parseRulings(fmt.Sprintf("%srulings.json", *dir))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	cards := parseCards(&files, rulings)

	sess, _ := session.NewSession(&aws.Config{Region: aws.String("eu-north-1")})
	db := dynamo.New(sess)

	tables, err := db.ListTables().All()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tableName := fmt.Sprintf("cards_%s", *lang)
	if !contains(tables, tableName) {
		fmt.Printf("table %s does not exist, trying to create it\n", tableName)
		err = createTable(db, tableName)
		if err != nil {
			fmt.Printf("failed to create table %s \n", tableName)
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	table := db.Table(tableName)
	total := 0
	for _, value := range *cards {
		total += len(value)
		for _, c := range value {
			fmt.Printf("saving %s into %s \n", c.CardTitle, tableName)
			err = table.Put(c).Run()
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	fmt.Printf("%d stored in %s", total, tableName)

}

func createTable(db *dynamo.DB, tableName string) error {
	err := db.CreateTable(tableName, Card{}).Provision(5, 5).ProvisionIndex("set-index", 5, 5).Run()
	return err
}

func parseCards(files *[]string, rulings *[]RulingJson) *map[string]map[string]Card {
	cards := make(map[string]map[string]Card)

	for _, v := range *files {
		if strings.Contains(v, "errata") || strings.Contains(v, "ruling") {
			continue
		} else {

			var c Card

			if err := unmarshal(v, &c); err != nil {
				fmt.Printf("%v", err.Error())
				os.Exit(1)
			}

			for _, r := range *rulings {
				if contains(r.Cards, c.Id) {
					c.Rulings = append(c.Rulings, *r.Ruling)
				}
			}

			split := strings.Split(v, "/")
			set := split[len(split)-2]
			c.Set = set

			if _, ok := cards[set]; !ok {
				cards[set] = make(map[string]Card)
			}

			if val, ok := cards[set][c.CardTitle]; ok {
				c.Houses = append(c.Houses, val.Houses...)
				c.FrontImages = append(c.FrontImages, val.FrontImages...)
				cards[set][c.CardTitle] = c
			} else {
				cards[set][c.CardTitle] = c
			}
		}
	}
	return &cards
}

func parseRulings(rulingsPath string) (*[]RulingJson, error) {
	rs := make([]RulingJson, 0)

	if err := unmarshal(rulingsPath, &rs); err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}

	return &rs, nil
}

func unmarshal(file string, v interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)

	return err
}
func getFileList(dir *string) ([]string, error) {

	fileList := make([]string, 0)
	e := filepath.Walk(*dir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".json") {
			fileList = append(fileList, path)
		}
		return err
	})

	if e != nil {
		panic(e)
	}

	return fileList, nil
}

func contains(list []string, value string) bool {
	if len(list) == 0 {
		return false
	}

	for _, v := range list {
		return v == value
	}

	return false
}
