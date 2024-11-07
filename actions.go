package main

import (
	"dragonmail/ansi"
	"fmt"
	"os"
	"strings"

	"github.com/Github-Reneon/dicetools"
)

func look(currentZone *Zone, currentCoords *Coordinates, args ...string) {
	// Temp POC Implementation
	fmt.Println(ansi.Green + currentZone.Name + ansi.Reset)
	fmt.Println(currentZone.Rooms[*currentCoords].Description)
	for _, character := range currentZone.Rooms[*currentCoords].Characters {
		fmt.Println(strings.Join(character.Tags, " ") + " " + ansi.Yellow + character.Name + ansi.Reset)
	}
}

func exits(currentZone *Zone, currentCoords *Coordinates, args ...string) {
	// Temp POC Implementation

}

func attack(currentZone *Zone, currentCoords *Coordinates, args ...string) {
	// Temp POC Implementation
	roll := dicetools.RollNotation("1d100")
	fmt.Println("Skill (Attack) beat an skill value of 100: roll was ", roll)

	roll = dicetools.RollNotation("1d20")
	fmt.Println("Attack beat an AC of 1: roll was ", roll)
}

func badinput() {
	fmt.Println("I don't know how to do that.")
}

func quit(currentZone *Zone, currentCoords *Coordinates, args ...string) {
	os.Exit(0)
}

func failQuit() {
	os.Exit(0)
}
