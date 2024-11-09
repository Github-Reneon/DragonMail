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
	
	if len(args) == 0 {
		// Zone Look
		fmt.Println(ansi.Green + currentZone.Name + ansi.Reset)
		fmt.Println(currentZone.Rooms[*currentCoords].Description)
		fmt.Print("Directions: ")
		exits(currentZone, currentCoords, []string{}...)
		for _, character := range currentZone.Rooms[*currentCoords].Characters {
			fmt.Println(strings.Join(character.Flags, " ") + " " + ansi.Yellow + character.Name + ansi.Reset)
		}
	} else {
		// Monster & Object Look
		for _, character := range currentZone.Rooms[*currentCoords].Characters {
			for _, tag := range character.Tags {
				if strings.Join(args, " ") == tag {
					fmt.Println(character.Name)
					fmt.Printf("HP: %d/%d\n", character.HP.Current, character.HP.Max)
					return 
				}
			}
		} 
		fmt.Printf("You don't see a %s", strings.Join(args, " "))
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

/* Doing actions

*/

func say(currentZone *Zone, currentCoords *Coordinates, args ...string) {
	fmt.Println("You say \"" + strings.Join(args, " ") + "\"")
}

func attack(currentZone *Zone, currentCoords *Coordinates, args ...string) {
	// Temp POC Implementation
	roll := dicetools.RollNotation("1d100", currentZone.Random)
	fmt.Println("Skill (Attack) beat an skill value of 100: roll was ", roll)

	roll = dicetools.RollNotation("1d20", currentZone.Random)
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
