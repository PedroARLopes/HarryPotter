package main

import (
	"encoding/json"
	"log"
	"os"
)

var AppState *State

func init() {
	file, err := os.Open("datasets/wizards.json")
	if err != nil {
		log.Fatal(err)
	}

	AppState = NewState()

	dec := json.NewDecoder(file)
	for dec.More() {

		var wizard *Wizard
		if err = dec.Decode(&wizard); err != nil {
			log.Print(err)
			continue
		}

		AppState.Wizards[wizard.Name] = wizard
	}
}
