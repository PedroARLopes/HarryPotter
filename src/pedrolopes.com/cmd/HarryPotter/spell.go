package main

import (
	"bytes"
	"encoding/json"
)

type SpellType int

const (
	Offensive SpellType = iota
	Defensive
)

var spellTypesId = map[SpellType]string{
	Offensive: "Offensive",
	Defensive: "Defensive",
}

var spellTypesName = map[string]SpellType{
	"Offensive": Offensive,
	"Defensive": Defensive,
}

// Spell represents a move that be used by a Wizard
type Spell struct {
	Name  string    `json:"name"`
	Type  SpellType `json:"type"`
	Value int       `json:"value"`
}

func (s *SpellType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(spellTypesId[*s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *SpellType) UnmarshalJSON(b []byte) error {
	var value string
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}

	*s = spellTypesName[value]
	return nil
}
