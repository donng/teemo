package model

import (
	"encoding/json"
)

type Wikipedia struct {
	Model
	Date    string `json:"date"`
	News    string `json:"news"`
	History string `json:"history"`
}

type Event struct {
	Time string
	Desc string
}

type Wiki struct {
	Date    string
	News    []string
	History []Event
}

func GetWikiPediaByDate(date string) (Wiki, error) {
	var err error
	var wiki Wiki
	var rawWiki Wikipedia

	if db.Where("date = ?", date).First(&rawWiki).RecordNotFound() {
		return Wiki{
			date, []string{}, []Event{},
		}, nil
	}

	err = json.Unmarshal([]byte(rawWiki.News), &wiki.News)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(rawWiki.History), &wiki.History)
	if err != nil {
		panic(err)
	}
	return wiki, nil
}
