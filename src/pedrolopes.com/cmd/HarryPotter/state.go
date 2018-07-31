package main

// State contains the wizards loaded into memory
type State struct {
	Wizards map[string]*Wizard
}

// NewState returns a new instance of state
func NewState() *State {
	return &State{
		Wizards: map[string]*Wizard{},
	}
}
