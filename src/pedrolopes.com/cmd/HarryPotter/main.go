package main

import "log"

func main() {
	for name, wizard := range AppState.Wizards {
		log.Print(name)
		log.Printf("%+v", wizard)
		for _, spell := range wizard.Spells {
			log.Printf("%+v", spell)
			log.Print(spellTypesId[spell.Type])
		}
	}
}
