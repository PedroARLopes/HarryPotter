package main

import (
	"bytes"
	"encoding/json"
)

type HouseName int

const (
	Griffindor HouseName = iota
	Hufflepuff
	Ravenclaw
	Slytherin
)

var houseId = map[HouseName]string{
	Griffindor: "Griffindor",
	Hufflepuff: "Hufflepuff",
	Ravenclaw:  "Ravenclaw",
	Slytherin:  "Slytherin",
}

var houseName = map[string]HouseName{
	"Griffindor": Griffindor,
	"Hufflepuff": Hufflepuff,
	"Ravenclaw":  Ravenclaw,
	"Slytherin":  Slytherin,
}

func (h *HouseName) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(houseId[*h])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (h *HouseName) UnmarshalJSON(b []byte) error {
	var value string
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}

	*h = houseName[value]
	return nil
}
