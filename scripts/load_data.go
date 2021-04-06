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

type Card struct {
	CardTitle  string `json:"card_title" dynamo:",hash" index:"set-index,range"` // Hash key, a.k.a. partition key
	Set        string `json:"set" dynamo:",range" index:"set-index,hash"`        // Range key, a.k.a. sort key
	Amber      int    `json:"amber"`
	CardNumber string `json:"card_number"`
	CardText   string `json:"card_text"`
	CardType   string `json:"card_type"`
	Expansion  int64  `json:"expansion"`
	FlavorText string `json:"flavor_text"`
	FrontImage string `json:"front_image"`
	House      string `json:"house"`
	Id         string `json:"id"`
	IsAnomaly  bool   `json:"is_anomaly"`
	IsMaverick bool   `json:"is_maverick"`
	Power      string `json:"power"`
	Rarity     string `json:"rarity"`
	Traits     string `json:"traits"`
	Errata     string `json:"errata"`
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

	cards := parseCards(files)

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
		for i, c := range value {
			fmt.Printf("%d: saving %s into %s \n", i, c.CardTitle, tableName)
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

func contains(list []string, value string) bool {
	if len(list) == 0 {
		return false
	}

	for _, v := range list {
		return v == value
	}

	return false
}

func parseCards(files []string) *map[string][]Card {
	cards := make(map[string][]Card)
	for _, v := range files {
		if strings.Contains(v, "errata") || strings.Contains(v, "ruling") {
			continue
		} else {
			var c Card

			err := unmarshalJson(v, &c)

			if err != nil {
				fmt.Printf("%v", err)
			}

			split := strings.Split(v, "/")
			set := split[len(split)-2]
			c.Set = set
			if cards[set] != nil {
				cards[set] = append(cards[set], c)
			} else {
				slice := make([]Card, 0)
				slice = append(slice, c)
				cards[set] = slice
			}

		}
	}

	return &cards
}

func unmarshalJson(file string, v *Card) error {
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
