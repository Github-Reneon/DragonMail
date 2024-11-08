package main

import (
	"dragonmail/ansi"
	"dragonmail/directions"
	"fmt"
	"os"
	"strings"

	"github.com/Github-Reneon/dicetools"
)

/* Looking actions

	actions used to scan world

*/

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
	for _, exit := range currentZone.Rooms[*currentCoords].Exits {
		if exit.Door && exit.Locked {
			fmt.Print("(")
		}
		switch exit.Direction {
		case directions.North:
			fmt.Print("North")
		case directions.South:
			fmt.Print("South")
		case directions.West:
			fmt.Print("West")
		case directions.East:
			fmt.Print("East")
		}
		if exit.Door && exit.Locked {
			fmt.Print(")")
		}
		fmt.Print("\r\n")
	}
}

func attack(currentZone *Zone, currentCoords *Coordinates, args ...string) {
	// Temp POC Implementation
	roll := dicetools.RollNotation("1d100")
	fmt.Println("Skill (Attack) beat an skill value of 100: roll was ", roll)

	roll = dicetools.RollNotation("1d20")
	fmt.Println("Attack beat an AC of 1: roll was ", roll)
}

/* Movement commands

	Commands used to move throughout the world
*/

func west(currentZone *Zone, currentCoords *Coordinates, args ...string){
	for _, exit := range currentZone.Rooms[*currentCoords].Exits {
		if exit.Direction == directions.West && (!exit.Door || (exit.Door && !exit.Opened)) {
			currentCoords.X -= 1
			look(currentZone, currentCoords, []string{}...)
			return
		}
	}
	fmt.Println("You bump into a wall!")
}

func east(currentZone *Zone, currentCoords *Coordinates, args ...string){
	for _, exit := range currentZone.Rooms[*currentCoords].Exits {
		if exit.Direction == directions.East && (!exit.Door || (exit.Door && !exit.Opened)) {
			currentCoords.X += 1
			look(currentZone, currentCoords, []string{}...)
			return
		}
	}
	fmt.Println("You bump into a wall!")
}

func south(currentZone *Zone, currentCoords *Coordinates, args ...string){
	for _, exit := range currentZone.Rooms[*currentCoords].Exits {
		if exit.Direction == directions.South && (!exit.Door || (exit.Door && !exit.Opened)) {
			currentCoords.Y -= 1
			look(currentZone, currentCoords, []string{}...)
			return
		}
	}
	fmt.Println("You bump into a wall!")
}

func north(currentZone *Zone, currentCoords *Coordinates, args ...string){
	for _, exit := range currentZone.Rooms[*currentCoords].Exits {
		if exit.Direction == directions.North && (!exit.Door || (exit.Door && !exit.Opened)) {
			currentCoords.Y += 1
			look(currentZone, currentCoords, []string{}...)
			return
		}
	}
	fmt.Println("You bump into a wall!")
}

/* Sys commands

*/

func badinput() {
	fmt.Println("I don't know how to do that.")
}

func quit(currentZone *Zone, currentCoords *Coordinates, args ...string) {
	os.Exit(0)
}

func failQuit() {
	os.Exit(0)
}
