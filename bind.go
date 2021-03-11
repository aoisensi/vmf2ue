package main

import (
	"encoding/json"
	"os"
)

var bind struct {
	Materials       map[string]string `json:"materials"`
	DefaultMaterial string            `json:"default_material"`
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
