package main

import (
	"fmt"
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

	input := []string{}

	World := createZone()
	CurrentPosition := Coordinates{
		X: 0,
		Y: 0,
	}

	for {
		fmt.Print("> ")

		_, err := fmt.Scanln(&input)

		if nil != err {
			fmt.Println(err)
			failQuit()
		} else {

			command := input[0]

			args := input[1:]

			if function, exists := functionMap[command]; exists {
				function(&World, &CurrentPosition, args...)
			} else {
				badinput()
			}

			input = []string{}
		}
	}
}
