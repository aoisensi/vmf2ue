package main

import (
	"encoding/json"
	"os"
)

var bind struct {
	Materials map[string]struct {
		Asset string `json:"asset"`
		W     int    `json:"w"`
		H     int    `json:"h"`
	} `json:"materials"`
	Props map[string]struct {
		Asset string `json:"asset"`
	} `json:"props"`
}

func readBind() {
	f, err := os.Open("binds.json")
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(f).Decode(&bind); err != nil {
		panic(err)
	}
}
