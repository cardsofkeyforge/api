//#!/usr/bin/env gorun
/*
	install go get github.com/erning/gorun
	run: gorun load.go -d <path> -l <cards_language>
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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
	CardTitle  string  `json:"card_title"`
	Set        string  `json:"set" dynamo:",hash"`
	Amber      int     `json:"amber"`
	CardNumber string  `json:"card_number" dynamo:",range"`
	CardText   string  `json:"card_text"`
	CardType   string  `json:"card_type"`
	Expansion  int64   `json:"expansion"`
	FlavorText string  `json:"flavor_text"`
	Houses     []House `json:"houses"`
	IsAnomaly  bool    `json:"is_anomaly"`
	IsMaverick bool    `json:"is_maverick"`
	Power      string  `json:"power"`
	Armor      string  `json:"armor"`
	Rarity     string  `json:"rarity"`
	Traits     string  `json:"traits"`
	Errata     string  `json:"errata"`
	Rules      []Rules `json:"rules"`
}

func (c *Card) UnmarshalJSON(data []byte) error {
	type Alias Card
	aux := &struct {
		Power interface{} `json:"power"`
		Armor interface{} `json:"armor"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var power string
	if aux.Power != nil && aux.Power != 0 {
		power = fmt.Sprintf("%v", aux.Power)
	} else {
		power = "0"
	}

	var armor string
	if aux.Armor != nil && aux.Armor != 0 {
		armor = fmt.Sprintf("%v", aux.Armor)
	} else {
		armor = "0"
	}

	c.CardTitle = aux.CardTitle
	c.Set = aux.Set
	c.Amber = aux.Amber
	c.CardNumber = aux.CardNumber
	c.CardText = aux.CardText
	c.CardType = aux.CardType
	c.Expansion = aux.Expansion
	c.FlavorText = aux.FlavorText
	c.Houses = aux.Houses
	c.IsAnomaly = aux.IsAnomaly
	c.IsMaverick = aux.IsMaverick
	c.Power = power
	c.Armor = armor
	c.Rarity = aux.Rarity
	c.Traits = aux.Traits
	c.Errata = aux.Errata
	c.Rules = aux.Rules
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

	cards := parseCards(&files)

	sess, _ := session.NewSession(&aws.Config{Region: aws.String("eu-north-1")})
	db := dynamo.New(sess)

	tables, err := db.ListTables().All()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tableName := fmt.Sprintf("cards-%s", *lang)
	if !contains(tables, tableName) {
		fmt.Printf("table %s does not exist, trying to create it\n", tableName)
		err = createTable(db, tableName)
		if err != nil {
			fmt.Printf("failed to create table %s \n", tableName)
			fmt.Println(err.Error())
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}

	fmt.Printf("Waiting for %s to be active...", tableName)

	expectedStatus := "active"
	var tableStatus string
	for strings.ToLower(tableStatus) != expectedStatus {
		t, _ := db.Client().DescribeTable(&dynamodb.DescribeTableInput{TableName: &tableName})
		tableStatus = *t.Table.TableStatus
		time.Sleep(1 * time.Second)
	}

	table := db.Table(tableName)
	total := 0
	for _, c := range *cards {
		fmt.Printf("saving %s into %s \n", c.CardTitle, tableName)
		err = table.Put(c).Run()
		if err != nil {
			fmt.Println(err.Error())
		}
		total++
	}
	fmt.Printf("%d stored in %s", total, tableName)

}

func createTable(db *dynamo.DB, tableName string) error {
	err := db.CreateTable(tableName, Card{}).
		Provision(5, 5).
		Run()
	return err
}

func parseCards(files *[]string) *[]Card {
	cards := make([]Card, 0)

	for _, v := range *files {

		var c Card

		if err := unmarshal(v, &c); err != nil {
			fmt.Printf("%v", err.Error())
			os.Exit(1)
		}

		split := strings.Split(v, "/")
		set := split[len(split)-2]
		c.Set = set

		cards = append(cards, c)
	}

	return &cards
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
