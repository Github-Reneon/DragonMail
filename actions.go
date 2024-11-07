package main

import (
	"fmt"
	"os"
	"github.com/Github-Reneon/dicetools"
)

func look() {
	// -- Temp POC Implementation
	fmt.Println("You see a verdant forest, with a river to the north.")
}

func exits() {
	fmt.Println("north west (up)")
}

func attack() {
	// Temp POC Implementation
	roll := dicetools.RollNotation("1d100")

	fmt.Println("Skill (Attack) beat an skill value of 100: roll was ", roll)

	roll = dicetools.RollNotation("1d20")

	fmt.Println("Attack beat an AC of 1: roll was ", roll)
}

func badinput() {
	fmt.Println("I don't know how to do that.")
}

func quit() {
	os.Exit(0)
}