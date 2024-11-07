package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	functionMap := map[string]func(zone *Zone, coords *Coordinates, args ...string){
		"look":   look,
		"l":      look,
		"attack": attack,
		"a":      attack,
		"exits":  exits,
		"e":      exits,
		"quit":   quit,
		"qq":     quit,
	}

	input := ""

	World := createZone()
	CurrentPosition := Coordinates{
		X: 0,
		Y: 0,
	}

	for {
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			input = scanner.Text()
		}

		parts := strings.Fields(input)

		command := parts[0]

		args := parts[1:]

		if function, exists := functionMap[command]; exists {
			function(&World, &CurrentPosition, args...)
		} else {
			badinput()
		}
	}
}
