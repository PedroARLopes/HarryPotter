package main

// Wizard represents a player
type Wizard struct {
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	House  HouseName `json:"house"`
	HP     int       `json:"hp"`
	Spells []*Spell  `json:"spells"`
}
