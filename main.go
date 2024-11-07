package main

import "fmt"

func main() {
	
	functionMap := map[string]func(){
		"look": look,
		"l": look,
		"attack": attack,
		"a": attack,
		"exits": exits,
		"e": exits,
		"quit": quit,
		"qq": quit,
	}

	input := ""

	for {
		fmt.Print("> ")

		fmt.Scanln(&input)

		if function, exists := functionMap[input]; exists {
			function()
		} else {
			badinput()
		}
	}
}